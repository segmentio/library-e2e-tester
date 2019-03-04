// Code generated by go-bindata.
// sources:
// fixtures/alias/advance-alias.json
// fixtures/alias/alias.json
// fixtures/group/advance-group.json
// fixtures/group/group.json
// fixtures/identify/advance-identify.json
// fixtures/identify/identify.json
// fixtures/page/page.json
// fixtures/screen/advance-screen.json
// fixtures/screen/screen.json
// fixtures/track/advance-track.json
// fixtures/track/track.json
// DO NOT EDIT!

package tester

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

var _fixturesAliasAdvanceAliasJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8f\x31\x4e\xc6\x30\x0c\x85\xf7\x9c\xc2\xf2\xcc\xc4\x46\x37\x60\x00\x76\x38\x40\x94\x5a\xa9\xa5\xd6\xa9\x12\x07\x11\x55\xb9\x3b\x8a\xfb\xab\x1d\xed\xef\xcb\x7b\xf1\xe1\x00\x50\xdb\x4e\x38\x01\xfa\x95\x7d\xc1\xa7\xb1\xaa\x85\xf2\xd7\x3c\x96\xa1\xe6\x4c\xa2\x3f\x85\xf2\x89\xf6\x4c\xbf\x9c\x6a\x39\xf1\x98\x6e\xe6\x25\x49\xdb\x2e\x38\xc6\xe7\x97\x13\xb1\x28\xc5\xec\x95\x93\x14\x9c\x60\x14\x03\xe0\x47\x4a\x71\x25\x78\x15\xbf\x36\xe5\x60\xa4\x3b\x80\x6e\x6f\x42\x12\xa5\x3f\xbd\x75\xb6\xd8\xe3\xe0\xb9\x77\x8b\x35\x69\xdb\xbd\xb4\x4b\x02\x40\xf1\x9b\x1d\xf4\xc6\xf1\x3d\x3d\x3c\xfb\xc2\x5c\x8b\xe6\xa1\xe2\x27\xc7\x05\xbe\x29\x2c\x68\xd4\x3a\x5d\x77\xff\x01\x00\x00\xff\xff\xa9\x1a\x38\x65\x10\x01\x00\x00")

func fixturesAliasAdvanceAliasJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesAliasAdvanceAliasJson,
		"fixtures/alias/advance-alias.json",
	)
}

func fixturesAliasAdvanceAliasJson() (*asset, error) {
	bytes, err := fixturesAliasAdvanceAliasJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/alias/advance-alias.json", size: 272, mode: os.FileMode(420), modTime: time.Unix(1551658315, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesAliasAliasJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x2a\xa9\x2c\x48\x55\xb2\x52\x50\x4a\xcc\xc9\x4c\x2c\x56\xd2\x01\x09\x95\x16\xa7\x16\x79\xa6\x80\x04\x93\x4b\x8b\x8a\x52\xf3\x4a\x42\x8b\x53\x8b\x20\x52\x05\x45\xa9\x65\x99\xf9\xa5\xc5\x10\x69\x10\x0f\x21\x97\x9c\x9f\x57\x92\x5a\x51\xa2\x64\xa5\x00\x32\x58\x41\x41\x29\x13\xac\xa8\xba\x3a\x33\xa5\xb6\x56\x89\x4b\x41\xa1\x96\xab\x96\x0b\x10\x00\x00\xff\xff\xf7\x3f\x36\x89\x74\x00\x00\x00")

func fixturesAliasAliasJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesAliasAliasJson,
		"fixtures/alias/alias.json",
	)
}

