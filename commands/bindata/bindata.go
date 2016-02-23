// Code generated by go-bindata.
// sources:
// ../templates/client_helper_resource.tmpl
// ../templates/client_resource.tmpl
// ../templates/generic_main.tmpl
// ../templates/python_client.tmpl
// ../templates/python_client_utils.tmpl
// ../templates/python_server_resource.tmpl
// ../templates/server_main.tmpl
// ../templates/server_python_main.tmpl
// ../templates/server_resources_api.tmpl
// ../templates/server_resources_interface.tmpl
// ../templates/struct.tmpl
// DO NOT EDIT!

package bindata

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

var _TemplatesClient_helper_resourceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x55\x51\x6f\xd3\x30\x10\x7e\xae\x7f\xc5\x11\x09\x64\xb7\x21\x05\xed\x69\x93\x2a\x24\x36\x4d\x08\x89\x01\xed\x10\x0f\xd3\xb4\xb9\xcd\x75\x0d\x6b\xec\xd4\xb9\x14\xaa\xaa\xff\x1d\x9f\x93\x66\x6d\x19\xd3\x5e\x90\x78\xc9\xec\xf3\x7d\xbe\xef\xfb\x7c\xb7\xae\xd7\x29\x4e\x33\x83\x10\x4d\xe6\x19\x1a\xba\x99\xe1\xbc\x40\x77\xe3\xb0\xb4\x95\x9b\x60\x19\x6d\x36\xa2\xd0\x93\x7b\x7d\x87\x50\xa7\x08\x91\xe5\x85\x75\x04\x52\x74\xa2\xf1\x8a\x7c\x8e\x5f\xa0\x99\xd8\x34\x33\x77\xfd\x1f\xa5\x35\x1c\xc8\x2c\x7f\x0d\x52\x7f\x46\x54\xf0\x7a\x9a\x13\xff\xa1\x2c\xc7\x48\x28\x21\xa6\x95\x99\x40\xc0\xe1\x7b\x9b\xae\x64\xaa\x49\x43\x66\x08\xdd\x54\x4f\x70\xbd\x51\x20\x33\x9b\x0c\x51\xa7\xe8\x62\x40\xe7\xac\x53\xb0\x16\x9d\x71\xd8\xc0\xc9\x00\xb8\x56\xf2\x49\xbb\x72\xa6\xe7\x01\xae\x44\x27\x9b\x86\xd3\x17\x03\x30\xd9\x9c\xd3\x3b\x0e\xa9\x72\x86\xb7\x01\x28\x3a\x1b\xb1\x8d\x05\xfa\xc9\x05\xfe\xac\xab\xc8\xb1\x8a\x39\x4f\x6c\x1a\x76\xa9\x1d\xe2\xe2\x7b\x46\xb3\x40\x30\x47\x9a\xd9\x34\x86\xca\xcd\x47\xe4\xa0\x24\xe7\x05\xc7\x70\xc8\x3b\x6e\x8c\x02\x16\x9e\x9c\x86\x75\x0c\xb3\x50\xa1\x84\x5c\x17\x57\x35\xf2\x7a\x0f\xb3\x28\xbf\x68\xa7\xf3\xe6\x56\xaf\xbd\x1b\xe0\x43\x2c\x0b\x6b\x4a\xdc\x33\xc0\x93\x69\x3d\x38\x30\xf0\xb9\x0e\x88\x4e\xbf\x0f\x13\x87\x9a\x10\x68\x86\xe0\x70\x51\x61\x49\xec\xcc\xa2\xbd\x3b\x30\x08\xee\x84\xc3\x43\x03\x7a\x5b\xd2\x31\x30\xa5\xe7\x97\x9e\x5a\x07\xf7\x31\x2c\xb9\x86\xd3\xc6\xb7\xd6\xd6\x1d\x86\x74\x98\x4c\xf2\x21\x44\x92\x11\x92\xf4\xa9\xbe\x77\x92\x51\xe1\x9d\xa1\xa9\x8c\x5e\x2e\xa3\x78\xa9\x54\x7d\x57\x53\xa0\xb6\x3c\x39\xb3\xd2\x83\x55\xfb\x80\xe3\x2a\x9b\xa7\x5f\x2b\x74\xab\x51\xf0\xb5\x6e\xb2\xc7\xdf\x40\x35\xde\xaf\x83\x8e\x39\x9a\xda\x50\x18\x0c\xe0\xcd\xae\x96\x28\xaa\x2b\x8f\x75\x89\xe1\x6a\x96\x11\xbd\x8b\x1e\xd3\x15\xca\x31\xf6\x21\xb9\x37\x80\x7b\x6f\x5d\x34\x88\xfc\xf7\x41\x97\x5c\x2a\x8e\xbe\x8a\xf6\x64\xb5\xb0\xab\x13\x26\xd4\x6e\xd5\xeb\xb7\xd7\x2c\xb2\xdf\x3f\xe3\x17\x74\x58\xf8\x89\xe5\x9e\x1b\x9e\x9f\x1e\x1d\x1d\x1f\x73\x61\x14\xb4\x2a\x10\x42\x02\x0f\x5d\x72\xe9\x3f\x0c\x69\x26\xe6\xe3\xe8\xf3\x05\xd8\xa5\x7f\x95\x2c\x45\xef\x49\x1b\xac\xad\x93\x04\x5d\xc6\x2a\xd8\xc9\x97\xbe\x31\xaf\xae\x79\x6c\x76\x1b\xb2\x21\x5b\x1f\xc8\xb6\x96\xec\x92\x4a\xce\xad\xcb\x35\xc9\xdb\xe8\xd6\xcb\x0b\x47\x81\xe2\xd1\xb1\xdf\xfa\xa0\x7a\x98\xb8\x96\xd8\x25\xfe\xa2\x3f\x88\x71\xf0\x2f\xc4\xf8\xe8\xdf\x12\xfb\x66\xf2\xc7\x3c\xab\xcc\x13\xae\xed\x61\xe4\xb8\x21\xa1\x6a\x76\x4c\x8e\xca\x76\xd0\x42\x79\x3f\x4a\x25\x32\x9f\xde\x2e\x9b\x9e\x0f\xc4\x4d\x67\xfa\xff\x4f\x4f\x0d\x59\x3b\x5f\x5d\x82\x41\x78\x77\x49\xa5\x12\x3b\x33\x78\xa0\x66\xdf\xe8\xca\x3c\x61\xf5\x1e\xe6\x3f\x52\x73\x40\xb3\x99\xf2\xed\x28\xef\xf4\xc0\xfe\xe3\x6f\xf3\xf8\x8a\xf5\x1a\x4d\xea\x7f\xe6\x7e\x07\x00\x00\xff\xff\x51\x79\xd6\x10\x0c\x07\x00\x00")

