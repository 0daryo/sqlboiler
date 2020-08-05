// Code generated for package driver by go-bindata DO NOT EDIT. (@generated)
// sources:
// override/templates/17_upsert.go.tpl
// override/templates/singleton/mysql_upsert.go.tpl
// override/templates_test/singleton/mysql_main_test.go.tpl
// override/templates_test/singleton/mysql_suites_test.go.tpl
// override/templates_test/upsert.go.tpl
package driver

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templates17_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x5f\x6f\xdb\xc8\x11\x7f\x96\x3e\xc5\x9c\x90\xbb\xa3\x0a\x86\x49\x81\xa2\x0f\x2e\xfc\x10\xff\x49\xce\x8d\x9d\xb3\xad\xb8\x06\x6a\x18\xc1\x9a\x1c\xca\x0b\xaf\x76\x99\xe5\xd2\xb6\xca\xf2\xbb\x17\x33\x4b\x8a\xa4\x2c\xc9\x4a\x1a\x1f\xee\xc9\x22\x77\x76\xe6\xb7\xf3\x9b\x7f\x4b\x97\xe5\x6b\x78\x25\x94\x14\x39\xec\xec\x42\xf4\x8e\x7e\x61\x1e\x7d\x16\x37\x0a\xc1\xff\x89\x3e\x89\x19\x56\xd5\x90\x45\xf3\xf8\x16\x67\xc2\x2f\xd3\x86\x56\x02\xfe\x0b\xd1\xa4\x5d\xe5\x0d\x32\x85\xe8\x5d\x92\x7c\x50\xe6\x46\x28\x78\x5d\x55\xc3\x37\x6f\xe0\x22\xcb\xd1\xba\x0f\x20\x9c\xc3\x59\xe6\x72\x10\x1a\xa4\xa6\x77\x21\x08\x9d\x40\x62\x90\xdf\x15\x59\x22\x1c\x82\xb1\x20\xa7\xda\x58\x04\xa3\x21\x36\x3a\x55\x32\x76\xd1\x30\x2d\x74\x0c\x81\x81\xbf\x94\xa5\xc7\x1f\x5d\x64\x13\xa9\xa7\x85\x12\xb6\xaa\xc6\x8d\x95\x80\x41\x68\xe3\x20\xfa\x64\xf6\x8d\x76\xf8\xe8\xaa\x2a\x76\x8f\xa4\x8a\x1e\xa2\xfa\x65\x08\x65\x89\x3a\x21\x90\xb5\xe5\x7d\xa3\x8a\x99\xce\xc3\x1a\x5c\xfd\x08\x37\x46\xaa\xa8\x7e\x18\x03\x5a\x6b\x2c\x94\xc3\x81\x45\x57\x58\x0d\x26\xf2\x86\xbd\xdd\xae\x4d\xde\xf7\x01\xdd\xc1\x5e\x30\x2e\x4b\x54\x39\x32\x8e\x10\x9a\x85\x5a\xb2\x5e\xd7\x49\x55\x85\x1b\x91\x8c\x87\xd5\x70\xb8\x00\x3d\xf4\xee\x26\x07\x76\x5c\x4e\x3f\x4f\x85\x96\xf1\x92\xf3\x4f\xff\x3f\xef\x03\xeb\xcc\xe9\x1d\x3b\x60\x6b\x3a\x4e\x5f\x9a\x8f\x72\x38\x90\x29\x81\xa2\xe8\xfc\x23\xc9\xf8\x07\x1b\xfd\x69\x17\xb4\x54\x84\x62\x90\x91\x8b\x02\xd6\x77\x69\x45\x76\x68\x6d\x80\xd6\x8e\xc7\xc3\x41\xb5\x8a\xb8\x35\x4c\xad\x22\x0a\x8a\x5c\xea\x29\x3d\xe3\x23\xc6\x85\x33\xf6\x5b\x12\xa7\xa3\x3a\xfb\x3e\x16\x4f\x9f\xfa\x93\x80\x78\xdf\x1d\xd6\x90\x3a\x5e\x7d\x4a\x6d\x2b\x5e\xbf\xea\xec\x7a\xde\xd7\xdb\x53\xbe\x22\xce\xba\x71\x45\x30\x5e\x8e\xd6\x7b\x61\x61\x36\x9f\x9c\x1d\xaf\x74\xe6\x85\x96\x5f\x8b\xc6\x2a\xec\xc2\xd5\x75\xee\xac\xd4\xd3\x92\xeb\xac\x15\x7a\x8a\xf0\x4a\x86\xf0\x2a\x36\xaa\x53\x69\x9b\x0d\x64\x61\x40\x92\x32\x65\x91\xc8\xeb\xa3\xb7\xa3\xb2\xe4\x37\xbe\x6c\x8f\x42\x2f\xd7\xc0\xaa\x7f\x57\x8c\x76\x11\x0b\x2f\x11\x65\x13\xc4\x1e\x53\x90\x98\xb8\x98\xa1\x76\xc2\x49\xa3\x21\x35\x16\x6e\xcd\x03\x38\x03\x99\x35\x19\x5a\x35\x87\x22\xc7\x3e\x1d\x6c\xb1\xc7\xc8\xb6\x41\xfa\xe7\x8a\xd1\x45\x9b\x90\x29\x18\xd8\x6d\xc3\xa9\x6e\x1b\xbc\x9e\x47\x9f\xf0\x21\x18\x95\x65\x74\x7a\x37\xf5\xec\xed\x80\x36\x50\x96\xbd\x46\x4c\xee\xba\x97\x09\x26\xec\xc2\x82\x4f\x3b\xe2\xf8\xf3\x4c\x13\x91\x8a\xa8\x19\x39\x39\xc3\xdc\x89\x59\xf6\xc5\x4b\x7d\xb9\x45\x95\xa1\x1d\x41\x04\x95\x97\x6e\x73\xe4\x37\x63\xee\xea\xb0\xea\x66\x53\x62\xf6\x30\x35\x16\xbd\x53\x59\x68\xeb\xd4\x7a\x9a\x3c\xed\x69\x09\xee\xa0\x8d\xc5\xe1\x40\xff\xe7\x00\x53\x51\x28\xc7\x83\xc8\xd7\x02\xad\xc4\x3c\xfa\x64\xf4\xbf\xd1\x9a\x7a\x69\x82\x44\x6b\x4d\xfa\x81\x79\xd0\x2d\xed\xb5\xa7\x2f\xa5\xbb\xad\x85\x43\x30\x63\x52\xeb\x13\xe3\x19\xad\x5b\xe6\x29\xeb\x64\x07\x29\xd4\xc1\x42\xf7\x98\x18\x7d\xbb\x8e\xcf\x58\x68\x72\x96\xa7\x00\x1e\xa4\xbb\x05\x01\x8e\x27\x28\x77\x2b\x1c\xd4\xeb\x4d\xee\x50\x1e\x09\x28\x58\x33\xc4\x6c\xb7\x61\xf7\xcd\x1b\xd8\x2b\xa4\x4a\x20\x16\xf1\x2d\xc2\x1d\xce\x41\xea\xd7\x4a\x6a\x84\x62\xaa\xa4\x9a\xc3\x6b\x98\xcd\xf3\xaf\x0a\xee\x73\xc8\xe8\x6f\x66\xcd\x8d\xc2\x59\x3e\x1c\xdc\x14\x29\xb9\x20\x77\x76\x26\xf4\x54\x21\x35\xb9\xbd\x22\x4d\xd1\x06\x63\x5e\x8d\x2e\xad\x74\x38\xe1\x22\x14\xe4\xce\xc6\x46\xdf\x47\x47\xce\x88\xa0\x17\xe7\xd1\x47\xa9\x13\x2a\x77\x14\x7c\x5f\x42\x88\x49\xab\x2f\x57\x7d\xb9\x7d\xa3\x72\x76\xc9\xb2\xee\x98\x4f\xd3\xbe\xde\x9b\x3b\x0c\x7e\x8d\x7e\x7d\x0e\x46\xbf\x0c\xac\x87\xd1\x97\xfb\x1e\x18\x4f\x75\x76\xa2\xf3\x07\xe8\x6a\x42\x72\x83\x2a\xe2\x76\x67\x17\x68\xb5\x5e\x18\x0f\x07\x2d\x79\xa7\x45\x43\xde\x4d\x91\x8e\x39\x95\x57\xa6\x85\x4f\xdb\x7d\x0a\x97\x93\xc2\x45\xe7\xc7\x26\xbe\x23\x4d\x1c\x40\xa1\x8f\xa3\x84\x0c\x3d\xbf\xff\xea\x0e\xe7\xd7\x5b\x1b\xba\xd0\xca\x9b\x1a\x0e\xa8\x0f\x52\x1d\xe0\x9c\xf0\xd9\xf3\x53\x6d\x98\x1c\xd0\x0c\x9f\x16\x1d\x01\xe9\xb3\x77\xd4\x79\xa2\x3c\x1d\x0e\x06\xeb\x10\xbc\x53\xaa\xc9\xd2\x0d\x52\x2b\xea\xc4\x76\xd2\xa6\x70\xdd\x0d\x6d\x40\xd0\xe3\x78\x38\x18\xd4\xfd\x70\x67\x77\x29\x0f\x2e\x3a\x4f\x3f\xe4\x08\xa7\x56\xce\x84\x9d\x7f\xc4\x79\x47\x98\x1c\xdd\xd4\x25\x6f\xbf\x53\x94\x9e\xef\x32\x85\xf6\xf5\xc8\x34\x65\x6a\xa9\xe7\x84\x10\x9b\x42\x25\x5c\xf5\x6f\xb8\x04\xd5\xc7\xf5\x05\x0a\x94\xcc\xb9\x07\x71\x99\x22\x73\xd0\x2d\x35\x13\x9a\xa7\x67\x99\x42\xea\xfe\x81\x45\x17\xb6\x49\x40\x9b\x38\x1a\x22\xaa\xce\x73\xd8\xf5\xfa\x7d\x3c\x9d\xd1\xab\x13\xaa\xcd\x41\x22\x85\xc2\xd8\x85\x40\xe3\x4d\xde\xbd\x67\x8e\x9a\x4e\xdc\xb4\xe0\x56\xa5\x45\xaf\x02\x76\x21\x9d\xb9\x68\x92\x59\xa9\x5d\xca\x14\x8c\x26\x87\xc7\x87\xfb\x9f\xe1\xe7\x1c\xde\x9f\xff\x7e\x42\x07\x3e\x3e\xab\xaa\xa5\x73\x97\x65\x74\x7e\x56\x55\x70\xf9\xdb\xe1\xf9\x21\xfc\x9c\x8f\x98\x18\x3f\xa9\xe5\xd1\x3f\x8d\xd4\x41\x7b\xcc\xa3\x04\xb5\x3b\x2b\x8c\xc3\x89\x92\x31\x36\x90\xa3\xe3\xb3\x10\x9a\xdf\xe7\x67\x1c\xe9\xe3\x10\x46\xe1\x68\xdc\x68\xab\x15\x5c\xde\xa2\xc5\x7d\x25\x8a\x1c\x99\x20\x02\x34\xe2\x13\x33\x8a\x51\x08\x6f\xbb\x9e\x5b\xf0\xee\x0f\x7b\x2f\x54\x81\x27\x22\xcb\xa4\x9e\x86\x9c\x71\x6d\xc7\xdb\x93\x3a\xa9\x97\xd6\x75\xd0\xcf\xf3\x0c\xc3\x75\x75\x60\xa1\xb6\xf5\x70\x3d\x25\x74\xba\x7b\xaf\xbd\x53\x11\x6b\x02\x92\x0e\x4c\x82\x75\x34\x2e\xb8\x79\x69\xb0\x64\x97\x0c\xae\x80\xda\xc7\xca\x60\x2b\xdf\x64\xd9\x8d\x5c\xad\x31\x65\xca\x8e\x74\x22\x2d\xc6\x14\xb8\xfe\xc5\xbf\x48\xe2\xf7\x34\x30\xd4\x7f\xee\x85\xea\xcd\x16\xbc\x98\xbf\xb7\x66\xd6\x1c\x81\x15\xd6\xb5\xb6\x47\xd2\xd8\xd7\x46\x8f\x24\x87\xab\x6b\xa9\x1d\xda\x54\xc4\x58\x56\x8b\x21\x63\xd9\x59\x1d\x47\x36\x1b\x5b\xe3\xa7\xce\xae\x37\xdd\xd1\xd1\x0c\x8b\xbd\x09\x79\x31\xfc\xf1\xe8\x7a\x80\x37\xc5\xf4\xc4\x24\xc8\xa6\x28\x7b\xde\x73\xf6\x28\x1d\xb4\xeb\xdc\xb8\x6c\x63\x80\x13\x78\xfc\xbc\x34\xb9\x6c\x5c\x0f\x80\x34\x80\xf7\x0d\x1f\xe5\x2c\x1c\xc4\xee\x91\xef\x76\x83\x07\xde\xc6\xdd\x70\x49\x15\x1d\x95\xe5\x96\x6d\x3e\x6c\x81\xeb\x61\x15\x9a\xe6\xee\x46\xb5\x37\x16\xfa\x58\xe4\xce\xb7\xa0\xa3\x83\xee\x25\x6c\x69\xa5\xbe\x8c\xf1\x55\x6c\xd5\xd2\x6a\x4f\x5b\xcc\x79\x4c\xad\x67\x6d\x9a\x98\xf9\x66\x12\x74\x50\x7b\x78\x51\x14\x8d\x59\x4b\xeb\xad\x75\x9b\x6b\x0b\x01\x8f\xe3\x1b\x14\xd5\x07\xed\xe9\x5c\x0d\xf3\x4b\x93\x9e\xdf\x06\xf0\xe9\xb6\x6f\x87\xd6\xdc\x0e\x56\x24\x70\xbf\xa3\xd1\x4d\x9c\xae\xe1\xbe\x56\x6e\xea\x6b\x34\x8e\x2d\xd5\xf8\x51\x37\x1d\x56\x13\x48\x69\xaa\xe8\xed\x01\x48\xed\xfe\xfe\xb7\x1e\x38\x5a\xf4\xe3\xfa\x89\xc8\xe0\xea\xba\xa8\x45\xe8\x7d\x53\xac\x79\x0a\xed\x27\xf8\x86\x0c\x5f\x34\xee\xa9\x71\x06\x78\xa8\xaa\x2f\x68\xcf\x22\xf5\x28\x1b\xdf\xfb\x28\x89\x3a\x62\x09\x4d\x7f\x6b\xdd\x79\x68\xed\x64\xae\xe3\xf7\x42\xaa\xc6\xd2\xab\xd8\x28\xfe\xbc\xcb\xb3\x59\x82\x8f\x4d\x12\x9c\x7e\xc4\xf9\xe2\x6a\xff\xb6\xa5\x6c\xe9\x83\x05\x7f\x4b\xe3\x19\x61\xa1\xa9\x27\xfa\x59\x3a\xe5\x47\xd0\xba\x96\x2f\x49\x93\xac\x89\x3c\x0e\x2f\x5b\x55\xc0\xf3\x6a\x6c\x54\x44\x7d\xa0\xaa\x02\x7f\x6a\x7f\xb2\x9a\x27\xae\x92\xbf\xfc\xb2\xde\xc3\x7f\xa5\xd5\xe5\x95\xab\xb7\xd7\xb4\xb6\xb9\xb1\x5c\x8d\x5a\xb7\x54\xd5\xe8\x7a\x3d\x55\xbd\x1b\xee\x22\x46\x5e\xac\xdf\x75\x87\xaa\x1f\x90\x32\x16\x9d\x95\x78\x8f\xcd\x65\x94\xbb\x49\xbe\x26\x85\x80\x8e\xdb\x0b\xf7\x4d\x3d\x71\x9b\xde\x1a\xb6\x59\x35\x7e\x81\x6e\xd5\xcc\x86\x5b\x34\xac\xee\xb1\x7c\x9d\xfa\xc3\x7a\xd7\x5a\x94\x0f\xcf\x60\xeb\x74\xb2\x15\x7e\xeb\x94\x66\x56\x7f\x6e\x1e\x82\xbe\xc1\xa7\x9a\xa3\x49\x2c\x78\x96\xa3\xa1\xc3\x9b\xea\xfa\x60\x85\xca\x15\x15\xff\x5b\xd5\x37\xcd\xe0\x07\x84\x73\x66\xb2\x82\xbf\x8b\x25\xfe\xe6\xb6\x39\x9e\xa9\xfc\x75\xd3\x79\xe7\xc9\xc5\x75\xbb\x9b\x70\x73\xe3\xde\x42\x9c\x6f\xd8\xb0\xeb\x3d\xb5\xb5\x81\xc5\x4d\x7b\xb0\xe1\x93\xde\xe2\xbf\x53\x89\x79\x97\x3a\xb4\xdf\xf5\x39\xaf\x2e\x67\x9d\x19\x84\x95\x6a\x6a\x16\xdd\xcf\xca\xff\x0b\x00\x00\xff\xff\x4c\x12\x80\x9a\x54\x1c\x00\x00")