func fixturesAliasAliasJson() (*asset, error) {
	bytes, err := fixturesAliasAliasJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/alias/alias.json", size: 116, mode: os.FileMode(420), modTime: time.Unix(1551658297, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesGroupAdvanceGroupJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x41\x4f\xc3\x30\x0c\x85\xef\xfd\x15\x56\xce\x50\x0d\xd6\x03\xec\xc6\x8d\x5d\x00\xa9\xfb\x03\x5e\x66\x75\x96\x52\x27\x4a\xdc\x89\x6a\xea\x7f\x47\x49\xd7\x51\x38\xe6\xbd\xe7\xe7\xcf\xb9\x56\x00\x46\xc7\x40\x66\x07\xa6\x8b\x7e\x08\xe6\x21\x4b\x43\xa2\xb8\x3f\x65\x51\xce\x3c\x4b\xc5\x9d\x35\xc7\xc7\x88\x91\x29\xcd\x8e\x46\x64\x4d\x66\x07\xb9\x0d\xc0\x08\xf6\xa5\x6f\x2f\xac\x64\xcf\x25\x04\x60\x58\x4e\x43\xd2\x38\x66\xeb\x40\xf6\x2c\xde\xf9\x6e\x5c\x5c\xea\x83\xf3\x23\x51\xee\xd9\x3e\xbf\xde\xd4\xe0\x50\x72\x9e\x44\x29\x86\xc8\x89\x96\xbc\x7a\x45\x07\x47\x76\x8e\x32\xd3\xcb\x76\xb3\xac\x29\x88\xd7\x2b\x9f\xa6\xc9\x54\x00\x53\x61\xb4\x5e\x94\xbe\xf5\x17\x12\x43\xb8\x3f\x00\xcc\x71\x60\x57\x06\x9b\x66\x73\x5b\xb1\xba\xa4\xa5\xae\x27\x51\x68\xb1\x0f\x8e\xfe\xfa\x29\xa0\x2d\x21\xeb\xfb\x3a\xcd\xc1\x1a\x05\xdd\xa8\x6c\x53\x9d\xfe\x8d\x5c\x28\x26\xf6\xe5\xa8\xa6\x6e\xea\xcd\x63\xfb\xf1\xf6\xd5\xbe\x7f\x1e\x4c\x49\x4c\x77\x62\x16\xa5\x2e\xa2\xb2\x97\xd5\xdf\xae\xd4\xa7\x35\xff\x42\x7a\x41\x37\xd0\xaa\xaa\x9a\xaa\x9f\x00\x00\x00\xff\xff\xab\x88\xe5\xa7\xe4\x01\x00\x00")

func fixturesGroupAdvanceGroupJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesGroupAdvanceGroupJson,
		"fixtures/group/advance-group.json",
	)
}

func fixturesGroupAdvanceGroupJson() (*asset, error) {
	bytes, err := fixturesGroupAdvanceGroupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/group/advance-group.json", size: 484, mode: os.FileMode(420), modTime: time.Unix(1551656475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesGroupGroupJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8e\x31\x8e\x83\x30\x10\x45\x7b\x9f\x62\x34\xf5\x16\xab\xa5\xd9\x70\x03\xfa\x5c\xc0\x81\x2f\xb0\x62\x6c\xcb\x1e\x0a\x0b\xf9\xee\xd1\x98\xd0\xbe\x79\x7a\xf3\x4f\x43\xc4\x52\x13\x78\x24\x5e\x73\x3c\x12\xff\x28\x3a\x0a\xf2\xb4\x28\x4c\xd9\x0a\xf0\xbe\x70\x37\x2e\x5e\xb0\xee\x08\x72\x71\xc9\xd6\x49\xe1\x91\xb4\x47\xc4\xc1\xee\xbd\x38\x05\x27\x98\xb7\x2e\x11\xb1\x0b\xcb\x51\x24\x57\x3d\x3d\x31\x6f\x21\xfa\xb8\xd6\xfb\x8a\x3d\xf9\x58\x01\xed\x0c\x7f\x8f\x2f\x4d\xde\x06\xf5\x11\x04\x39\x65\x57\x70\xfb\x12\xc5\x7a\x7a\x39\xef\xa1\x8b\xfe\x87\xdf\xfb\x4d\x1f\x78\x9e\x6e\x69\x8d\x0d\x51\x33\xcd\x7c\x02\x00\x00\xff\xff\x2f\x3b\x70\x18\xe9\x00\x00\x00")

func fixturesGroupGroupJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesGroupGroupJson,
		"fixtures/group/group.json",
	)
}

func fixturesGroupGroupJson() (*asset, error) {
	bytes, err := fixturesGroupGroupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/group/group.json", size: 233, mode: os.FileMode(420), modTime: time.Unix(1551640808, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesIdentifyAdvanceIdentifyJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\x5d\x6b\xf2\x40\x10\x85\xef\xf3\x2b\x96\xbd\x7e\x79\x89\x2e\xc1\x8f\xab\x42\xad\xc1\x52\x63\x05\x7b\xd1\xde\xad\x71\xaa\x63\xf6\x23\x24\x93\x62\x08\xfb\xdf\xcb\x6e\x4c\x10\x7a\x7b\xe6\xcc\x79\xce\x4c\x17\x31\xc6\xa9\x2d\x81\x2f\x19\xc7\x13\x18\xc2\xef\x96\xff\xf3\x6a\x53\x43\xb5\x39\x79\x3d\x3b\x6c\xc4\x6e\xb5\x17\x5f\xfa\x83\x76\xab\xf3\x2d\x7b\x8e\xe3\xed\xf5\x82\x6f\x87\x22\xfe\xbc\xe6\xb4\x9d\xbe\xaa\x2c\xdd\x27\xbb\x74\xad\xb2\xf4\x45\xf4\xfb\x54\x49\xa4\x9a\x2f\x99\x67\x30\xc6\x8d\xd4\x81\xf2\x0e\x04\x15\x4b\xf1\x78\xb4\xa6\x0e\x56\xc6\x38\x68\x89\xca\x4f\x4b\x3f\x7d\x42\x83\x04\xf9\xe5\x7f\x6e\xf5\xe0\x28\x95\x34\xc1\x50\x81\xc6\x66\x94\x31\x34\xec\x3a\x3c\x39\xc7\x23\xc6\x5c\x60\xa3\x21\x38\x57\x92\xd0\x23\xc6\x06\x6b\x99\xc3\xd1\xda\xc2\x2b\xee\xbe\x9f\x5a\x7b\x56\x30\x7a\xee\x3d\x27\x3e\xf4\x47\xaa\x06\x26\x3c\xe8\x6e\x8c\x96\xc6\x9a\x56\xdb\xa6\xee\x7f\xb3\x88\xc5\x7c\x96\x88\x69\x32\x5f\xcc\xc4\x74\x91\xcc\x67\xfd\xf5\xb9\x35\x04\x37\xfa\x7b\x7e\x48\x1d\xda\x6b\x59\x3e\xa2\x0b\x68\x07\xf2\x23\x37\x72\xd1\x6f\x00\x00\x00\xff\xff\x95\x96\xdb\x41\xa8\x01\x00\x00")

func fixturesIdentifyAdvanceIdentifyJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesIdentifyAdvanceIdentifyJson,
		"fixtures/identify/advance-identify.json",
	)
}