func TemplatesClient_helper_resourceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesClient_helper_resourceTmpl,
		"../templates/client_helper_resource.tmpl",
	)
}

func TemplatesClient_helper_resourceTmpl() (*asset, error) {
	bytes, err := TemplatesClient_helper_resourceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/client_helper_resource.tmpl", size: 1804, mode: os.FileMode(420), modTime: time.Unix(1455619563, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesClient_resourceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\x4d\x4f\xdb\x40\x10\x3d\x7b\x7f\xc5\xd4\x4a\x91\x4d\x2d\xe7\x5e\x29\x17\x20\xfd\x90\x28\xa2\x01\xda\x63\x65\xec\x31\x31\x24\x5e\x67\x77\x1d\x84\x5c\xff\xf7\xce\xec\xda\x8e\x09\x2d\x52\x91\xca\xa9\x9c\x96\xf9\x7a\x6f\xde\x3e\x6f\x9a\x26\xc3\xbc\x28\x11\xfc\x74\x55\x60\x69\x7e\x28\xd4\xb2\x56\x29\xfa\x6d\x2b\xaa\x24\xbd\x4b\x6e\x10\x5c\x4a\x88\x62\x5d\x49\x65\x20\x10\x9e\x5f\xa2\x99\x2e\x8d\xa9\x7c\xe1\x79\x3e\x96\xa9\xcc\x8a\xf2\x66\x7a\xab\x65\x69\x23\xf9\xda\xf8\x22\x14\x22\x95\xa5\xb6\x0d\x4a\x4a\x73\xb5\x38\x85\x19\xf8\x4d\x13\x1f\x25\x1a\xaf\x16\x9f\xdb\xd6\x16\x35\xcd\x24\xa9\x8a\xb3\x64\x8d\xf0\x7e\x06\x31\x1f\x08\xdc\x3c\x54\x08\x54\xeb\xfe\x05\x6d\x54\x9d\x1a\x68\x84\xe7\xd8\x00\xa3\xc7\xc7\x8e\x59\x2b\x44\x5e\x97\x29\x9c\xe1\xfd\xd0\x11\x84\x70\xb8\x6b\xe7\x3e\x9e\x5e\xe2\x7d\x30\x44\x43\x0a\xc6\xdd\xb8\xd9\x78\x60\xd3\x12\x63\x34\xb5\x2a\x21\xe5\xe1\x4d\x03\x2a\x29\x49\x89\xc9\x5d\x04\x93\xad\xa5\xf9\x05\xcd\x52\x66\x1a\x88\xea\x28\x9d\x73\x3e\xe7\x82\xc9\x36\xfe\x40\x9c\x8e\xe5\x7a\x4d\x13\x6d\xdd\x74\x4a\x0b\x51\x9a\xf9\x34\x58\x66\x14\x62\xda\x41\xca\x4c\x7b\x0d\x88\x96\xad\xea\x00\xba\x65\x6c\xe4\x3c\x51\xc9\x5a\x53\x81\xfb\x77\x81\xba\x3a\x92\xd9\x83\x1d\x57\xe4\xb4\x1b\x8c\xa2\xe0\xfb\x84\x19\xf5\x48\x87\x76\x3d\x4e\xd2\x95\x60\x84\x4a\x49\x15\xb2\x2c\x1b\x6d\xc7\x32\xe5\xeb\xba\x58\x65\x5f\x6b\x54\x0f\x17\x46\xd1\x7d\x06\x1b\x3e\x3b\x54\xd2\xca\x82\xe0\x86\x41\xbe\xa1\xba\x06\xff\xe3\xfc\x92\x6d\xe2\x79\x7f\x84\x87\x6d\xa2\xa0\x86\xa7\x74\xdd\xf6\xd4\x4a\x9a\xa4\x0a\x13\x83\xa0\x90\xe0\xc8\x2d\xf2\xfa\x16\x53\x43\x29\x0a\x44\x40\x44\x99\x9a\x65\x4f\xd7\xbb\x70\x45\x81\xc5\x8e\xa0\xb7\xd5\x63\x02\xd6\xc0\xe7\x89\x59\x76\x24\xde\xf5\x80\x03\x91\xa1\xc2\x66\x3b\x09\x22\x28\x8b\x15\xed\xe9\xf1\x9a\x04\xfb\x66\xc6\x01\xd6\xe8\xb9\x0d\x3b\x9f\xd4\xb6\xdb\xf2\x75\xf5\xb8\xd2\xb8\x4b\xef\xe5\xec\xf6\x9e\x67\x15\xc8\xa5\x02\xb2\x95\x75\x95\xb3\xd1\x12\x93\x0c\x95\x26\x64\xe8\xff\x48\x8b\xf8\x93\x0d\xc7\x17\x68\x02\xaa\xa7\x4f\x2c\xbe\xa8\xe8\x9a\x4c\x1e\xf8\x6f\xb7\x24\xc6\x36\x0c\x6d\x43\xa7\x6b\x26\xc1\x2c\x07\x59\xad\x9e\xba\x1a\x04\xed\x8d\x1f\x9f\xc8\x80\x4a\x5e\x79\x6d\xcf\xa3\x67\x07\x15\x30\xa5\x98\xa7\xd2\x87\x27\x35\x06\xa1\x78\xce\x4d\x3c\x65\x87\xeb\xd6\xe1\x27\x87\x8d\x71\x82\xf4\x06\xa1\x0a\x86\x89\x61\xec\x42\xc1\x41\x1d\x8a\x1d\xb3\xd1\x0c\x37\x80\x28\x8a\x1d\x39\xd1\x15\xc2\x9e\xd5\x4f\xe6\xa7\xf3\xcb\xb9\x73\xfb\x4b\x2d\xdb\xcd\x18\x5c\xfb\xcf\x4d\xfb\xf4\x12\x5e\xdb\x70\xee\x09\xdd\xb7\x9a\x18\xdf\xc6\xcb\xdf\x0e\x47\x73\xe4\xe9\x4c\x92\xd6\xdf\x0b\xb3\xe4\xd2\xc0\xb7\x9d\x7c\x7d\xf4\x3b\xf3\xb7\xa2\xff\x46\xf3\xe8\xf1\x2b\xb3\xe9\x98\xb6\xed\xc1\xe3\x1e\x97\xf8\x09\x97\xf2\x54\xde\xa3\xe2\xc6\x5e\x81\xa8\x57\x3a\xea\xaf\xef\xff\x77\xd7\x91\xeb\x0f\xf6\xd7\x94\x8e\x0c\x3b\x3a\xfe\x0a\x00\x00\xff\xff\x5c\x16\xe4\x69\xa8\x08\x00\x00")