func templates17_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates17_upsertGoTpl,
		"templates/17_upsert.go.tpl",
	)
}

func templates17_upsertGoTpl() (*asset, error) {
	bytes, err := templates17_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/17_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSingletonMysql_upsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\x8e\xd3\x30\x10\xc7\xcf\xf6\x53\x0c\x91\x56\x8d\x25\x2b\xcb\x5e\x91\x7a\xd8\xa5\x65\x15\x28\xfd\x2e\x08\x21\x0e\x6e\x3d\x6e\x2d\xa5\x4e\xf1\x47\xa1\x42\x7d\x77\xe4\x24\x6d\xb3\x4b\x41\x1c\xf6\x92\x8c\x3d\x33\x7f\xcf\x6f\x66\x6e\x6f\x61\x19\x74\x21\x17\x3b\x87\xd6\x4f\x02\xda\xc3\xc7\xc3\x6c\x32\xa8\x6f\x1d\x08\x88\x07\xe7\x85\xc7\x2d\x1a\x0f\xce\x5b\x6d\xd6\x10\x5c\xfc\xfa\x0d\x42\xa8\x12\x7b\xc2\x0b\xd8\xd9\x72\xaf\x25\xca\x8c\xaa\x60\x56\xd7\x75\x53\xa9\x05\x48\xab\xf7\x68\x5d\xd6\xd3\xa2\xc0\x95\xe7\xe0\xc5\xb2\xc0\xa1\xd8\x62\xa3\xcf\x21\xec\xa4\xf0\xc8\xe1\xc7\x46\x7b\x2c\xb4\xf3\xf0\xf5\x5b\xed\x63\xa7\x1a\x7e\x51\x72\xf1\x76\xe3\xed\x56\x98\x75\x81\x59\x2e\xd1\xf8\x49\x28\x3d\xce\x0a\xbd\xc2\xf8\x64\x36\x98\x70\x88\xff\xe9\xa4\xa5\xc9\x28\xb9\xbc\x7c\x5d\xe1\x8f\xe4\x73\x02\xa3\x94\x2c\x83\x82\x37\xed\xc4\x47\xf4\x0f\x41\x29\xb4\x29\xa3\x44\xa2\x42\xdb\x72\x8e\xc3\xc9\xb9\x0c\x2a\xa6\xef\x85\x85\x55\x59\x84\xad\x71\x0d\x14\x25\x5a\x41\x81\x26\xbd\xd4\x08\xaf\xba\xf0\x3a\xc2\x92\x53\x68\xb7\x09\x76\xd9\xfb\x52\xb7\x42\x39\x24\x3c\x61\x94\x1c\xe9\x59\xa6\x6e\x23\x83\xee\x49\x43\x6d\x7d\xf6\x6e\x67\xb5\xf1\x2a\xa5\x84\x44\x02\x1e\xff\x49\x3e\x9c\xf5\xa7\x73\xc8\x1f\x87\xa3\x69\x1f\xf2\xe1\x7c\x04\x37\x0e\xd2\x1b\xc7\xe0\xd3\xfd\x60\xd1\x9f\x55\x76\x52\x05\x9f\x7b\x50\x9d\x9a\xb2\x2a\xbb\x05\x5b\x88\x15\x6e\xca\x42\xa2\x75\x55\x13\x17\x0e\x73\x23\xf1\x67\xdb\xc1\x9f\xb1\x72\xb8\xe3\x70\xc7\xa2\x14\xa3\x84\x58\xf4\xc1\x1a\x58\x06\x95\xcd\x2a\xe2\xb4\xa1\x7b\x46\xd1\x40\x9c\x19\xfe\x52\x3c\x8c\x86\xd0\x5b\x8c\x07\xf9\xdb\xfb\x79\x1f\x3e\xf4\xbf\xc0\x62\xdc\x8b\x66\x45\xf5\x04\xaa\xc5\xf4\x62\x48\x71\xe2\xaa\xb4\xa0\x39\xec\xe3\xd6\x58\x61\xd6\xd8\x2c\x7a\x35\x1b\xad\x40\x5f\xa6\x1d\xa9\xb2\xcf\x56\x7b\x7c\x38\x78\x4c\x3b\xbc\x13\x5b\x72\xa4\x84\x7c\x8f\x8b\x29\x9f\x2e\xde\x3f\x36\x76\xcf\x68\x4b\xac\x69\x64\xad\x71\xcd\x93\x40\xb7\x69\x5a\x9a\xfc\x67\x66\x5d\x20\xeb\x34\xd3\xb9\x36\xb6\x23\xfd\x1d\x00\x00\xff\xff\x4c\x0d\x4e\x35\x6a\x04\x00\x00")

func templatesSingletonMysql_upsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingletonMysql_upsertGoTpl,
		"templates/singleton/mysql_upsert.go.tpl",
	)
}