func fixturesIdentifyAdvanceIdentifyJson() (*asset, error) {
	bytes, err := fixturesIdentifyAdvanceIdentifyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/identify/advance-identify.json", size: 424, mode: os.FileMode(420), modTime: time.Unix(1551657365, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesIdentifyIdentifyJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcd\x31\x0e\xc2\x30\x0c\x05\xd0\x3d\xa7\xf8\xf2\x8c\x38\x40\x27\x36\xc4\xc6\x15\xd2\xc6\x08\x8b\xc6\x8d\x12\x77\xa8\xa2\xdc\x1d\x25\x15\xac\xcf\xff\xfb\x57\x07\x90\x1d\x89\x69\x02\x49\x60\x35\x79\x1d\x74\xe9\xba\x17\xce\x8f\xd0\x3d\x65\x6f\xcc\x9f\x93\x2d\x7b\xb1\x42\x13\x7a\x15\x20\xf5\x71\x94\x9f\x6c\x9c\x71\x97\x79\xde\xb4\x8c\x28\x40\x1c\xbd\xac\xe3\x45\xbf\xde\x44\xc5\x78\x79\x5f\x97\x2d\xfe\x12\x69\xf5\x7a\x6e\x70\x94\xfd\xcf\x32\x86\x6b\x95\xd0\x1a\x39\xa0\xb9\xe6\xbe\x01\x00\x00\xff\xff\x9a\xe7\x17\x53\xac\x00\x00\x00")

func fixturesIdentifyIdentifyJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesIdentifyIdentifyJson,
		"fixtures/identify/identify.json",
	)
}

func fixturesIdentifyIdentifyJson() (*asset, error) {
	bytes, err := fixturesIdentifyIdentifyJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/identify/identify.json", size: 172, mode: os.FileMode(420), modTime: time.Unix(1551640808, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesPagePageJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8d\x31\x0e\xc3\x20\x0c\x45\x77\x4e\x61\x79\xae\x92\x9d\x13\x34\x27\xe8\x8c\x88\xd5\x58\x85\x80\x88\x2b\x54\x51\xee\x5e\xe1\x34\xeb\xf3\xfb\xcf\xcd\x00\xa0\x7c\x32\xa1\x05\xcc\xee\x49\x78\x1b\x64\x77\x51\xc9\x3d\xc5\x3f\x79\x1f\x54\x96\x55\xad\xe2\x84\xe8\x75\xe2\x5c\x52\xa6\x22\x4c\x07\x5a\x18\xb1\x91\x63\x09\xba\x7e\x50\xf0\x29\x12\x7c\x61\xd9\x59\xc8\x6f\xba\x19\xb1\x12\xc6\x7d\x13\xc9\x76\x9e\x6b\xad\x13\x9f\xc2\xe4\x53\xbc\x24\xd6\x6f\xad\xf1\xda\x3b\x1a\x80\x6e\xba\xf9\x05\x00\x00\xff\xff\xff\x51\xc4\xf5\xaf\x00\x00\x00")

func fixturesPagePageJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesPagePageJson,
		"fixtures/page/page.json",
	)
}