func TemplatesClient_resourceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesClient_resourceTmpl,
		"../templates/client_resource.tmpl",
	)
}

func TemplatesClient_resourceTmpl() (*asset, error) {
	bytes, err := TemplatesClient_resourceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/client_resource.tmpl", size: 2216, mode: os.FileMode(420), modTime: time.Unix(1456206796, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesGeneric_mainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\x4f\xcd\x4b\x2d\xca\x4c\x8e\xcf\x4d\xcc\xcc\x8b\x2f\x49\xcd\x2d\xc8\x49\x2c\x49\x55\xaa\xad\xe5\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x00\x49\x70\x71\xa5\x95\xe6\x25\x83\x99\x1a\x9a\xd5\x5c\x5c\xb5\x5c\xd5\xd5\xa9\x79\x29\x40\x55\x80\x00\x00\x00\xff\xff\xdc\x57\x73\x81\x49\x00\x00\x00")

func TemplatesGeneric_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesGeneric_mainTmpl,
		"../templates/generic_main.tmpl",
	)
}

func TemplatesGeneric_mainTmpl() (*asset, error) {
	bytes, err := TemplatesGeneric_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/generic_main.tmpl", size: 73, mode: os.FileMode(420), modTime: time.Unix(1455614346, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesPython_clientTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x52\x4d\x6f\xdb\x30\x0c\xbd\xeb\x57\x10\x46\x0e\x49\x5b\xb8\xf7\x02\x39\xb4\x45\x07\x14\xd8\x86\x20\xfb\xb8\x0a\xaa\x4d\xc7\xc2\x6c\xc9\xa3\xa8\x14\x81\xa7\xff\x3e\xca\x4e\x0d\x1f\xca\x83\x4c\x12\xe4\xe3\x7b\xa4\xc7\x11\x6a\x6c\xac\x43\x28\x86\x0b\xb7\xde\xe9\xaa\xb3\xe8\x58\x33\xf6\x43\x67\x18\x0b\x48\x49\xd9\x7e\xf0\xc4\x40\xf8\x37\x62\xe0\xa0\x1a\xf2\x3d\x5c\x0b\x23\xdb\x2e\xc0\xb5\xe2\x2d\xda\xae\xd6\x52\x45\x17\x1d\x98\xac\x3b\x29\xf5\xf4\xf8\xe3\x45\xff\x3a\xbe\xc2\x1e\x8a\x71\x2c\x9f\x4c\x40\x89\x52\x2a\x94\x52\x55\x67\x42\x80\xe7\x09\xe9\x41\x81\x98\xb0\x01\xad\xad\xb3\xac\xf5\x36\x60\xd7\xec\xe6\x7c\xb6\x1c\x96\x91\x3a\x41\xfa\x00\x55\x22\x80\x8c\x3b\x21\x6c\xfe\xdc\xc1\xe6\x0c\x0f\x7b\x28\x0f\x93\x92\x6f\x28\x6f\x1d\x32\xff\x0f\xe0\x71\xdc\x9c\xcb\x39\xff\xdd\xf4\x98\xd2\x76\xca\x1c\x0c\x99\x3e\xa4\xb4\x1a\x55\x14\xc5\x0a\xb9\xc9\xd0\x4d\xc6\x96\xea\x2f\xd1\x55\xcf\xbe\xef\x85\xf2\x02\x9e\x2d\x43\x35\x29\x8d\x23\xba\x7a\x95\x7e\x65\xb0\x01\xfa\x69\x28\x34\x9e\x66\x12\xbf\x91\xde\x52\x9a\xfd\x17\x57\x0f\xde\x3a\x5e\x35\xc9\xf8\xc5\x8f\x64\x45\xf0\xa2\xfd\x76\x6e\x3a\x62\xf0\x91\x2a\x3c\x18\x6e\x57\x8d\x73\x71\x7e\x6f\x3f\x39\xc6\x76\x0e\x86\x49\xef\x6e\x69\x22\xe4\x48\x6e\x39\x6f\xb9\x30\x84\x7f\xf0\xd3\x7f\xf5\xef\x48\xb2\x2a\x01\x9d\xb7\x75\x7c\xa4\x93\x6c\xeb\x0e\x5a\x34\x35\x52\xd8\x5f\xbf\xbb\x7c\x0c\x11\x0f\x79\x09\x8b\x73\x7f\x03\x43\x64\xe0\x16\xa7\x54\x8b\x84\xc0\x1e\xcc\xd9\xdb\x1a\x44\x35\xf9\x3a\x56\xc2\x0d\xa2\x73\x58\x61\x08\x86\x2e\xe0\xf0\xbd\xcb\x3f\xe5\xcd\xbd\x68\xfb\x1f\x00\x00\xff\xff\x76\x46\xb2\xf7\xa4\x02\x00\x00")

func TemplatesPython_clientTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPython_clientTmpl,
		"../templates/python_client.tmpl",
	)
}