func templatesSingletonMysql_upsertGoTpl() (*asset, error) {
	bytes, err := templatesSingletonMysql_upsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/singleton/mysql_upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_main_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5b\x6f\xdc\x36\x16\x7e\x96\x7e\xc5\x89\x80\x74\xa5\x54\x66\x02\x14\xd8\x07\x07\x42\x10\x8f\xc7\x85\xd1\xc4\x97\x19\xef\x16\x8b\xba\x68\x69\x89\x63\x13\x96\x48\x99\xa4\x3c\x9e\x35\xfc\xdf\x17\x87\xd4\x85\x1a\x8f\xbc\xce\x6e\x1e\xfb\x94\x48\xfc\xce\xed\x3b\x47\x87\x9f\xe7\x9e\x2a\x50\xd7\x0f\x5f\x37\xcb\xf3\x2f\xb7\x6c\x03\x19\x28\x76\xcd\x1e\x6a\xf2\xb5\xd1\x66\x26\xab\x9a\x97\x2c\xfe\x33\xfe\x54\x25\x71\x9c\x5e\x8a\xe4\xd3\xa5\xfe\x71\x76\x7a\xb2\xbc\x58\x7c\x3e\x3e\xb9\x20\xef\x3e\x1d\x9d\x2e\xe6\xc7\x3f\x9f\xc0\x2f\xf3\x7f\x91\x77\x9f\x2e\x45\xf2\xe3\x9f\x49\x18\x9a\x4d\xcd\xa0\xda\xe8\xbb\xf2\x82\x69\xc3\x14\x68\xa3\x9a\xdc\xc0\x63\x18\x14\x57\x33\x29\x04\xbc\xd3\x77\x25\x39\x3c\x08\xf1\xc5\x09\xad\x18\x20\x84\x8b\xeb\x30\xb8\x91\xda\x00\x0c\xcf\x8d\x66\xca\x7f\xae\xa9\xd6\xfe\xb3\xd6\x65\x25\x0b\x36\x9c\x4b\x65\xed\xb9\x30\x61\x18\xc8\xda\x70\x29\x8e\x78\xd9\x03\xc2\xc0\x30\x6d\x0e\x0f\x6c\xd4\xde\xc9\x2d\xaf\x97\xe7\x5f\x66\x55\x01\x57\x52\x96\xe1\x53\x18\xae\x1a\x91\x03\x17\xdc\xc4\x89\xcb\xfb\x2b\xe5\x02\x32\xf8\xc1\xab\xeb\xf1\xa9\x47\xc6\x15\xbc\xf3\x4e\x12\xd0\xcc\x34\x75\x9c\x00\x53\x4a\x2a\xf4\x80\x5c\x33\xa5\xdc\x8b\x30\x0c\xee\x79\xcd\x14\x59\x32\x73\xc8\x56\xb4\x29\x4d\x1c\x59\x7b\xd2\x16\x14\xa5\x10\x19\xd5\xb0\x28\x99\x86\x62\xad\x51\x0a\x3f\xfd\xf4\xe1\xef\x49\x18\x06\x15\x69\xc9\xcc\xc0\x59\xfc\xcc\xcc\xd2\x56\xd8\x19\x14\x57\x82\x56\xd6\x65\x45\x2c\xd1\x93\x48\x3c\x75\x38\xdb\x80\x49\x1c\x9e\x3a\x9c\x6d\xcc\x24\x0e\x4f\x5b\x1c\x36\xc8\xc3\x1d\x8b\x71\x3d\x16\xd4\x75\x75\xd2\x5f\xc7\x92\x45\x7b\x1d\x9d\x34\x40\x8c\x5f\xbe\xd7\x72\xcf\xe6\x40\xca\xb2\x0f\x71\xcb\x6b\x7d\x57\xe6\x55\x11\x21\xbb\xd8\xbb\x0c\xee\x69\x49\xc9\x01\xbb\xe6\xe2\x9f\xb4\xe4\x05\xc5\xf1\x8a\x13\xd2\x3e\xb0\x38\x0c\x02\x0b\x71\xc1\x4f\xa4\x99\x57\xb5\xd9\xc4\x8e\xc6\x14\x46\xac\xa5\x93\x60\x64\xbf\x07\xbb\x56\xf4\xe0\x13\x69\x62\xfb\x9f\xf9\x5d\x43\x4b\x1d\x3b\x46\x53\xf8\xd0\x1b\x38\x1a\x5f\x70\xef\xc6\xa4\xc7\x77\xb4\x4c\x1b\xb4\x6c\xf7\x16\x3d\xfb\x69\x18\x24\x64\x76\xc3\xf2\xdb\x18\x39\xe2\x2b\x3b\xe2\x6f\x32\x10\xbc\xc4\xa1\x0f\x14\x33\x8d\x12\xf8\x36\x0c\x9e\xc2\x30\x78\xff\x1e\x66\x8a\x51\xc3\x80\x82\xa2\xa2\x90\x15\xff\x37\x2b\xa0\xb8\x02\xcc\x81\x58\x17\x25\x13\xb1\xdf\xd4\x04\xb2\x0c\x3e\x58\x77\x5b\xbd\xee\x3d\x90\xa5\xa1\x57\x25\x73\x07\x7d\x85\x89\x8b\xd9\x66\x95\x41\x45\x2a\x7a\xcb\x4e\xfb\x9d\x10\x27\x1f\xa7\xf3\x95\x4a\x93\x5f\x15\xad\x63\xa6\xb0\x71\xb9\x6c\xca\x42\xfc\xcd\x00\xba\x00\xb7\x57\x60\xc5\x4b\x3b\x4e\x6d\x94\x37\xa3\xb1\x42\x77\x5e\xe8\x42\xc9\xfa\xc2\x26\xbf\x23\xec\x88\xa7\xe0\x69\x6c\x99\x5b\xc2\x5e\x6d\x1b\x06\x41\xd1\x54\x35\xa6\xb0\x9f\x01\x7b\x60\x39\x99\xc9\xaa\xa2\xa2\x68\x27\x1b\x4f\xa3\x14\x53\x72\xeb\x44\x3b\x2e\x52\x88\xf6\xf6\x84\xdc\x2b\xa8\xa1\xee\xb8\x23\x31\x70\x19\x4c\x7b\x9c\xf2\x86\xae\xae\xa8\x66\xf6\xdc\x6b\x28\xe6\xa8\x52\x58\xa3\x3b\x2e\xc9\x19\xaf\x59\x9c\x0c\x79\x2f\x4d\x81\x35\xee\x67\xf0\xc3\xd5\xc6\x30\x4d\x0e\x9a\xd5\xca\xae\x5b\x2f\x95\x69\xd0\xe0\x88\x2c\x4d\x21\x1b\x5c\x37\xeb\xf1\x4b\x47\xed\x28\x5c\xe8\x3b\x47\x8c\x5d\xf7\x82\xad\x8f\x7e\x61\x9b\x43\xa6\x8d\x92\x1b\xa6\x62\xef\xba\x4c\x41\x25\xdb\x46\xce\xf1\x56\x92\xa1\xdf\xcf\x21\x0b\xaa\xcc\xcb\xed\xdc\x1a\xc1\x15\xe5\x25\x2b\xc0\x48\xd0\x68\x0b\x7d\x33\x21\x77\xdd\xc0\x51\x1c\x0f\x8f\x9f\xdb\x77\x09\xb7\x15\x6a\x57\x61\xbf\x52\xbe\x33\xd0\xaa\x32\xe4\x4c\x71\x61\x4a\x81\x11\x92\xed\x77\xa3\x6e\xb4\x3b\x28\x4e\x92\x57\xe6\xb8\xa6\xdc\xc0\x4a\xaa\x49\x56\xc2\x20\xf8\x03\x07\x81\xcc\x4a\xa9\x59\x9c\xc0\xfb\xf7\xf0\x79\x85\xea\xa4\xfb\x5a\xb8\x86\x42\x0a\x96\x42\x8e\x08\x30\x37\x0c\xd6\x8a\x1b\x06\x4c\x14\x20\x57\xf6\x45\xcd\x6b\x16\xee\x66\xf8\x7f\xad\x7b\x6b\x58\xfe\xcf\xca\xb7\x67\x01\x0b\x6f\x9d\x08\x5e\xbe\xa0\x57\x74\xf9\x55\x16\x2c\xf6\xc4\x54\xd2\xfe\x8b\x65\xe8\x35\x37\xf9\x0d\xd8\xd3\xc7\x30\xc8\xa9\x66\xad\x3e\xd9\x1f\xb6\x66\xb4\x98\x9f\xff\xe3\x78\x31\x3f\x8c\x3a\xc4\x8a\x96\x7a\x0c\x39\x3c\x5e\x7e\x3e\xf8\xe2\x41\xce\x16\xf3\xa3\xf9\x02\x8d\x7c\x58\x14\x06\xed\x3e\xf1\xde\x62\x74\xac\x68\xb2\x88\xf1\x0a\xf2\xd2\x6f\x1d\x20\xed\xcb\x1a\x79\x5f\xc5\xb8\x9e\x5a\xf8\x1e\xee\xf1\xec\xad\xb6\x6b\x6a\x90\x8c\xc9\x74\xa0\xed\x7b\x64\x90\x79\xa6\xaa\x53\x68\x17\x13\x97\x8d\xe1\x25\xb9\x60\x55\x6d\x61\x11\x8a\x3a\xe7\xbf\xbb\x39\x5e\xba\x31\x27\x3b\xee\x26\x66\xe7\x25\xa4\x2f\x66\x67\x18\xda\x12\x1f\x06\x7f\xa4\xed\x98\x4a\x8d\x1b\xc0\xb4\xda\xc2\x05\x96\x9a\x1c\x6b\xbc\xe5\x1f\xb8\x36\x76\x34\xdd\x9d\x65\x7d\x64\x80\xdd\x0d\x83\x27\x60\xa5\x66\xf0\x0d\x79\xda\x9b\x12\x84\x34\xb8\x37\x0c\x54\xbd\x96\xc4\x04\xb1\x03\x47\x75\x3b\xf9\x96\xab\xe8\xb7\xbc\xe4\x4c\x98\xdf\x11\x32\x1c\xaf\xda\x53\x34\xce\xde\xea\x4b\x61\x9b\xd3\x26\xff\x1c\x86\x9a\x27\x7b\x5b\xb4\x30\x7c\xda\x09\x43\xe1\x35\x78\xc3\xa7\xc4\x93\x1c\x28\x52\x13\xac\xd1\x89\x8d\x1d\x51\xa8\xd6\x6b\xa9\x8a\xc1\x85\x35\xc1\xd2\xd0\x8b\xfd\x3a\xf6\xf1\xca\xee\xbf\xa6\x4e\x2a\x25\x1f\xdd\xe9\x9b\x0c\xa2\x68\xc2\xbb\xd6\xe5\x1e\x82\x7a\xef\x68\xd7\xf9\x76\x5d\x19\x1b\xf6\x14\xd6\x4a\x1a\x99\xcb\x32\x33\x79\xfd\x12\xd3\xfd\x6e\xfc\x8b\xec\xef\x4b\xb6\xbf\x36\xf0\xd3\xa9\x6a\x62\x95\x68\x32\x6c\x5f\x7c\xd7\x5e\x3d\xd3\x7b\x65\x2c\xf5\x86\xad\x82\x8b\x1d\xbf\x6a\x7f\x7f\xb5\x5b\xa0\xd3\x58\xf0\x56\x7f\x7c\xa6\xb3\xfa\xcd\x49\x54\x23\x66\x55\x11\xeb\xbb\xb2\x53\xf1\xd1\x0b\x79\xf8\x62\xf5\xe5\x2c\x10\x39\xe4\x80\x6b\x02\xb7\x89\xfe\xae\xd9\x18\x46\x55\x21\xd7\xc2\xcf\x05\x27\x80\xb4\xbf\x26\x3c\xdf\x4a\xdd\x51\xcf\xf8\x7f\x95\xe8\xfb\xdf\xae\xd1\xbd\xab\x55\x6a\xb2\x60\x95\xbc\xc7\x21\x7c\xd5\x05\xd2\x11\x80\x32\x33\xed\xee\xec\xf6\xc2\x4a\x81\xaa\x6b\x0d\x84\x90\xee\x1e\xee\xab\xb6\x07\x19\xd0\xba\x66\xa2\x88\x7f\xfb\xdd\x01\x1e\xb7\xc5\xf7\x93\x73\x41\x08\xc1\x01\xcc\x77\xe8\xf6\x36\xa2\x87\x43\x58\x2f\x7b\x9d\x5f\x4d\x4e\xd8\x7a\xc1\x68\xc1\x94\xcb\x14\xbd\x69\x27\xa9\x77\x89\x73\x3d\xad\xdb\x73\x5f\x8c\x3b\x17\xfd\x4b\x77\x43\x39\xe3\xd0\xeb\x07\x1e\x2f\x1a\xf1\xbc\x15\xbe\x7a\xea\xae\x45\xd5\x08\xc1\xc5\xf5\x7e\xd4\xb3\xe9\x6a\x4b\xc6\x70\x17\xda\xd7\x58\x5b\xa7\x5b\x0a\x6c\xfb\xef\xd7\xd7\x48\xa9\x5c\x0a\x1c\xd5\xb8\xfd\x91\x2b\x75\xed\x4b\xa6\xa7\x76\x6b\x68\x53\xeb\xde\x86\x1b\xff\x68\x14\x0c\x88\x96\xb3\xbb\x92\x9c\xd6\x4c\x0c\x7f\x86\x15\x8a\xdf\x33\x45\xec\x9f\x28\x07\x0d\x2f\x8b\xf3\x86\xa9\x4d\x5b\x50\xf7\x2b\x84\xdb\xa4\xe3\xaf\xb3\x5b\xf8\xdd\x46\x4f\x61\x58\xa7\xbb\x74\xca\x40\x44\xfa\x8c\x9d\x71\x21\x4f\xe1\x7f\x02\x00\x00\xff\xff\x2f\xbf\x35\x64\x67\x14\x00\x00")

func templates_testSingletonMysql_main_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_main_testGoTpl,
		"templates_test/singleton/mysql_main_test.go.tpl",
	)
}