func fixturesPagePageJson() (*asset, error) {
	bytes, err := fixturesPagePageJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/page/page.json", size: 175, mode: os.FileMode(420), modTime: time.Unix(1551640808, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesScreenAdvanceScreenJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\xc1\x6e\x9c\x30\x10\xbd\xef\x57\x58\x9c\x0b\x02\x03\x1b\x88\xd4\x43\xda\xf4\x90\x43\xa3\x2a\x55\xa5\xde\xa2\xa9\x3d\x4b\xac\x05\xdb\x32\x66\x95\xed\x8a\x7f\xaf\x6c\xf0\xc2\x52\xc5\x27\xfc\xe6\xcd\x9b\x99\xc7\xf8\xb2\x23\x24\xb2\x67\x8d\xd1\x3d\x89\x7a\x66\x10\x65\xf4\xc9\x61\x12\x3a\x8f\x3d\xa2\x05\xd1\x92\xef\x28\x07\xf2\x03\x1a\x24\x4f\x16\xbb\x89\x32\xf4\x68\x9e\xb8\x23\xc9\x37\x31\x41\xda\x28\x8d\xc6\x0a\xec\xa3\x7b\xe2\xb4\x09\x89\x84\xa7\x5c\x2e\x82\x8f\xa3\x67\x11\x12\x31\xe8\x34\x88\x46\xbe\x02\xe7\xc2\x0a\x25\xa1\xf5\x3a\x8a\xac\x80\xff\xb8\x0d\x4a\x76\x0e\xbc\xe9\xb2\xe5\x74\x60\x8e\x68\x85\x6c\x5e\xd9\x1b\x48\x89\x41\xf6\x8a\x93\x80\x6f\x33\xc3\xc0\x52\x11\xff\xb9\x8d\x2b\xd3\x80\x14\x7f\xc1\x35\x37\xf3\x6e\xa0\x2d\xbf\x57\x83\x61\x41\x71\xbe\x6c\x39\xd6\x00\x3b\xfa\x66\x15\x0f\xd4\x80\x11\x8f\xed\x08\x19\xbd\xb3\x20\x95\x3c\x77\x6a\xe8\x27\xc7\xdd\xf5\xf3\xc3\x97\xaf\x8f\x93\xef\x4c\x49\x8b\xef\x76\x31\x1d\xb4\x76\xb4\x67\xb4\x87\x56\xbc\x93\xe7\x9b\x1e\x41\xeb\x65\xdc\x89\xb1\x0e\xe9\x16\xec\x41\x99\xce\x85\x8d\x3a\x0e\xeb\xd8\x09\x4d\x3f\x1b\x50\x24\x34\xc9\x42\x8c\xe3\x49\xf8\x69\xa7\xfa\x8e\xcd\x4f\x6e\x13\x7a\x21\x9b\x69\x03\x32\x9a\x17\x65\x5c\x94\xfb\xbb\xb8\xae\xaa\x7d\x5c\x65\x55\x15\xd3\x7d\x91\x16\x45\x51\x50\x4a\xe9\xac\x75\x5d\x99\xba\xe0\x99\x3b\x71\x9e\xe7\x79\x5c\x96\x65\x19\x03\x00\xc4\xe8\x4e\xce\x52\x77\x96\x9c\x0e\xe4\x70\x00\x66\x07\x83\xc6\x65\xbf\x2c\x9d\xbb\xa8\xe2\xd3\x2e\xe4\x77\x59\xfa\x7b\xc1\xc3\xea\x9b\x1b\xf6\x6a\x4a\x5a\xd5\x49\x9a\x7e\x4b\x8b\xb4\x2a\x1e\x22\x4f\x18\xe7\x99\x85\xb7\x38\xab\x69\x92\xed\xab\x24\xab\xeb\x24\x4b\xaf\x86\xb4\x8a\x41\xeb\xa5\x51\xc6\xbf\x7e\x06\x58\xf5\x6b\x8f\xc2\x4f\x78\xf9\xa8\x7a\x9d\xa4\xe4\xcf\x20\x5a\x4e\x5c\xfd\xdb\xf2\xd6\x80\xb0\x37\x72\xd0\xf7\x8a\x09\xb0\xc8\xe7\x67\xa9\x88\xe0\x73\xd2\x75\x91\x84\xb4\xd8\x18\xbf\x0f\xeb\x47\xba\xa0\xd9\xb6\x43\x07\x44\x27\x68\x07\xcc\x36\x06\x2c\x49\x74\x9d\xd4\x81\xf6\x22\xe3\xea\xe7\xe8\x35\x63\x16\xa6\x57\x61\x1a\xcd\x91\x71\xe9\x76\x37\xee\xfe\x05\x00\x00\xff\xff\x9b\x1e\x2e\x13\x9e\x04\x00\x00")

func fixturesScreenAdvanceScreenJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesScreenAdvanceScreenJson,
		"fixtures/screen/advance-screen.json",
	)
}