func TemplatesPython_clientTmpl() (*asset, error) {
	bytes, err := TemplatesPython_clientTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/python_client.tmpl", size: 676, mode: os.FileMode(420), modTime: time.Unix(1455866780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesPython_client_utilsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xcd\x6e\xdb\x46\x10\xbe\xeb\x29\x06\x0c\x14\x93\xb5\xa8\xd0\x52\x90\x1f\x21\x6a\x0e\x45\x51\xf4\x92\x93\x73\x0a\x02\x79\x45\x0e\x2d\xd6\xe4\x2e\xbd\xbb\xb4\xc3\x04\x79\xf7\xce\xec\x92\x14\x29\x2b\x69\x0e\x35\x0c\x6a\xb5\xf3\xff\xcd\x7c\x43\x7d\xfb\x16\x43\x86\x79\x21\x11\x82\xba\xb5\x07\x25\x77\x69\x59\xa0\xb4\xbb\xc6\x16\xa5\x09\x20\xfe\xfe\x7d\x56\x54\xb5\xd2\x16\x32\x61\xd1\x16\x15\xf6\xdf\xdd\x79\x36\x23\x73\xd8\x37\x45\x99\xed\xee\x1b\xd4\xed\xce\x58\x5d\xc8\xdb\xd0\x7f\xa9\x85\x16\x95\xd9\x7e\x50\x12\xa3\xcd\x0c\xe8\x2f\x08\x02\xf7\xe9\x4c\xc0\x69\x81\xd3\x42\x8b\xda\x49\x0a\x59\x37\x16\xe0\x66\xec\xe1\x06\x36\x90\x15\xa9\x2d\x94\x14\xba\x75\x6a\xaa\xb1\xac\xe7\xc3\xc1\x63\x61\x0f\x90\x2b\x5d\x09\x0b\x37\xef\xef\xb0\xdd\x3e\x88\xf2\x39\x7d\xae\xf8\xb0\xba\x99\xc4\x2e\x72\x18\x3b\x87\xc2\x00\x67\xe8\x13\xe4\x3f\x8d\xb6\xd1\x92\xf4\x67\xee\xea\xde\xc0\x16\x82\xf7\xde\x98\x82\x00\xf9\x5d\x00\x96\x58\x51\xb2\x13\x57\x47\x17\x64\x73\xb9\x65\x45\xb8\x84\x60\x1b\xd0\x93\x32\x0d\xd9\x26\xe2\x9b\xe7\x9d\xeb\x2e\xd2\xbd\xf9\xb4\x89\xaf\x3e\x77\x70\xde\xa2\x44\x4d\x68\xef\x74\x9e\xae\xd7\xeb\xb7\x61\xb6\x80\x52\xa5\xa2\xdc\xd9\xaf\xdb\x6b\xdd\x9c\x62\xd9\xeb\x43\xa7\xef\x5a\xd3\xa1\x31\x82\xd4\x1b\x65\x54\x0b\xb7\x12\x6c\x5b\xa3\xbb\xe9\x5d\x93\xa0\x31\xe8\xbf\x7a\x17\x5f\x09\x15\x46\xcb\x52\xcc\x85\x47\xdd\x1e\x50\x3f\x16\xa4\x56\x09\x7d\x07\xc2\x40\x63\xd3\xd9\xb8\x21\x3e\x4a\x9f\x49\xd7\x1e\x17\xd0\x27\xb4\x04\xfc\x42\xdd\xbc\x59\x25\xc9\x9b\x38\x79\x19\x27\xab\xeb\x55\xb2\x49\xf8\xff\x32\x79\x4d\xcf\x69\xaf\xac\x6e\x8f\xa0\x52\x2a\x7d\xb2\xc7\xcb\x71\x51\x9c\xf5\x72\x38\xe4\x5a\x55\x7c\x30\x56\x54\x75\x98\x45\x83\x09\x96\x06\x7f\xc5\x01\xd5\x76\xd6\x07\x7e\x49\xb1\xb6\x70\x4d\x08\xfe\xa9\xb5\xd2\x47\x5f\xb5\x30\x66\xd6\x0f\x99\x54\x96\x66\xab\x90\x64\x2b\x53\xe4\x2e\x4e\x42\x44\xa3\x81\x13\x0c\xe9\xe0\x2f\xbc\xf8\xa0\x3c\xc1\x5c\x58\xa0\x91\x73\x08\xaa\xfd\x3f\x98\x12\x82\x7f\x91\x74\xae\x97\x17\x30\x77\x5d\xa4\xb4\xa2\x5f\x0d\xca\x87\x51\xe0\xb3\x75\x87\xbf\x65\x4b\xfe\xb4\x4d\x5d\x62\x18\x7d\xda\xac\x3f\x47\x93\x71\x0d\x2f\xe6\xc9\xcb\x2c\x9e\x27\x2b\xff\xb8\xe6\xc7\x66\x78\xcc\x0d\x65\x36\x81\x37\xcc\x96\x2d\x0a\x4d\xb9\x2c\x2b\x25\xed\x81\x0f\x99\x68\xf9\xe3\xa0\x1a\x7f\x5f\xc8\xc6\x22\x9f\x0c\xa6\x4a\x66\x8b\x89\x03\xd8\x0d\xa4\xe0\xcc\x78\x30\xc7\xac\x88\x18\x00\xc7\x9d\x1d\x5d\xa4\x4d\xc9\x8a\x2a\xcf\x0d\xda\x90\xab\x1a\x69\x4e\xa9\x33\x61\x06\x43\xbc\xf9\x11\x39\x36\x03\x11\x1c\x4b\x4c\x6b\x2c\xf1\xbf\x4f\x66\x31\x62\x46\x07\x52\x32\x81\x8c\xa4\x5d\x0f\x73\xf8\x78\xfd\x07\xf8\xe4\x96\x4e\xe5\xef\xdc\x8b\x32\x85\xc6\xf5\xef\x20\x1e\x10\x84\x6c\x07\xf7\x94\x67\xae\x16\xf0\x88\x27\x0c\x75\xa1\x4f\x88\x39\x84\x9f\xd4\x79\x8e\x39\xcf\x0e\x42\x66\x25\x02\xb7\x06\xf6\x48\xfc\x44\xb8\x7a\xfb\x3a\x81\x4a\x19\xda\xac\xad\x2b\x90\xdc\x22\x6f\x48\xa9\xa6\xd9\x30\x97\x69\x1d\x8f\xed\x96\x63\xa2\x72\x45\xae\xe9\xf0\xce\x09\xa7\x7c\x7b\x06\x1f\x8d\x33\x5a\x91\x83\x54\x34\xa6\x8b\xcc\x10\xc8\x8b\x1e\x01\x28\x51\xd4\xe4\xa9\x9d\xd8\x5a\x9a\x59\x37\xaf\xd5\x9d\x9b\x56\x17\x49\x63\x5d\x0a\x9a\x77\x8e\xb8\x65\xbf\xd1\x71\x84\x7f\x46\xfc\xb3\xce\x46\xc3\xdf\x0d\xbe\xcf\xb9\xc3\x8b\x12\x2a\x8b\xdb\x83\x35\xe2\x81\x96\xdb\xc2\x57\x3b\xbe\x72\x5d\x12\xa5\x3d\x9d\x0c\xbe\xee\x41\x1c\x63\xe5\x12\x70\xfd\x71\x39\x58\xca\xbd\xda\x15\x26\x33\x76\x9a\x6c\xd7\xda\xd8\xe9\x77\xfe\x7f\x52\xdb\x44\x7d\x12\x77\xaa\x7b\x1c\x58\x4f\xa1\x33\x54\xfb\xbf\x38\xb4\x57\xaa\xec\x5e\x17\x8e\x00\xc3\xbb\x2b\x03\xde\xb4\x4f\xe9\x3b\xd6\xa5\x97\xad\xdf\x0d\x66\x7c\xbb\x85\x04\xe2\xdf\xe1\xd2\xbd\x42\xa6\x82\xab\x37\xc9\x20\x5b\x9f\xc8\xe2\xf5\x2b\x2f\x8c\x93\xab\xde\xb0\xaf\x68\x50\xfa\xcf\x6d\xe2\x6b\xe1\x1d\x46\xda\x62\x6f\x42\xaf\x16\xc1\x8b\x17\xc0\x01\x9c\xd8\xaf\xb6\x13\x85\xb9\x93\xb3\xde\xab\x64\x58\xde\x5d\xe0\x77\x90\x3c\xe9\xce\xc5\x3c\x1d\x16\x2c\x2f\xfe\x30\x88\x83\x05\xf8\xed\xe9\x03\x44\x3f\x6e\xed\x53\xe3\xcb\x27\xc6\xfc\x5b\x10\x65\xe6\x7e\xf2\xfd\x1b\x00\x00\xff\xff\x4b\x7e\x1f\x35\x18\x0a\x00\x00")