func templates_testSingletonMysql_main_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_main_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_main_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testSingletonMysql_suites_testGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x0a\xc2\x30\x10\x44\xef\xfd\x8a\xa5\xe4\xd0\x4a\x9b\x0f\x10\x3c\x78\xd4\x83\x88\xb4\x1f\x10\xed\xb6\x04\xe2\x5a\xba\x5b\x10\x42\xfe\x5d\xd2\x46\xe9\xc1\xdb\x0c\x6f\x32\x99\xed\x67\x7a\x40\x83\x2c\xed\xc8\x38\x49\x21\xb0\x13\x64\xb1\x34\xe8\xa6\x04\x9f\x01\x78\x5f\xc3\x64\x68\x40\x50\x96\x3a\x7c\x57\xa0\xc4\xdc\x1d\xc2\xfe\x00\xba\x89\x8a\x43\x48\x39\xdb\x27\xa8\x4f\x7c\x7e\x59\x5a\x30\xd4\x3f\x8e\x8e\xb7\x56\x19\x67\x0d\xc7\x22\xa5\x8f\x51\x22\xaf\x8d\xdf\x96\x8b\x79\xe2\x92\x16\x7d\x9b\xa9\xc8\xbd\x5f\x9f\xe8\x76\xbc\xba\x79\x32\x2e\x84\xbc\x82\x38\xf8\x0f\x59\x2f\x2a\x97\xbf\x90\xba\xed\x8c\xe4\x42\xf6\x09\x00\x00\xff\xff\x11\x5d\x4c\xce\xff\x00\x00\x00")