func fixturesScreenAdvanceScreenJson() (*asset, error) {
	bytes, err := fixturesScreenAdvanceScreenJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/screen/advance-screen.json", size: 1182, mode: os.FileMode(420), modTime: time.Unix(1551657737, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesScreenScreenJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x4f\xca\x03\x21\x0c\x05\xf0\xbd\xa7\x08\x59\x7f\xcc\xec\x3d\xc1\x37\x27\xe8\xba\x38\x01\x03\xfe\x43\x53\xa4\x58\xef\x5e\xa2\x74\xfb\xcb\x7b\x2f\xc3\x00\xa0\xbc\x0b\xa1\x05\x6c\xae\x12\x25\xfc\x53\x4b\xcf\xb8\xec\x3f\x47\xda\xf2\x6a\x54\xaf\x5b\x2d\x79\xde\x54\x6a\x2e\x54\x85\xa9\xa1\x05\x9d\xd2\x31\x96\xb0\x9a\x0f\x0a\x2e\x47\x82\x0f\x5c\x89\x85\x9c\x5f\x1d\x1d\xaa\x41\xef\x5e\xa4\xd8\xf3\xec\xbd\x1f\xbc\x03\x87\xcb\xf1\x17\xe2\xf5\x69\x0c\xbe\xe7\x44\x03\x30\xcd\x34\xdf\x00\x00\x00\xff\xff\x71\xd0\x49\x1a\xad\x00\x00\x00")

func fixturesScreenScreenJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesScreenScreenJson,
		"fixtures/screen/screen.json",
	)
}