func TemplatesPython_client_utilsTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPython_client_utilsTmpl,
		"../templates/python_client_utils.tmpl",
	)
}

func TemplatesPython_client_utilsTmpl() (*asset, error) {
	bytes, err := TemplatesPython_client_utilsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/python_client_utils.tmpl", size: 2584, mode: os.FileMode(420), modTime: time.Unix(1456220676, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesPython_server_resourceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x91\x51\x4b\xc3\x30\x14\x85\xdf\xf3\x2b\x2e\xa3\xd0\x0d\xba\xfd\x80\xc1\x40\x14\x05\x41\xc5\x07\xf1\x45\x24\xc4\xf5\xc6\xd5\x35\x49\xbd\x49\x27\xa3\xe6\xbf\x7b\x9b\xb0\x4d\x67\x5f\x1a\xee\x3d\xf9\xce\xe9\xe9\x30\xcc\xa1\x46\xdd\x58\x84\x09\xa1\x77\x3d\xad\x51\x76\xfb\xb0\x71\x56\x06\x34\x5d\xab\x02\x4e\x60\x1e\xa3\x18\x95\x85\xea\x9a\x07\x65\x10\x96\x2b\x58\xa4\xc3\xb8\xd1\xe4\x0c\xe8\x56\xf9\x2d\x34\xa6\x73\x14\xe0\xb2\xed\xb1\xa3\xc6\x86\x0a\x3e\xbc\xb3\x8d\xde\x57\x40\xf8\xd9\xa3\x0f\x82\x41\xf9\xea\x37\x3c\xb9\x3b\xf7\x85\x04\x31\x4a\x06\xc3\xea\x74\x6f\x5a\xfe\x53\x65\x51\x59\x81\x94\x96\x17\x52\xce\x98\x04\xa4\xec\x3b\x42\xb1\xad\xa0\xd8\xa5\x54\xf7\xc8\xd9\x6b\xcf\x4c\x21\x2e\x86\xe1\x98\xf8\x8c\xb3\x20\xd7\x07\x1c\x6d\x8a\xdd\xe2\xda\xd6\x9d\x63\xd7\x18\x19\x6f\x32\x60\xf5\x92\x77\xcf\x48\x6f\x3c\x7f\x9d\x09\xae\x09\xd2\x28\x5b\x8c\xd4\x18\xa7\x69\xf2\xa8\x48\x19\x1f\xe3\x6c\x29\x80\x9f\xb2\x2c\xd3\x7b\x18\x0e\xf1\xf4\x98\x4f\x8f\x01\x59\x7d\xd3\xdb\xf5\x95\x33\x06\x6d\xf0\xa9\xc0\xac\x65\xc1\xf1\x8c\xb6\x3e\x6e\x6e\x03\x34\x1e\x36\xca\xd6\x2d\x77\xa5\x1d\xc1\xaf\x60\x70\xf6\x01\x7f\xfc\x09\x43\x4f\xf6\xf0\x0b\xa6\xa9\xb0\x03\x59\x9c\x4c\x7e\x02\x00\x00\xff\xff\x47\xce\x2f\x6b\x04\x02\x00\x00")

func TemplatesPython_server_resourceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesPython_server_resourceTmpl,
		"../templates/python_server_resource.tmpl",
	)
}