func templates_testSingletonMysql_suites_testGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonMysql_suites_testGoTpl,
		"templates_test/singleton/mysql_suites_test.go.tpl",
	)
}

func templates_testSingletonMysql_suites_testGoTpl() (*asset, error) {
	bytes, err := templates_testSingletonMysql_suites_testGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/mysql_suites_test.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_testUpsertGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x54\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\xd8\x5d\x50\x0b\x85\xd9\x5e\x53\xf8\xe0\x7c\x1c\x82\xb6\x86\x1b\xcb\xe7\x82\x91\x46\x0e\x61\x9a\x54\xc9\x51\x6d\x57\xe0\x7f\x2f\x28\xd9\x8e\x13\x3b\x6d\x0e\xed\x21\x07\x7f\x70\xf0\x66\xde\x7b\x33\x1c\xb6\xed\x19\xfc\x2d\xb5\x92\x1e\x2e\x86\x20\x46\xf1\x1f\x7a\x91\xcb\x7b\x8d\xd0\xff\x88\xb1\x5c\x62\x08\xac\x6a\x4c\x01\x84\x9e\xda\xb6\xcf\x10\xb3\x7a\xa2\x1b\x27\x75\x08\xb3\xda\xa3\x23\x4e\xf0\x5f\x04\x28\x33\x17\x79\x0a\x2d\x4b\x48\x4c\xa4\x93\x5a\xa3\xe6\x29\x63\x89\xaa\x40\xa3\xe1\xfb\x02\xd7\x76\x65\xa6\xca\xcc\x1b\x2d\x5d\x08\x23\xad\xaf\xac\x6e\x96\xc6\xa7\x30\x1c\xfe\x0c\x39\x71\x6a\x29\xdd\xe6\x03\x6e\xf6\x09\x2d\x4b\x12\x12\xd3\x85\xaa\xf9\x20\x7e\xd7\xca\xcc\x81\x3a\x1b\x2b\x45\x0f\x60\x8d\xde\x40\xdd\xe7\xc1\x02\x37\x50\xf4\x99\x83\x94\x25\x61\xaf\x6c\xb9\x99\x7e\xfe\x78\xe0\xef\x91\x72\x66\xd4\xd7\x06\x0f\xf5\xfd\xff\x4b\x4e\x63\xa1\xe9\xd2\x76\x64\x40\x16\x0a\x6b\x2a\xad\x0a\x02\x6b\x7a\x6e\x96\x78\xc4\x32\xb6\xdf\x49\x53\xda\xa5\xfa\x8e\x62\x8c\xab\x29\x62\xc9\x53\x96\x7c\x93\x0e\xd0\x75\x1f\xeb\x58\x72\x7e\x0e\x23\x22\x5c\xd6\x04\xf4\x80\x70\x3b\x9e\xde\xdc\xe5\xe0\x55\x89\x60\x2b\x90\x06\x66\x93\x18\x61\x89\x8d\x15\x4f\x5a\x69\x7b\xbf\xb1\xe8\x21\xe7\x94\x5c\x53\x10\x8f\x62\x32\xf8\xd7\x66\xf0\x42\xf3\xaf\x2f\xf3\x4d\x8d\x3e\x83\x4a\x6a\x8f\xe9\xfb\xae\xd0\x5f\x43\x30\x4a\x6f\x3b\x72\x13\xa5\x56\x7c\x30\x33\x5d\x2f\xc8\x3e\xb2\x9c\x56\x04\xbe\xe3\xbe\x80\x7f\xfc\x20\x8b\xf5\xb6\x8d\x69\x5b\x55\x81\xb1\x04\x62\x6c\xaf\xac\x21\x5c\x53\x08\x05\xad\xa3\xb5\xa2\x3f\x8b\x4b\x59\x2c\xe6\xce\x36\xa6\xe4\x69\xdb\xa2\x29\x43\x60\x49\x0f\xf9\xd4\x78\xca\xd7\xbc\xab\x72\x58\xe1\x28\x70\x6f\x95\x16\x97\x38\x57\xa6\xab\xa1\x3d\x1e\xc6\xf2\x35\x2f\x68\x9d\x45\x83\x3b\x86\x57\x81\x52\x96\x94\x58\xa1\x83\xb8\x39\x3c\x85\x16\xbe\xc0\x10\x68\x2d\xee\xac\xd6\xf7\xb2\x58\xf0\x14\x42\x1c\xf1\x7e\x18\x56\x6c\x17\xe9\x25\xe3\x71\x28\x68\x4a\x38\x0b\x01\xe2\xa9\xe3\xbf\x35\x15\x3a\x9e\x3e\x3d\xbd\x6e\x2e\x4d\x47\x77\x7a\x28\x47\xd3\x28\x6c\x63\xa8\x0b\x3c\xbb\x5a\xbb\x57\x80\xa7\xe2\x2a\x62\x5e\x29\xff\xd1\xf9\xb1\x4a\xbe\xa3\x8d\x90\x8e\x38\x82\xde\x3d\x81\x0c\x56\xd2\xc4\x35\x42\x70\x58\x58\x57\x66\x30\xb7\x74\x31\xc8\x7a\xfc\x56\xf4\xb3\x7d\x99\x4d\xae\x47\xf9\xcd\xa9\x7d\xf9\x6d\x1b\xf1\x22\xec\xe8\xd5\x12\x42\xfc\xd1\xf5\x79\x7b\xf7\xea\x8d\x5c\xab\xc0\x7e\x04\x00\x00\xff\xff\x63\xfd\x7b\x9f\x38\x07\x00\x00")

func templates_testUpsertGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testUpsertGoTpl,
		"templates_test/upsert.go.tpl",
	)
}