func fixturesScreenScreenJson() (*asset, error) {
	bytes, err := fixturesScreenScreenJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/screen/screen.json", size: 173, mode: os.FileMode(420), modTime: time.Unix(1551657600, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesTrackAdvanceTrackJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\xcb\x6e\xdb\x3a\x10\xdd\xfb\x2b\x08\xae\x23\x5d\x52\xb2\x0d\xcb\x3b\x23\x09\x2e\x0a\xa4\x41\x11\xb7\x45\xd1\x8d\x41\x52\x63\x87\x88\x44\x0a\x24\xa5\xd4\x08\xf4\xef\x05\xa9\xa7\x9d\x04\xd5\x4a\x9c\x33\x67\x1e\x67\x86\x7c\x5b\x20\x84\x99\xd2\xea\x5c\xea\xda\x7e\xc9\xf1\xb6\x3b\x52\x42\xf0\x8d\xc7\x84\x56\x0e\xfe\x38\xbc\x45\xde\xd5\x3b\x57\x95\x77\xda\xbd\x82\xd5\x25\xa0\x9f\x32\x07\x8d\x76\x55\x15\xdc\x3b\xfc\xa0\x58\x09\x21\x52\xc3\xe6\xe6\xaa\x60\xee\xa8\x4d\xe9\x21\xa3\x5f\xea\x39\xd6\x80\xb1\x52\x2b\x0f\xa5\x31\x4d\x63\x32\x80\x39\x34\x52\xc0\x98\xdf\xbb\xe7\x0d\x18\x27\xad\x54\x27\xd9\x15\xcc\x18\xf1\x5f\xc4\x38\xe7\xd1\x2a\x4d\xd3\x68\xe3\x8f\xb0\xe2\x22\xe7\x99\x20\x9c\xf3\x3e\x1c\x42\xb8\xe3\x50\x66\x08\xa5\x94\x46\x42\x08\x11\xad\x56\xab\x55\x94\x65\x59\x16\xad\x39\x5d\x6e\x36\x22\x7c\x13\xa7\x64\xaa\x3e\x32\xe1\x6a\x03\xc6\xb3\x9f\xa6\xea\x3d\xaa\x73\x28\xbc\x79\xb9\x4e\xc8\xaf\xc9\xee\xce\x15\x5c\xf5\x8a\x10\x9e\x75\x9a\x64\x59\x4c\xc8\x3d\x59\xd2\x65\xb2\xc3\xc1\xa1\xed\xdb\x96\x41\x65\x9a\x25\x31\x5d\x93\x98\xc4\x49\xb2\x1e\x14\x29\x24\x37\xcc\x9c\xe7\x92\x8c\x82\x2b\x56\x9c\x9d\x14\x36\xfa\x34\x27\x8d\x49\x4c\x2e\x73\x15\x5a\xb0\x22\xf0\x41\x45\x3f\xf6\x43\x1e\x6d\x3f\x4a\xf1\xf4\x59\xe0\x2c\x26\x88\xd7\xb2\xc8\x91\x6f\xe7\x32\x43\x65\x64\xc9\xcc\xf9\xc0\x6b\x2b\x15\x58\x7b\xa8\x95\xf4\x3b\x85\x15\xbc\xda\x21\x9f\x05\xa1\x55\xfe\x4f\x37\x67\x98\x78\x91\xea\x74\xc7\x1c\x9b\x17\x78\xbd\xa9\xef\x64\xdc\xc4\x24\xa6\x37\x68\x3a\x2c\x97\x63\x23\xbd\x0a\xae\xeb\x65\x0a\xe0\x17\xb0\xf4\x69\xf0\x86\xa4\x78\x34\xb7\x33\x9e\x93\x25\xfc\xd6\x2a\x88\xb3\x2b\xc1\x48\xc1\xfe\x7b\xd0\xf6\xb0\x53\x27\x28\xc0\x0e\xa4\x91\x82\x2b\xa3\x2b\xbf\xc0\x60\x2f\x6b\xbd\xbd\x7b\xf4\x31\xee\xef\xfe\xbf\xbf\xdd\xed\xbf\xcf\x6b\x93\xf6\x41\x36\x3e\xc3\x91\x15\x16\xde\x45\xac\x2d\x98\xee\xee\xfa\x3f\x9a\xa4\x97\xea\x3b\xc3\xa4\xbb\x18\x26\x28\x27\x5d\x01\x25\xa8\x60\xc7\x8c\x8b\xbc\xe7\x2c\x7a\x1e\x86\x06\x54\x10\xff\xd6\x0b\xab\x1c\xda\x3b\x66\x1c\xe4\xdd\xbb\x30\xec\x76\x98\x46\x67\x9a\xaa\x50\xcf\xb2\x33\x49\xe5\xe0\x64\x82\xac\x53\xfe\xb9\x95\x5e\xaf\x98\x37\xe0\x86\x15\x35\xd0\xab\x0b\x31\x91\x92\x39\xa9\x64\x55\x08\xd2\xce\x2e\x6b\x95\x5c\x2a\xeb\x03\x27\x63\xe0\x64\x1c\xc9\x55\xc7\x1f\x0c\xa6\x7f\x2d\xde\xde\x64\xde\xb6\xc3\x06\x3e\xc1\x11\x8c\x01\x83\x1e\xfb\x3b\xf1\xe8\xc7\x3f\x82\xb6\xd3\x09\x7d\x2b\xd8\x99\x7b\x75\x82\xc7\x0c\xaf\xcb\x4f\xd1\xfe\x39\xcd\xd1\x03\xa8\x93\x7b\xc6\x5b\x44\xae\x91\xaf\xc3\x6b\xf3\x8e\xb3\xaf\xab\x4a\x87\x11\x6d\x91\x33\xb5\xdf\x93\x76\xd1\x2e\xfe\x06\x00\x00\xff\xff\xbb\x34\xde\xca\xe4\x05\x00\x00")

func fixturesTrackAdvanceTrackJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesTrackAdvanceTrackJson,
		"fixtures/track/advance-track.json",
	)
}