func TemplatesPython_server_resourceTmpl() (*asset, error) {
	bytes, err := TemplatesPython_server_resourceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/python_server_resource.tmpl", size: 516, mode: os.FileMode(420), modTime: time.Unix(1456221131, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_mainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x94\x51\x6f\xda\x30\x10\xc7\x9f\xe3\x4f\x71\xb3\xfa\x60\x43\x14\x2a\xf1\x52\x2a\xf1\x50\xad\xaa\xd4\x69\x63\x08\xba\xa7\x69\x6a\x0d\x1c\x60\x35\x71\x22\xfb\xd2\xb5\x8a\xf8\xee\x3b\x27\x2d\x85\x75\xeb\xdb\xa4\xbd\x80\x73\xfe\xff\xcf\xbf\xbb\xe4\xdc\x34\x2b\x5c\x5b\x87\x20\x03\xfa\x07\xf4\xb7\x85\xb1\xee\x96\xb0\xa8\x72\x43\x28\x77\x3b\x51\x99\xe5\xbd\xd9\x20\x34\x4d\x36\xed\x96\x13\x53\x20\x6f\x08\x5b\x54\xa5\x27\x50\x22\x91\x79\xb9\x91\xfc\xe7\x90\x06\x5b\xa2\x2a\xae\xc9\x16\x28\x05\x2f\x36\x96\xb6\xf5\x22\x5b\x96\xc5\x60\x53\x7a\x9b\xe7\x66\x50\xd4\x8f\x52\x68\x21\x06\x83\x4b\x3e\x05\x3c\x56\x1e\x03\x3a\x82\xd9\xd5\xc7\xe1\x70\x34\x82\x15\x87\x05\x3d\x55\x08\xad\x20\xe6\xca\x6e\xf8\x27\x5a\xbe\x18\x1f\xb6\x26\xff\x34\xff\x3a\x81\x92\x91\xbd\x5d\x21\x14\xaf\x41\xb1\xae\xdd\x12\x14\x41\x2f\x7a\x35\x1c\xe8\x95\x06\xf5\xfd\xc7\xe2\x89\x30\x05\x36\x96\x5e\x43\x23\x12\x8f\x54\x7b\x07\xdd\x86\xda\x9f\xa5\x7a\xa4\xb3\xab\xd2\x17\x86\xd4\x9d\xbc\x83\x7e\x87\xd1\x22\x0e\x47\xfc\xc8\x41\xad\x53\x70\x36\x17\xbb\x03\xb0\x1b\x7c\xa4\x37\x60\x31\xf8\x17\xb0\xb8\xf5\x6f\xc1\xbe\xb9\xe2\x4f\x3d\xab\xdd\x3b\x5d\x3b\xf2\xa8\xc5\x33\x84\xee\xe8\x22\x1c\x85\x16\x15\xce\xc7\xdd\xf1\x53\x96\x63\xe4\xe9\x1f\xd2\xf4\x39\x90\x42\x20\x6f\xdd\x46\x2d\xb4\x16\x89\x5d\xb7\xb6\x0f\xe3\xc8\x17\x13\xbd\x94\xc9\x51\x91\x30\x6f\xd2\x23\x18\xb7\xef\x5d\x51\xd0\xfb\x2e\xbc\xad\xe6\xb8\xd1\xb5\x7b\xa7\xd5\x47\x9e\xff\xa8\x9a\xdf\x30\xe7\x5d\x66\xfd\x7c\xc4\xc1\x37\x70\xfc\xf2\x5f\x74\xfb\x14\x71\x6a\x55\xf7\xcd\xc4\x12\x78\xbe\xb2\x09\xfe\x9c\x95\x35\xa1\x67\x99\x48\x9a\x06\xbc\x71\x3c\xc6\x27\xf7\x29\x9c\x3c\x44\x51\x36\xc3\x50\xd6\x7e\x89\xe1\x12\xd7\xc0\x03\xcd\xa2\xac\x9b\xed\x6b\xc7\xbe\xb5\x59\x62\x9b\x21\x28\x9f\xc2\x7e\xef\x62\x7a\xdd\xec\x74\x9b\x11\xdd\x2a\xfa\x44\xc2\xe3\x9f\x4d\x19\x89\x72\xa7\x64\x20\xe3\x29\xd2\x77\x37\x8a\x64\x6d\xbc\x13\xb2\xcf\x36\x10\xba\x0b\xb7\x9a\xc7\xb8\x92\xe7\x67\xa7\x67\xa7\x32\x05\xdf\x96\xf1\x9a\xee\x57\x00\x00\x00\xff\xff\x29\x06\x41\xfe\x93\x04\x00\x00")

func TemplatesServer_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_mainTmpl,
		"../templates/server_main.tmpl",
	)
}

func TemplatesServer_mainTmpl() (*asset, error) {
	bytes, err := TemplatesServer_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_main.tmpl", size: 1171, mode: os.FileMode(420), modTime: time.Unix(1455760541, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_python_mainTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\xd0\xcd\x6a\xeb\x30\x10\x05\xe0\xbd\x9f\xe2\x60\xb2\x70\xe0\xc6\x0f\x10\xf0\xee\xd2\x55\xe9\xa2\x64\x3f\x28\xf5\xc8\x15\xb1\x7e\x18\x49\x29\x45\xf5\xbb\x57\x8e\x1b\xba\x68\xa1\xda\x09\xe9\x9c\x8f\x99\x52\x0e\x18\x59\x1b\xc7\x68\x23\xcb\x95\x85\xc2\x7b\x7a\xf5\x8e\xac\x32\x8e\x12\xdb\x30\xab\xc4\x2d\x0e\xcb\xd2\x68\xf1\x16\x7a\x56\xf1\x02\x63\x83\x97\x84\x87\xf5\xd2\x34\xa5\x40\x94\x9b\x18\xbb\xcb\x3f\xec\xae\x38\x0e\xe8\x9f\x39\xfa\x2c\x2f\x1c\xff\xb3\xfe\x4e\x97\xd2\x3f\x29\xcb\xf8\xc0\xc9\x3f\xfa\x37\x96\x65\xb9\x77\xfd\xf2\x44\x2a\x98\x5a\xce\x6e\x44\x2d\x68\x54\x08\x18\x36\xb3\x23\x72\xf5\x33\xd1\x7e\xd5\xff\xc6\x6b\xb4\x17\x9e\x4c\x4c\x75\xc2\xf3\x9c\x39\x88\x71\xa9\xfb\x61\x62\x43\xf7\x9b\xba\xa2\x46\xe3\x6e\x61\x18\xd0\xd2\xb6\x19\x6a\x8f\x0d\xea\xb9\x15\x67\xd7\x8d\x7c\xce\xd3\x70\x92\xcc\x5f\xd9\x1b\xfb\x19\x00\x00\xff\xff\xba\x73\xf0\x48\x5f\x01\x00\x00")

func TemplatesServer_python_mainTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_python_mainTmpl,
		"../templates/server_python_main.tmpl",
	)
}

