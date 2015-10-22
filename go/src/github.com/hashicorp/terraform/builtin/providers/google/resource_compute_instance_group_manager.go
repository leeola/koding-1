package google

import (
	"fmt"
	"log"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceComputeInstanceGroupManager() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeInstanceGroupManagerCreate,
		Read:   resourceComputeInstanceGroupManagerRead,
		Update: resourceComputeInstanceGroupManagerUpdate,
		Delete: resourceComputeInstanceGroupManagerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"base_instance_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"instance_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"instance_template": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"target_pools": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set: func(v interface{}) int {
					return hashcode.String(v.(string))
				},
			},

			"target_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},

			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"self_link": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeInstanceGroupManagerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	// Get group size, default to 1 if not given
	var target_size int64 = 1
	if v, ok := d.GetOk("target_size"); ok {
		target_size = int64(v.(int))
	}

	// Build the parameter
	manager := &compute.InstanceGroupManager{
		Name:             d.Get("name").(string),
		BaseInstanceName: d.Get("base_instance_name").(string),
		InstanceTemplate: d.Get("instance_template").(string),
		TargetSize:       target_size,
	}

	// Set optional fields
	if v, ok := d.GetOk("description"); ok {
		manager.Description = v.(string)
	}

	if attr := d.Get("target_pools").(*schema.Set); attr.Len() > 0 {
		var s []string
		for _, v := range attr.List() {
			s = append(s, v.(string))
		}
		manager.TargetPools = s
	}

	log.Printf("[DEBUG] InstanceGroupManager insert request: %#v", manager)
	op, err := config.clientCompute.InstanceGroupManagers.Insert(
		config.Project, d.Get("zone").(string), manager).Do()
	if err != nil {
		return fmt.Errorf("Error creating InstanceGroupManager: %s", err)
	}

	// It probably maybe worked, so store the ID now
	d.SetId(manager.Name)

	// Wait for the operation to complete
	err = computeOperationWaitZone(config, op, d.Get("zone").(string), "Creating InstanceGroupManager")
	if err != nil {
		return err
	}

	return resourceComputeInstanceGroupManagerRead(d, meta)
}

func resourceComputeInstanceGroupManagerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	manager, err := config.clientCompute.InstanceGroupManagers.Get(
		config.Project, d.Get("zone").(string), d.Id()).Do()
	if err != nil {
		if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == 404 {
			// The resource doesn't exist anymore
			d.SetId("")

			return nil
		}

		return fmt.Errorf("Error reading instance group manager: %s", err)
	}

	// Set computed fields
	d.Set("fingerprint", manager.Fingerprint)
	d.Set("instance_group", manager.InstanceGroup)
	d.Set("target_size", manager.TargetSize)
	d.Set("self_link", manager.SelfLink)

	return nil
}
func resourceComputeInstanceGroupManagerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.Partial(true)

	// If target_pools changes then update
	if d.HasChange("target_pools") {
		var targetPools []string
		if attr := d.Get("target_pools").(*schema.Set); attr.Len() > 0 {
			for _, v := range attr.List() {
				targetPools = append(targetPools, v.(string))
			}
		}

		// Build the parameter
		setTargetPools := &compute.InstanceGroupManagersSetTargetPoolsRequest{
			Fingerprint: d.Get("fingerprint").(string),
			TargetPools: targetPools,
		}

		op, err := config.clientCompute.InstanceGroupManagers.SetTargetPools(
			config.Project, d.Get("zone").(string), d.Id(), setTargetPools).Do()
		if err != nil {
			return fmt.Errorf("Error updating InstanceGroupManager: %s", err)
		}

		// Wait for the operation to complete
		err = computeOperationWaitZone(config, op, d.Get("zone").(string), "Updating InstanceGroupManager")
		if err != nil {
			return err
		}

		d.SetPartial("target_pools")
	}

	// If instance_template changes then update
	if d.HasChange("instance_template") {
		// Build the parameter
		setInstanceTemplate := &compute.InstanceGroupManagersSetInstanceTemplateRequest{
			InstanceTemplate: d.Get("instance_template").(string),
		}

		op, err := config.clientCompute.InstanceGroupManagers.SetInstanceTemplate(
			config.Project, d.Get("zone").(string), d.Id(), setInstanceTemplate).Do()
		if err != nil {
			return fmt.Errorf("Error updating InstanceGroupManager: %s", err)
		}

		// Wait for the operation to complete
		err = computeOperationWaitZone(config, op, d.Get("zone").(string), "Updating InstanceGroupManager")
		if err != nil {
			return err
		}

		d.SetPartial("instance_template")
	}

	// If size changes trigger a resize
	if d.HasChange("target_size") {
		if v, ok := d.GetOk("target_size"); ok {
			// Only do anything if the new size is set
			target_size := int64(v.(int))

			op, err := config.clientCompute.InstanceGroupManagers.Resize(
				config.Project, d.Get("zone").(string), d.Id(), target_size).Do()
			if err != nil {
				return fmt.Errorf("Error updating InstanceGroupManager: %s", err)
			}

			// Wait for the operation to complete
			err = computeOperationWaitZone(config, op, d.Get("zone").(string), "Updating InstanceGroupManager")
			if err != nil {
				return err
			}
		}

		d.SetPartial("target_size")
	}

	d.Partial(false)

	return resourceComputeInstanceGroupManagerRead(d, meta)
}

func resourceComputeInstanceGroupManagerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	zone := d.Get("zone").(string)
	op, err := config.clientCompute.InstanceGroupManagers.Delete(config.Project, zone, d.Id()).Do()
	if err != nil {
		return fmt.Errorf("Error deleting instance group manager: %s", err)
	}

	// Wait for the operation to complete
	err = computeOperationWaitZone(config, op, d.Get("zone").(string), "Deleting InstanceGroupManager")
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}