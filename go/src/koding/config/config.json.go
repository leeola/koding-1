// Code generated by go-bindata.
// sources:
// config.json
// DO NOT EDIT!

package config

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x94\x41\x6a\xc3\x30\x10\x45\xf7\x3a\xc5\x30\x6b\x47\x4a\xb2\x29\xe8\x2a\xa5\x0b\xd7\x1a\x12\x61\x7b\x24\xa4\xb1\x28\x14\xdf\xbd\xd8\x49\x5a\x25\x10\x68\x37\x25\x59\xea\xff\xf9\xf0\x78\x0b\x7d\x2a\x00\x4c\x61\x12\xca\x68\x61\x79\x01\xa0\xa3\xa2\xfb\xe0\x3c\x1f\x74\x17\x46\xb4\x80\xbb\xfd\x8b\xde\xea\xad\xde\xa1\x02\x98\x9b\x65\x44\xec\x62\xf0\x2c\xd5\x4e\x26\x66\x1a\x32\xa5\x42\x09\x2d\xbc\xae\x29\x9c\xdb\xf5\x82\xb8\xf8\x14\x78\x24\x96\xea\x60\xad\x62\x0a\x6e\xea\xc4\x07\xc6\xa6\xce\xc7\x96\xdb\x03\x39\xfc\xce\xde\x7e\x6a\x9c\xd2\xb0\xe0\x1d\x45\xa2\x35\x46\x2a\x6a\xd3\x7b\xa1\xcb\x68\x6e\xfe\x84\xe2\xa8\xd0\x10\xe2\x5a\x35\x37\xc5\x05\xe7\x2a\xcf\x2d\xbb\xf7\xf0\xf1\x1b\x46\x47\x65\x73\x9f\x53\x55\x63\xf4\xf1\xdf\x25\x66\x6b\x4c\xac\xe9\x36\xc6\xc7\x47\xb3\x98\xcf\x1a\xe3\xfe\x2e\xe9\x8d\xc7\xee\x48\x5d\xff\x18\x32\x4f\x28\x4f\x63\xf4\x1a\xf7\xa4\x75\xf9\x02\xd4\xac\xbe\x02\x00\x00\xff\xff\xc1\x6e\x85\xfb\x3b\x04\x00\x00")

func configJsonBytes() ([]byte, error) {
	return bindataRead(
		_configJson,
		"config.json",
	)
}

func configJson() (*asset, error) {
	bytes, err := configJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.json", size: 1083, mode: os.FileMode(420), modTime: time.Unix(1446555960, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config.json": configJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"config.json": &bintree{configJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