func TemplatesServer_python_mainTmpl() (*asset, error) {
	bytes, err := TemplatesServer_python_mainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_python_main.tmpl", size: 351, mode: os.FileMode(420), modTime: time.Unix(1456221178, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_resources_apiTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x53\x4f\x6f\xdb\x3e\x0c\x3d\xdb\x9f\x82\x3f\x21\xf8\xc1\x1e\x0a\xa5\x87\x9d\x56\xf4\xb0\x3f\x2d\xd6\x01\xcb\xb2\x15\xe8\x8e\x85\x1a\xd3\x8d\x17\x5b\x72\x68\x39\x46\x20\xf8\xbb\x8f\x94\x92\x65\xf3\x89\x22\x1f\x1f\xdf\xa3\xe4\x10\x2a\xac\x1b\x8b\xa0\x08\x07\x37\xd2\x06\x9f\x4d\xdf\x3c\x7b\xec\xfa\xd6\x78\x54\xf3\x9c\xf7\x66\xb3\x33\xaf\x08\x21\xe8\x75\x0a\x57\xa6\x43\x2e\xe4\x4d\xd7\x3b\xf2\x50\xe4\x59\x08\xd0\xd4\xa0\x57\x88\xd5\x97\xc7\x6f\x2b\x98\x67\x85\x76\xe3\xaa\xc6\xbe\x2e\x7f\x0d\xce\x2a\x06\xa0\xad\x38\x9f\x67\xca\xa2\x5f\x6e\xbd\xef\x55\x5e\xe6\x21\x2c\x78\x9e\x10\xc2\xbb\x5b\x26\x48\xcc\xfe\xd8\xc7\x79\xe9\xf8\x7e\xfd\x00\x83\xa7\x71\xe3\x21\xe4\x3c\x96\x9b\x80\x8c\x65\x49\x8b\xdd\x15\x2c\x0e\xb1\xf3\x2b\xfa\xad\xab\x06\x99\xf0\x57\xb9\x96\x7a\x2d\x80\xc5\x41\xdf\x8f\x76\xf3\xd1\x75\x1d\x5a\x3f\x30\x6c\xb9\xe4\x11\x5c\x9d\xe7\x10\x58\x5b\xca\x3c\x78\x68\x06\xd8\x1a\x5b\xb5\x48\x50\x3b\x8a\x18\xfd\x84\xf4\x32\xcf\x29\xbe\xb3\x55\xef\x1a\xeb\xf9\x9c\xd7\x4c\x59\xb0\x01\xb8\xf8\x88\x7a\xcb\x04\x4d\xa2\x52\xb6\x98\x40\x4c\xeb\x1f\x38\xf4\xce\x0e\xf8\x93\x1a\x8f\x74\x05\x04\x6f\x4e\xf9\xfd\x88\x83\xe7\x4e\x59\xe7\xd9\xc0\x5e\x0c\xec\x4f\x06\xbe\x8f\x48\xc7\xb5\x21\xe6\xe3\x56\xf1\x00\xfc\x25\x1f\xbb\x3d\xeb\x61\x18\xe1\x5e\xdf\x3b\xea\x9e\x4c\x3b\x62\xa1\x4e\x15\x55\x0a\x69\x72\xc9\x81\xdc\x15\xcf\xfb\xe0\xaa\xa3\x24\x0e\x86\xa4\x4f\x8e\xb2\xf5\x4b\x25\xcf\x18\x89\x44\x42\x2c\xd7\xc8\x17\x3c\x7d\x42\xbe\x57\xa4\x82\xb4\x80\x4a\x9d\xce\xc5\xff\x27\x82\xf2\x26\x36\xfc\x77\x0b\xb6\x69\xc5\x4b\x36\xe9\x68\xf5\x33\x1a\x69\x7b\x7b\x7d\xcd\x5a\x32\x42\x3f\x92\xcd\xb3\x28\xe7\xfc\x34\xfe\xbc\x23\xd9\x51\x54\x73\x51\x77\x4a\x44\x79\x29\x96\xda\x59\xd4\x9d\x4d\xa2\xa6\x52\xa7\x50\xf4\x24\x58\x79\xf3\xef\x08\x5e\x17\x5f\x5a\x7a\x06\xf0\x82\xad\x9b\xa0\x95\x1f\xc0\x3b\x30\x55\x05\xdb\x28\x33\xc2\x26\x9d\x34\xeb\x47\xf4\x85\xda\xe1\x51\x5d\xa9\x83\xac\x95\xb7\x19\x5f\xd9\x99\xf3\x12\xfe\x0e\x00\x00\xff\xff\xe9\xd7\x49\x0a\x4f\x03\x00\x00")

func TemplatesServer_resources_apiTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_resources_apiTmpl,
		"../templates/server_resources_api.tmpl",
	)
}

func TemplatesServer_resources_apiTmpl() (*asset, error) {
	bytes, err := TemplatesServer_resources_apiTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_resources_api.tmpl", size: 847, mode: os.FileMode(420), modTime: time.Unix(1455866780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesServer_resources_interfaceTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x3f\x6f\xdc\x30\x0c\xc5\x67\xeb\x53\x10\xc6\x0d\x77\x41\x22\xef\x05\x3a\xf5\x0f\x9a\xa1\x45\x11\x14\xed\x18\xe8\x6c\xda\x16\x62\x53\x2e\x45\x5f\x12\x08\xfe\xee\xa5\xec\x4b\x5b\xa4\x37\x74\x32\x45\x3d\x91\xef\xf7\x9c\x52\x83\xad\x27\x84\x92\x31\x86\x99\x6b\xbc\xf7\xed\xbd\xe0\x38\x0d\x4e\xb0\x5c\x16\x33\xb9\xfa\xc1\x75\x08\x29\xd9\xaf\x5b\xf9\xc5\x8d\xa8\x17\xa6\xaa\xbe\xf5\x3e\x42\xeb\x07\x04\xfd\xba\x59\xc2\x4d\x87\x84\xac\x2f\x1b\x38\x3e\x43\x17\x6e\xd8\x8d\x83\x0a\xdf\x07\xa0\x20\x80\x8d\x17\x90\xdf\x8f\x54\xd2\x3b\x6a\x20\x7a\xaa\x75\x84\xc0\xa3\x1f\x06\x38\x22\x84\x13\xf2\x23\x7b\x11\x24\x68\x66\xf6\xd4\xe9\x2b\x04\xc2\x27\x81\xf3\x06\x1f\xc8\x18\x3f\x4e\x81\x05\xf6\xa6\x28\x3b\x2f\xfd\x7c\xb4\x75\x18\xab\x2e\xb0\xce\x71\xd5\x38\x3f\x95\x7a\x43\x28\x55\x2f\x32\x95\xe6\x60\x8c\x3c\x4f\x2b\xca\xc6\x70\x4b\x82\xdc\xba\xbc\xfc\xa5\x4a\xa6\x48\x09\xd8\x91\x22\xef\x1e\xae\x61\x77\x82\x37\x6f\xc1\x7e\x46\xe9\x43\x13\x41\xb9\x0b\x80\xbf\x25\x6d\xd6\xb4\x59\xb4\x3b\xd9\x8f\x33\xd5\xef\xc2\x38\x22\x49\x3c\x4b\xab\x4a\xf7\xa9\x62\x59\x52\x42\x6a\xb6\x2e\x80\xb6\x6f\x25\xc7\x96\x13\x18\x90\xa1\x0d\xbc\x0a\xed\x77\xe4\xe3\xb2\x6c\xf5\x07\x6a\xa6\xa0\xde\xf4\x6c\x8a\x62\x6d\x6d\x4e\x36\xff\xfb\xcc\x65\xef\x30\x4e\x81\x22\xfe\xd0\xc4\x90\xaf\xe1\xea\xdc\xfd\x39\x63\x94\xc3\xca\xa3\x8b\xb3\x75\xfd\x69\xad\x3a\xbc\x10\xc0\x5d\x98\x05\xe3\x9e\xe1\x4a\x53\xb3\xeb\x49\x27\xf9\x0b\xca\x03\xc0\x7f\x64\x54\xb0\xfd\xb4\x82\xe5\x44\xf6\xe5\x2b\x98\x52\x47\xdb\x7f\x68\x0e\x2f\x13\xce\xfa\x2d\x88\xf2\x15\xc1\x9f\xfa\x57\x00\x00\x00\xff\xff\x14\xac\x1a\x3b\xbc\x02\x00\x00")