func fixturesTrackAdvanceTrackJson() (*asset, error) {
	bytes, err := fixturesTrackAdvanceTrackJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/track/advance-track.json", size: 1508, mode: os.FileMode(420), modTime: time.Unix(1551657117, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fixturesTrackTrackJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x31\x0a\x02\x31\x10\x45\xfb\x9c\xe2\x33\xb5\x27\xd8\xce\x46\xb0\x13\xdc\x0b\x8c\xd9\x29\xc2\xae\x49\x98\x9d\x08\x12\x72\x77\x99\x45\xb1\x7d\xef\xf1\xf9\x3d\x00\x64\xef\x2a\x34\x81\x4c\x39\xae\x74\x72\x24\x2f\xc9\xe6\xec\x9e\x9e\x75\x13\xcc\x7f\xd5\x76\xd1\xeb\xe2\xae\x2a\x9b\xc8\x17\x57\x2d\x55\xd4\x92\xec\x34\xc1\x67\x01\x4a\x47\xd6\x7b\x5a\xc6\x38\x2a\xef\x36\xce\x4e\x6f\x5a\x70\xce\xb9\xf1\xf6\x33\x1c\x63\x69\xd9\x66\x7f\x83\x09\x74\xe1\x28\x8f\x52\x56\x0a\xc0\x08\x23\x7c\x02\x00\x00\xff\xff\xad\x23\x77\x22\xac\x00\x00\x00")

func fixturesTrackTrackJsonBytes() ([]byte, error) {
	return bindataRead(
		_fixturesTrackTrackJson,
		"fixtures/track/track.json",
	)
}

func fixturesTrackTrackJson() (*asset, error) {
	bytes, err := fixturesTrackTrackJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fixtures/track/track.json", size: 172, mode: os.FileMode(420), modTime: time.Unix(1551640808, 0)}
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
	"fixtures/alias/advance-alias.json": fixturesAliasAdvanceAliasJson,
	"fixtures/alias/alias.json": fixturesAliasAliasJson,
	"fixtures/group/advance-group.json": fixturesGroupAdvanceGroupJson,
	"fixtures/group/group.json": fixturesGroupGroupJson,
	"fixtures/identify/advance-identify.json": fixturesIdentifyAdvanceIdentifyJson,
	"fixtures/identify/identify.json": fixturesIdentifyIdentifyJson,
	"fixtures/page/page.json": fixturesPagePageJson,
	"fixtures/screen/advance-screen.json": fixturesScreenAdvanceScreenJson,
	"fixtures/screen/screen.json": fixturesScreenScreenJson,
	"fixtures/track/advance-track.json": fixturesTrackAdvanceTrackJson,
	"fixtures/track/track.json": fixturesTrackTrackJson,
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
	"fixtures": &bintree{nil, map[string]*bintree{
		"alias": &bintree{nil, map[string]*bintree{
			"advance-alias.json": &bintree{fixturesAliasAdvanceAliasJson, map[string]*bintree{}},
			"alias.json": &bintree{fixturesAliasAliasJson, map[string]*bintree{}},
		}},
		"group": &bintree{nil, map[string]*bintree{
			"advance-group.json": &bintree{fixturesGroupAdvanceGroupJson, map[string]*bintree{}},
			"group.json": &bintree{fixturesGroupGroupJson, map[string]*bintree{}},
		}},
		"identify": &bintree{nil, map[string]*bintree{
			"advance-identify.json": &bintree{fixturesIdentifyAdvanceIdentifyJson, map[string]*bintree{}},
			"identify.json": &bintree{fixturesIdentifyIdentifyJson, map[string]*bintree{}},
		}},
		"page": &bintree{nil, map[string]*bintree{
			"page.json": &bintree{fixturesPagePageJson, map[string]*bintree{}},
		}},
		"screen": &bintree{nil, map[string]*bintree{
			"advance-screen.json": &bintree{fixturesScreenAdvanceScreenJson, map[string]*bintree{}},
			"screen.json": &bintree{fixturesScreenScreenJson, map[string]*bintree{}},
		}},
		"track": &bintree{nil, map[string]*bintree{
			"advance-track.json": &bintree{fixturesTrackAdvanceTrackJson, map[string]*bintree{}},
			"track.json": &bintree{fixturesTrackTrackJson, map[string]*bintree{}},
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

