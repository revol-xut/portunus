// Code generated by go-bindata. DO NOT EDIT.
// sources:
// static/css/portunus.css

package static


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataCssPortunuscss = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x57\xdb\x6e\xdb\x38\x13\xbe\xff\x9f\x42\x68\x51\xe0\x6f\x6a\x0a\xb2" +
	"\x83\x06\x0d\x85\xcd\x8d\x11\xbf\xc4\xee\x5e\x50\xd2\x28\x22\x42\x91\x02\x49\xf9\x50\xc1\xef\xbe\xe0\xc9\x26\x6d" +
	"\x35\x4e\x50\x60\xf7\x22\xb1\x44\xce\x7c\x73\xe4\xe8\xe3\xdd\xe2\x0e\x57\xd0\x0a\x09\x8b\x3b\x4c\x5a\x0d\x72\xaa" +
	"\xc4\x1e\x29\xfa\x93\xf2\x17\x5c\x09\xd9\x80\x44\x95\xd8\x1f\x3b\xdd\xb3\xa9\x15\x5c\x9b\x3d\xc0\xcb\x87\x61\x5f" +
	"\xda\xd7\x96\xf4\x94\x1d\x30\x22\xc3\xc0\x00\xa9\x83\xd2\xd0\x2f\x32\xf7\x8b\x46\xba\xc8\x14\xe1\x0a\x29\x90\xb4" +
	"\x2d\x19\xe5\x80\x3a\xa0\x2f\x9d\xc6\xcb\xfc\xbb\x05\xcd\x28\x1f\x46\xbd\xb0\x8f\xd5\xa8\xb5\xe0\x53\x8c\x4b\x79" +
	"\x07\x92\xea\x23\x59\x10\xdc\x8a\x7a\x54\x0b\x82\x3b\xb1\x05\xb9\x20\x98\xd4\x9a\x6e\x61\xaa\x05\x13\x12\x57\x6c" +
	"\x84\x52\xc3\x5e\xa3\x06\x6a\x21\x89\xa6\x82\x63\x2e\x38\x1c\xf3\x5a\xf4\x3d\x41\x0a\x06\x22\x89\x86\x06\x31\xaa" +
	"\xf4\x93\x5b\xc5\x8c\x28\x8d\xea\x8e\xb2\x66\x6a\xa8\x1a\x18\x39\x38\x25\xe3\xcf\xa2\x12\xcd\x61\xea\x89\x7c\xa1" +
	"\x1c\x17\xe5\x40\x9a\xc6\x64\xa5\x38\x72\xb2\x9d\x2a\x52\xbf\xbe\x48\x31\xf2\x06\x7f\x5e\xaf\xd7\x65\xf0\x82\xd4" +
	"\xaf\x66\x3f\xbb\x9b\x7c\xa0\xab\xfc\xbb\x84\x3e\x89\xdd\x2d\x59\xb1\x91\x4d\x3d\xd9\xa3\x1d\x6d\x74\x87\x7f\x14" +
	"\xc5\xb0\x2f\x83\xc1\x8c\x8c\x5a\x9c\xad\x66\x45\x00\x52\x1a\x29\x7d\x60\x60\x3d\x2d\x83\xdb\x2d\x83\xbd\x87\xcc" +
	"\x18\x9d\xb6\x20\x35\xad\x09\x43\x84\xd1\x17\x8e\x7b\xda\x34\x0c\xce\xfb\xb9\x1a\x48\x0d\x72\x32\x5a\xd8\x18\x3a" +
	"\x6f\x65\xdd\xf2\x1c\xb4\x7b\x40\x32\xd4\xcc\x78\x10\xb5\x41\x1c\x88\x51\x25\xa7\x2c\x52\x6e\x23\xae\x98\xa8\x5f" +
	"\x7d\x72\x7c\x2d\x67\xab\x74\x15\x67\x0c\x9a\xd7\xa3\x94\xc0\x75\x92\xf4\xcd\x66\x13\xcb\xf8\x6e\x58\xc4\x4b\xae" +
	"\x51\xe2\x15\xdb\x42\x09\xcc\xf3\xf3\x73\xea\xde\xb1\x27\x94\xbf\x5d\x94\x6c\x25\xa1\x9f\x2f\x8f\xd5\xce\x77\xb4" +
	"\x81\x08\x62\xb9\x32\x18\xc7\xbc\xd2\x7c\x3e\x3f\x27\x14\x97\xe0\xd9\xda\xa5\xa7\xa7\xdc\x75\x54\x03\xb2\x55\xc4" +
	"\x5c\xec\x24\x19\x5c\x5d\x76\x4e\x80\x0b\xd9\x13\x16\x97\xca\x38\x67\x3c\x30\x7f\x68\x90\xb4\x27\xf2\x90\x64\xa2" +
	"\x28\x36\x3e\x13\x16\xbb\x74\xa7\x1f\x17\x57\x5a\x21\xd5\x57\xeb\x2e\xdf\x57\xcb\xd7\x49\x7f\x78\xd8\x1c\x5b\x21" +
	"\xfb\xa7\xbb\x53\xa7\xf9\xf7\x6f\x61\x05\x69\x31\x60\x97\x52\xb3\x93\xe5\x52\xec\x10\x23\x15\xb0\x53\x0a\x5d\xee" +
	"\xce\x21\x16\xf9\x8f\x19\xf1\x4c\x0d\x84\xe7\x66\x11\x81\x94\x42\x4e\x3e\xd9\x88\x41\xab\x71\x91\x9b\x5a\xfa\xb8" +
	"\x25\x34\x91\xb6\x1b\x4c\xf6\xbd\xa1\x5b\xbb\xb4\x25\x6c\x84\x0b\xfb\xbe\xc4\x45\xf1\xe5\x52\x17\x73\xa1\xff\x9f" +
	"\x53\x33\x08\xcd\xa1\xfd\x3a\x83\x75\x29\x32\x5d\x1d\x77\x5f\x85\xe5\xb0\xcf\x94\x60\xb4\xc9\xec\xb4\xf1\x93\x59" +
	"\x92\x86\x8e\x0a\xdf\x87\x61\x1c\xa5\xe1\x62\xd8\x3e\x98\xa5\xe4\xcd\x79\x6b\x1d\x8d\x5c\xb6\x25\x9c\x3c\xbc\xcb" +
	"\xca\xe7\xc7\xc7\xc7\x79\x61\xdf\x07\xb3\x7b\xbe\xe8\x09\xd0\xfd\xfd\xfd\x2c\x50\x5c\x9d\x54\x63\x5d\x14\x65\x72" +
	"\xe8\xd7\xeb\x5b\x08\xf3\x21\x14\xc5\x4d\xbd\x37\xa2\x89\xc5\x66\x03\x0b\xf0\xa6\xb8\xbe\x3d\x90\x14\xbb\x1b\xc5" +
	"\xbe\x98\x43\x67\x88\x93\x4c\x66\x1f\xe7\x87\x46\x32\x9a\x7d\xb7\xfc\xdb\xed\x33\xe3\xac\xe7\x13\x53\x3a\xf3\xc3" +
	"\xa1\x0b\xde\xde\xbf\x05\xe1\xde\xeb\x0e\xea\x57\x68\x16\x33\x52\xae\xc5\xbc\xc0\x37\x7b\xcc\x5d\xa2\xde\x9b\xd1" +
	"\xc4\x42\x60\x40\x1f\x30\x14\x82\x4c\x0b\x53\xd6\x82\x6b\xe0\x1a\x7f\xfa\x6b\xf5\xb0\x5c\x7e\xba\x61\x7e\xe4\xef" +
	"\x75\xc0\x34\x4e\xf0\xe2\xeb\x07\xdd\x28\x66\xdd\xb0\xb8\x7f\xea\xc3\x00\x7f\x58\xdc\x4a\xec\xff\x9e\xc8\x30\x00" +
	"\x91\x84\xd7\x9e\x5a\xa0\x5e\xfc\x44\x57\x8b\x3b\xa8\x5e\xa9\xbe\x5a\x8f\xe9\x53\x39\x43\x9a\xde\xe9\x84\x3f\x86" +
	"\x51\x90\xbf\x4c\xcb\xa5\xa6\x3d\x99\x49\x37\xcc\x4f\x9f\x04\xea\xe2\x50\x22\x09\xa4\x11\x9c\x1d\xbe\x66\x11\x50" +
	"\x3d\x4a\x25\x24\x1e\x04\xe5\x1a\x64\x3c\x49\x5a\x21\x51\x4b\x81\x35\x0a\xfe\x83\x74\xde\xf4\xe1\xd4\xba\x61\xff" +
	"\xc9\x46\x85\x5b\x2a\x03\xef\xfd\x48\x2b\xdf\xb6\x97\x76\xea\xef\x59\x2d\x7e\xd7\xaa\x61\x13\x09\xab\xb7\x70\x61" +
	"\x7b\x0a\x04\x27\x9e\x9a\x59\x91\x2d\x4f\xa3\xe9\x0c\x94\xd2\x94\x2c\xb1\x30\xcf\x57\x7e\x1d\xfa\x34\x08\x45\x2d" +
	"\xf1\x95\xc0\x88\x69\xf6\xd2\xe8\x7a\x26\x62\x59\x09\x32\x3e\x94\x29\xcf\xb8\xee\x81\x32\x6d\xcb\x84\x97\xaf\x4e" +
	"\x44\xdd\x13\xc2\x4a\xb0\xe6\x96\x67\xa1\x28\x6f\x0f\x6b\x77\x4b\x33\x9f\xb7\x38\x74\xc7\x2f\x5b\x46\x54\x37\xcd" +
	"\x3b\x6e\x93\x93\x45\xfe\xe7\x2b\xb7\x12\x7f\xa9\xa2\x6f\x92\x47\x73\xff\x13\x76\x90\x7c\xce\x66\x18\x42\xa2\xa8" +
	"\xc6\xba\x06\xa5\x66\x54\x8b\x75\xaa\xba\xde\xac\x8f\x9a\x54\x0c\xa6\x33\xa7\x2b\xcf\x33\x84\x91\x41\x01\x0e\x0f" +
	"\x61\xc3\x30\x70\x77\x22\xad\x6a\xa6\x3b\x20\x4d\xa6\x4f\x04\xa4\x12\x5a\x8b\x3e\x32\xeb\xae\x88\xa9\x70\xa6\xbb" +
	"\x29\xfe\xe0\x9a\x84\xbb\x6b\x92\xbb\x01\x98\xa6\x08\x2a\xe6\x52\xfa\x16\xbe\xf9\xc0\x5f\xc8\x66\xba\x99\x66\xee" +
	"\x9f\x4e\x28\x37\x03\x57\x70\x15\x82\xfe\x12\x1b\xae\xc1\x76\xd6\xf5\x65\xe3\xf8\xbf\x7f\x02\x00\x00\xff\xff\x6c" +
	"\x72\x3c\x50\x3b\x10\x00\x00")

func bindataCssPortunuscssBytes() ([]byte, error) {
	return bindataRead(
		_bindataCssPortunuscss,
		"css/portunus.css",
	)
}



func bindataCssPortunuscss() (*asset, error) {
	bytes, err := bindataCssPortunuscssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "css/portunus.css",
		size: 4155,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1562431114, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"css/portunus.css": bindataCssPortunuscss,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"css": {Func: nil, Children: map[string]*bintree{
		"portunus.css": {Func: bindataCssPortunuscss, Children: map[string]*bintree{}},
	}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