func templates_testUpsertGoTpl() (*asset, error) {
	bytes, err := templates_testUpsertGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/upsert.go.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/17_upsert.go.tpl":                        templates17_upsertGoTpl,
	"templates/singleton/mysql_upsert.go.tpl":           templatesSingletonMysql_upsertGoTpl,
	"templates_test/singleton/mysql_main_test.go.tpl":   templates_testSingletonMysql_main_testGoTpl,
	"templates_test/singleton/mysql_suites_test.go.tpl": templates_testSingletonMysql_suites_testGoTpl,
	"templates_test/upsert.go.tpl":                      templates_testUpsertGoTpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"17_upsert.go.tpl": &bintree{templates17_upsertGoTpl, map[string]*bintree{}},
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_upsert.go.tpl": &bintree{templatesSingletonMysql_upsertGoTpl, map[string]*bintree{}},
		}},
	}},
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"mysql_main_test.go.tpl":   &bintree{templates_testSingletonMysql_main_testGoTpl, map[string]*bintree{}},
			"mysql_suites_test.go.tpl": &bintree{templates_testSingletonMysql_suites_testGoTpl, map[string]*bintree{}},
		}},
		"upsert.go.tpl": &bintree{templates_testUpsertGoTpl, map[string]*bintree{}},
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