func TemplatesServer_resources_interfaceTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesServer_resources_interfaceTmpl,
		"../templates/server_resources_interface.tmpl",
	)
}

func TemplatesServer_resources_interfaceTmpl() (*asset, error) {
	bytes, err := TemplatesServer_resources_interfaceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/server_resources_interface.tmpl", size: 700, mode: os.FileMode(420), modTime: time.Unix(1455866780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TemplatesStructTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\x41\x8f\x83\x20\x10\x85\xef\xfc\x8a\x09\xf1\xb8\xc1\xbb\xc9\x9e\x76\xb3\xc9\x5e\x9a\x1e\x7a\xaf\x44\xc7\x86\x8a\x40\x05\x4d\x0c\xe1\xbf\x77\x10\xdb\x7a\x71\x98\x37\xef\x7b\x2f\xc6\x1e\x07\x65\x10\xb8\x0f\xf3\xd2\x85\x6b\xc0\xc9\x69\x19\x90\xa7\xc4\x9c\xec\x46\x79\x43\x88\x51\x9c\xcb\x78\x92\x13\x92\xc0\x62\x84\x59\x1a\x92\xaa\x15\x9a\x6f\x10\xbf\xe8\xbb\x59\xb9\xa0\xac\x01\xd2\xeb\x9a\x3c\xd5\x9a\x12\xfd\xd0\xf4\xb4\x09\x9b\xcb\x1c\x10\x99\x40\x27\x50\xe2\x20\x32\xa0\xef\x83\x1b\x71\xfb\x22\xa8\xd4\x0b\xee\xe0\x3f\x85\xba\xf7\x64\xc8\xbc\xbc\x15\xa5\x42\xf6\xa8\x01\xf0\x71\x1c\x8b\x7f\xff\x63\x27\x67\xbd\xda\x3b\x0c\x52\x7b\xdc\xe3\x0f\xf9\x42\xf9\xf4\x6e\xef\xde\x9a\x86\xd3\x9a\x82\x52\xe2\xed\xbb\x60\xa9\x51\xe6\xc4\x5e\xd3\x33\x00\x00\xff\xff\x9a\x27\xfe\xad\x1f\x01\x00\x00")

func TemplatesStructTmplBytes() ([]byte, error) {
	return bindataRead(
		_TemplatesStructTmpl,
		"../templates/struct.tmpl",
	)
}

func TemplatesStructTmpl() (*asset, error) {
	bytes, err := TemplatesStructTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../templates/struct.tmpl", size: 287, mode: os.FileMode(420), modTime: time.Unix(1456191299, 0)}
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
	"../templates/client_helper_resource.tmpl": TemplatesClient_helper_resourceTmpl,
	"../templates/client_resource.tmpl": TemplatesClient_resourceTmpl,
	"../templates/generic_main.tmpl": TemplatesGeneric_mainTmpl,
	"../templates/python_client.tmpl": TemplatesPython_clientTmpl,
	"../templates/python_client_utils.tmpl": TemplatesPython_client_utilsTmpl,
	"../templates/python_server_resource.tmpl": TemplatesPython_server_resourceTmpl,
	"../templates/server_main.tmpl": TemplatesServer_mainTmpl,
	"../templates/server_python_main.tmpl": TemplatesServer_python_mainTmpl,
	"../templates/server_resources_api.tmpl": TemplatesServer_resources_apiTmpl,
	"../templates/server_resources_interface.tmpl": TemplatesServer_resources_interfaceTmpl,
	"../templates/struct.tmpl": TemplatesStructTmpl,
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
	"..": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"client_helper_resource.tmpl": &bintree{TemplatesClient_helper_resourceTmpl, map[string]*bintree{}},
			"client_resource.tmpl": &bintree{TemplatesClient_resourceTmpl, map[string]*bintree{}},
			"generic_main.tmpl": &bintree{TemplatesGeneric_mainTmpl, map[string]*bintree{}},
			"python_client.tmpl": &bintree{TemplatesPython_clientTmpl, map[string]*bintree{}},
			"python_client_utils.tmpl": &bintree{TemplatesPython_client_utilsTmpl, map[string]*bintree{}},
			"python_server_resource.tmpl": &bintree{TemplatesPython_server_resourceTmpl, map[string]*bintree{}},
			"server_main.tmpl": &bintree{TemplatesServer_mainTmpl, map[string]*bintree{}},
			"server_python_main.tmpl": &bintree{TemplatesServer_python_mainTmpl, map[string]*bintree{}},
			"server_resources_api.tmpl": &bintree{TemplatesServer_resources_apiTmpl, map[string]*bintree{}},
			"server_resources_interface.tmpl": &bintree{TemplatesServer_resources_interfaceTmpl, map[string]*bintree{}},
			"struct.tmpl": &bintree{TemplatesStructTmpl, map[string]*bintree{}},
		}},
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
