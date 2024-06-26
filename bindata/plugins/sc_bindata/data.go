// Code generated by go-bindata. DO NOT EDIT.
// sources:
// scripts/check-step.sh (7.585kB)
// scripts/set-dcgm-exporter.sh (686B)
// scripts/set-systemctl-service.sh (796B)

package sc_bindata

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _scriptsCheckStepSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\xeb\x6e\xdc\xc6\x15\xfe\xbf\x4f\x71\xca\x5d\x74\x93\xd4\x43\xea\x62\xab\xc1\x26\x32\xa0\xf8\x22\x0b\x88\x6d\x41\xb6\xd2\x16\x8e\x21\xcc\x92\x87\xcb\xa9\xc8\x19\x7a\x2e\xbb\x12\x2c\x03\xc9\x6b\xf4\x59\xfa\x36\x79\x80\xbe\x42\x31\x43\x0e\x97\x5c\x72\xe5\x4d\xeb\x56\x3f\x12\x6b\x2e\xdf\x9c\xcb\x37\xdf\x9c\x43\x8d\xff\x10\x19\x25\xa3\x39\xe3\x11\xf2\x25\xcc\xa9\xca\x46\x23\x85\x1a\x08\x56\xff\x13\x50\xb2\x12\x53\xca\xf2\xd1\x88\x71\xa5\x69\x9e\x9f\x1a\x96\xe0\x57\x5f\xc3\xc7\x11\x40\x21\x12\xcc\x8f\x83\xc9\x7e\x30\x02\x88\xa9\x42\x08\x26\x6e\x2c\x00\xc6\x47\xe0\x7e\x82\xa7\x22\xbe\x46\x19\x7c\x5d\xff\x6e\x7f\x30\xce\x84\x9f\x80\xb3\x0a\x97\x6a\x26\x38\x38\xf4\x59\xd0\x5b\xbb\xff\xdb\x2f\xbf\x5e\x2a\x84\x9f\x03\x5a\x6a\xa8\x4d\x81\xc4\x21\x84\x4c\xfc\x1c\x40\x2a\xa4\x1f\x77\x50\x0f\x40\x48\x90\x98\xa2\x04\x2d\x40\x67\x08\x22\x4d\x59\xcc\x68\x0e\xf5\xc1\x89\x88\x4d\x81\x5c\x57\x27\x6f\xee\x0f\xe1\x6f\xc2\x00\x47\x4c\xec\x7e\xa3\xd0\x6f\x5b\xa2\x54\x76\xc3\x63\x38\x78\x18\xee\x85\x7b\x61\xdf\x5a\x00\x52\xdb\xd6\x3d\xe4\x5f\xff\xfc\x47\xa6\x75\xa9\x66\x51\x94\x88\x58\x85\xb5\xfd\xb1\x28\x22\xe4\x0b\xc6\x31\xaa\x2d\x88\xcc\xdc\x70\x6d\xa2\x3e\xf4\xc1\x6f\xbf\xfc\x7a\x96\x82\xce\x98\x02\xa6\x80\xc2\xe9\xf9\x25\xbc\x10\x4a\x3f\x80\xf3\x1c\x6d\x0a\x3a\x3e\xa7\x22\xcf\xc5\x8a\xf1\x45\x63\x88\xb2\x53\x0b\x61\xff\x5b\x3b\x74\x7a\x7e\x79\x61\xb8\x66\x05\x6e\xf5\x64\xbb\x0b\x7c\xc9\x12\x46\x9d\x0b\x09\xd5\x34\x46\xae\x51\x46\x71\x2e\x4c\x42\x38\xd5\x6c\x89\x51\x2c\xb8\xa6\x8c\xa3\x24\x5a\x88\xfc\x9a\xe9\x28\xa7\x1a\x95\xf6\xce\x92\x85\x4d\x7a\x98\xe9\x22\x6f\x1b\xf0\xdd\x77\x9e\x41\xa7\xe7\x97\x4f\x25\x5b\x0e\x92\xa8\x99\xdb\x8d\x47\xaf\x4b\x37\xbd\x3f\x83\x13\xa3\x45\x41\x35\x8b\x3b\x59\x07\xa3\x6c\xb0\x6c\xe8\x94\xa6\x3c\xa1\x32\x81\x4b\x97\x0b\x90\x58\x0a\xc5\xb4\x90\xb7\x7d\xd8\x2a\x5d\x24\x71\x96\x28\x48\x70\xc9\x62\x54\xfd\x75\xca\x24\x02\x36\x16\x53\xa3\x45\x6d\xc2\x56\x7b\x0f\x66\xf0\x92\x72\x43\xf3\x6d\xc6\x36\xdc\x7e\xf5\xd3\xd9\xd3\xb3\x13\xa8\xc0\xfb\x78\x4f\xc5\x8a\xe7\x82\x26\x6e\x93\xa5\x4e\xb5\x10\x52\x29\x0a\x37\x56\xef\x5f\xe1\x5c\x31\x8d\x33\xf0\xc9\x5e\xad\x56\xed\x5c\x7b\x9c\x88\xf1\x04\x6f\x42\xaa\xca\x9b\xfe\x61\xab\x85\xd5\x90\xd8\x61\xcc\xa2\xc8\x58\xc2\x57\xdb\xda\x50\x7f\x7d\x2e\x11\xbf\x3d\x8a\x7e\x64\xdc\xdc\x90\x9b\x6f\x8f\xae\x8e\x1e\x46\x8f\x0e\x1f\x85\xdf\x7c\x13\x55\xd6\x90\xf6\x14\xa9\xa6\x42\x69\xf8\x96\xf8\xc6\x59\x21\x12\xa0\x7f\xba\x81\xa1\xed\x0f\x1f\xee\x85\xdf\x1e\xdc\xb3\x3d\x1c\x3c\x75\xbd\x0d\x08\x17\x44\x94\xc8\x17\x39\x49\x59\x8e\xca\x0d\x70\x61\x96\x48\x0d\x89\x33\x8c\xaf\x87\x69\xfc\xf4\xc9\xe9\xcb\x0e\x83\xc7\x10\x8b\xa2\xa0\x3c\x51\xfd\x3c\x3d\x39\x7d\xb9\x1b\xa3\x87\x6e\x7d\xc3\x86\x0e\x5b\xfa\x7a\x67\xd7\x4a\xfc\x60\x98\xc4\xa4\xd1\xb5\xc3\xf0\x30\x7c\x64\xc5\x93\xce\xc5\x12\x07\xd4\xad\xb9\xff\xb8\xc4\x5c\x94\x28\x3b\x22\x10\x2f\x8a\xb1\x67\x87\xda\x72\x9f\x7f\x3c\xfb\xe1\xc9\xae\x91\x70\x8b\x77\x7c\x24\x42\x78\x5d\x22\x77\x6e\x45\xa8\xe3\x88\x96\x3a\x52\xc2\xc8\x18\x55\x98\x33\xa5\xc1\xe6\x6b\x60\x23\x00\x2c\xd9\xf0\x96\x01\x05\x0e\xe1\x24\x49\x40\x70\x1b\xe7\x0d\x91\x5d\x6b\x04\x20\xd7\x92\xa1\x1a\x3e\x8d\x40\x82\x73\x7f\x2f\x14\xc6\x46\x32\x7d\x1b\x56\xc2\xe0\xc2\x58\xfd\x13\xfe\x4e\x8b\xe2\x96\xf8\x05\x50\x50\x36\xc0\xda\x4d\x3c\x9d\x85\x54\xc6\x19\x5b\xe2\x36\xc4\x2d\x40\x87\x21\x5c\x96\x09\xd5\xe8\xbc\x2a\x69\x7c\x4d\x17\x08\x36\x0a\x5b\xdc\xb0\x6f\xb1\x71\x3b\xfa\xd3\x0f\x2d\xd8\x42\xd2\xa4\x42\xab\x99\x88\x89\xc7\xbd\x17\xd2\xed\xeb\xcf\x3f\x0a\xe1\x02\x0b\xb1\x44\xa0\xfc\x16\x0c\x37\x6a\x27\x3c\xab\xb0\xd2\xed\xeb\x2f\x39\x0a\x3d\xb7\x2a\x49\x74\x6c\xab\x21\xb7\x23\xfa\x02\x24\x67\xf3\xf8\xa8\xbf\xea\xcf\x21\x3c\xb1\x3a\xd0\x82\xf4\xb7\x6b\x7e\x0b\xd2\x70\xce\xf8\x62\x18\x5d\x69\xc9\xf8\x42\x41\x94\xb3\x79\x54\x6b\x4f\xee\x84\x68\xc1\x8d\x1d\x8c\x43\x25\xc2\x23\xb8\x83\x85\xc4\xb2\x02\x1f\x04\x7a\x9b\x61\x73\xa8\xca\x84\xc9\x13\x98\xa3\xdd\x44\xb5\x15\x8a\x8c\x72\x38\x08\x0f\x1f\x85\x5b\xb4\xca\x95\x06\xad\xca\x60\x57\xdd\xda\xd8\xb7\xf3\xc5\xfd\x01\x53\x21\x1b\x9a\xf8\x67\x2d\x69\x2a\x14\x1b\x35\x0b\xf8\x00\xca\x4a\xf0\x9c\xd0\x02\x4b\xd7\x45\xa0\xad\x87\x1a\x9a\x85\x70\x96\x02\x17\xba\x59\xef\x05\x72\xfd\xa4\xfd\xc7\x25\x58\x08\xcf\x6b\xf1\xec\xb9\xdb\x11\xdc\x02\x75\x26\x92\x9e\x05\xed\x77\xb6\x51\xea\x8e\x38\x3f\x80\xff\x53\x8d\xd5\xdc\xfb\x93\xb4\x62\xc5\x86\x07\x4c\xd9\x54\x97\x39\x6a\x7b\xe5\x12\x9f\x0f\xa6\x40\xa2\xd2\x54\x6a\x6c\xb9\x67\xf8\x86\x1a\x7a\x96\x58\x9f\x9b\x74\xf9\xc2\x43\xfa\x88\x29\x50\x26\x8e\x51\xa9\xd4\xe4\xf9\xed\x3a\x83\x5b\x84\xee\x8d\x3d\x16\x28\x34\x1e\x5b\x2e\x6b\x5b\xa0\x2b\x8b\x3b\xf3\x36\xba\x97\x3a\x01\xc2\x69\x81\x60\x43\x41\x16\xa5\x21\x36\xc5\x37\x40\x88\x2c\x80\xd8\x01\x05\xf6\x1a\xbb\xd1\x2d\xe7\x3d\xbb\xc1\xd8\xd4\x92\x58\x3b\xd4\xf8\xd3\x1c\x86\x37\x18\x03\x61\x7a\xf3\xa0\x2a\x73\x44\x15\xac\x8f\x1e\x0b\x29\x31\xd6\xa0\x34\x96\x70\x08\xa5\x64\x5c\xdb\xa8\x89\x75\x8c\x18\x4f\x05\x14\x48\xb9\xba\x37\x35\x2d\xec\xe6\x06\x7f\xd3\xba\xb0\x63\x48\x30\xa5\x26\xd7\x03\xf7\xb6\xb2\xe5\xd9\x12\xe5\xc9\x19\xb9\xc0\xea\xe9\x23\xaf\x44\xf2\xb9\xab\xeb\x4e\x42\x45\xe3\xd1\xa7\xd1\xc8\x45\xa3\xba\x0c\x75\x6b\x38\xae\x15\x90\xa5\xbe\xcd\x68\xdf\xce\x11\xd8\x09\x1f\x4e\xb2\xf4\x71\xfc\xe3\x63\xb0\x65\x45\xc4\x4d\x9e\x7f\x67\x7d\xf6\x5d\x64\xa7\x63\xec\xdc\x73\x6b\x12\xe6\x0a\x3b\x0b\xdf\x3d\xbb\xb8\x78\x7d\xf1\x7e\xd6\x3a\x9b\x0b\xbd\xb1\xcb\xfe\xb4\x9b\xda\x7a\x71\x3d\x23\x51\x1b\xc9\x61\x7f\x04\x90\xb2\x51\xcb\xa3\xf5\xbd\xf7\xfa\x3a\x02\xc8\x45\x4c\x7d\x3f\x7a\x55\x0f\x1f\x4f\xbe\xaa\xdd\x22\xc4\x2b\xf1\x1d\xd0\xd5\x35\x4c\x3f\xba\x6c\xc3\xe4\xf0\xd3\xd4\xe6\xa9\xe3\x5d\xbd\x74\x06\x93\x2e\x5c\x30\xea\x86\x55\xb7\x04\xbe\x26\x03\xd5\x6c\x9e\x23\xac\x98\xce\x06\x6b\xbb\xc6\x50\x3f\xd1\x98\x1a\x1c\x1c\xd9\x5e\x36\xa8\x12\xf3\xee\x1d\x04\x9b\xa7\xc3\xf7\x10\x4c\x36\xf7\x05\xf0\xfe\xfd\x40\x9e\x36\xc3\xef\xcd\xdc\xc0\xf4\x69\x59\x9b\x6e\xdf\xf7\x8d\x72\xd4\x36\xb9\x1a\xac\xc4\x68\xe8\x9d\xff\xbb\xf3\xd8\xfc\xbe\xd7\xd0\xb6\xe9\x20\xfb\xcc\xed\x34\x53\xf7\x13\x78\x7d\xd3\xb7\x91\xb8\x0a\xcd\x76\xc8\x2e\x93\xb7\xac\x1e\x62\x71\xc7\xf7\xc6\x9b\xd1\x67\x49\xdc\x05\xef\x71\xd9\x0d\xb7\xb8\xdc\xf2\x90\x64\x70\x07\x19\x52\x2b\xaf\xb0\xdf\x25\xf5\xe4\xe8\xd3\xb4\x1e\x21\xcf\x61\xba\x9c\x36\x13\x07\x6d\xb2\x0f\x1e\x6e\x39\xdf\x39\x76\x88\xf3\xdd\x2d\x5f\x80\xfa\xb6\xa1\xdc\xeb\x52\xbf\x6b\xc4\x2e\xd4\xbf\xc7\xab\x4d\xa7\xbe\x18\xed\x7f\x4f\xe2\x7b\xb4\xb7\x1d\xe6\x80\x56\xdb\xbe\xf3\x7e\xa5\x8e\x17\x9f\xe3\x78\x0f\x64\x88\xda\x7e\xd1\x67\x19\x6d\x17\x7e\x9e\xcc\x0e\xae\xc7\xe1\x78\x51\xb4\xd5\xd8\x99\x4e\x96\xbe\x78\x0e\x7e\xaa\x53\x79\xaf\x2a\xb7\x80\x2d\x3f\x5b\x90\x43\xec\x6c\x2f\xff\x02\xdc\x74\xcd\x78\x87\x9a\xed\xf3\x77\x27\x66\xc7\xac\x0e\xc8\xff\x86\x8d\xdb\x93\xd6\xd7\x5f\xdb\xc2\x6c\x30\xb1\xd7\x37\x35\x21\x5a\xd8\x16\xa8\x95\xd3\x9d\xfb\xa5\x75\xbb\x74\x75\xa7\x29\xcb\xad\x74\x1d\xdc\x35\x22\x76\xe7\x05\xeb\x6a\xba\x26\x42\x47\xb1\x3a\xf6\xcc\x60\xd2\xb1\x64\x88\x0a\xd5\x86\x2f\xf9\x46\x87\x87\x1d\x2e\x74\x2d\xd8\x9d\x0c\x1b\x86\x75\x71\xfe\x7b\x3e\x0c\x89\x93\x3d\x71\x80\x0f\xdd\xd2\x71\xdd\x47\x39\x3a\x8c\x9d\xa7\x49\x79\xbd\x20\x1f\x0c\xca\x5b\x20\x7f\x01\x92\x1e\x4f\x27\x1f\xdf\x68\xaa\x8d\xfa\x34\xf5\xaf\x6e\xaf\xfb\x81\x83\xc7\x8d\x3c\xf9\xdb\x4e\x3e\x40\xe0\xfb\x76\x71\xdd\xea\x33\xea\xf8\x8c\xdb\x15\xcc\x56\x60\xeb\x79\x2e\x91\x26\xb7\x1d\xd5\x1a\xd7\xfa\xb6\x33\x48\x57\xf6\x7c\xad\x6d\xeb\x7f\x2e\x56\x61\x58\x41\x6e\xa9\x6c\xd6\x71\x1a\x8d\xc7\xe0\x3e\x5a\xd2\x52\x93\x05\xae\x3f\x4c\x90\xdb\xad\xb1\xb1\xc0\xd5\x55\x7c\xfb\xec\xcd\xdb\xab\x57\x27\x2f\x9f\x1d\x07\x3a\x63\x8a\x30\xd5\x5a\x6d\x9b\x23\xe2\xda\x99\xba\x3b\x1a\xd7\x3d\x5c\x5d\xd3\xae\x32\xd4\x19\x4a\x50\xa8\x9b\xa3\xf4\xb5\x6f\xeb\x2a\x96\x6e\xb4\x62\xb6\xeb\x02\x52\x75\x64\x93\xe6\xf0\x5e\x23\xb6\x91\xbc\xcd\x16\xab\xb5\xb3\x55\x92\xb4\xf6\xf4\xf8\x5e\x23\xf4\xad\x6c\x67\x53\xa1\xae\x85\xac\x5e\x7e\xcd\xf2\xbc\x75\x58\xff\x05\xbb\x17\xd6\xe6\x57\xa1\xee\x7e\xae\x08\x1a\xd6\xa7\x86\xc7\xae\xaf\x72\xe3\xb5\xfa\x31\x75\x7a\x7e\xf9\x42\x28\x7d\x1c\x4c\x3e\xee\xcf\x88\x96\x06\x3f\x39\x65\x69\x5d\x92\x51\x2d\x00\x10\x4c\x9a\xf5\x01\x1c\x83\x5d\x0c\xed\xbb\xde\x2d\x6e\xd7\x43\x8d\x34\x0f\xde\x3c\xe7\xe6\x10\x7e\x4a\x73\xd5\x3d\x60\x0c\x2f\x28\x4f\x72\x04\x2e\x38\xb1\xed\x6a\x26\x94\x76\x7f\x06\x6c\x05\xe9\xa2\xfa\xd8\x05\x82\x03\xed\x2c\x1c\x28\x09\xce\xf8\x92\xe6\x2c\x81\x25\xcd\x0d\xba\x0f\xd4\x67\x6f\xae\x4e\xcf\x2f\xaf\x5e\xbc\x7e\xf3\x36\xf4\x7f\xdf\x32\x0a\x61\x6a\xbd\x9d\x82\x90\x30\x75\x76\x4d\xeb\xd4\xe1\x0d\xd3\xad\xb7\x66\xfd\xbe\xd8\x98\x8f\xe1\x89\xff\xcc\x57\x31\xb9\x49\x42\x4b\x8c\xab\x06\x18\x4a\xaa\xb3\x51\x1d\x86\x71\x00\x84\x23\xec\xb7\x9c\xaf\xec\xbd\x54\x74\x81\x33\x98\xec\xc1\xf7\x2d\x43\x1f\x07\x30\x6e\x1b\x5e\xfd\x99\x6e\x2e\x44\x8e\x94\x57\xbe\x3d\x68\x1a\x72\xa6\x5c\xe2\x46\x8d\xe9\xd6\xf0\xca\xba\x60\xb2\x1f\xfc\x3b\x00\x00\xff\xff\xff\xd1\x00\x98\xa1\x1d\x00\x00")

func scriptsCheckStepShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsCheckStepSh,
		"scripts/check-step.sh",
	)
}

func scriptsCheckStepSh() (*asset, error) {
	bytes, err := scriptsCheckStepShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/check-step.sh", size: 7585, mode: os.FileMode(0644), modTime: time.Unix(1719212044, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0x91, 0x3e, 0x7, 0xf6, 0xb2, 0x92, 0xc5, 0xa2, 0x27, 0x3a, 0x1c, 0x55, 0x92, 0x54, 0x52, 0x3f, 0xe9, 0x91, 0x9d, 0xb9, 0xc, 0xa2, 0x8e, 0xd, 0x22, 0xce, 0xbf, 0xeb, 0xd1, 0x23, 0x1a}}
	return a, nil
}

var _scriptsSetDcgmExporterSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x41\x8f\xd3\x30\x10\x85\xef\xf3\x2b\x1e\x4e\xa4\xc2\x21\xeb\xf6\xba\xb0\x20\xd8\x06\xb1\x07\x58\xa9\xcb\xc2\x01\xa1\xca\x75\x26\x89\xb5\xae\x1d\x6c\x47\xdb\x82\xf8\xef\x28\x69\xab\x85\x54\x22\x97\x68\x3c\x33\x9f\xdf\xcc\x73\xf6\x4c\x6e\x8c\x93\x1b\x15\x5b\xa2\xc8\x09\x05\x1f\x7e\x1e\x9d\xe9\xb8\x56\xc6\x12\x65\xa8\x62\x82\xe4\xa4\x29\x5b\xde\x7d\x5e\x2f\x6f\x56\x57\x62\x08\x05\x3d\x85\xf7\x91\x43\x94\x3f\x5b\xe5\x77\xc6\xc9\xaf\x3e\x3c\xc4\x4e\x69\x96\x8d\x49\x6d\xbf\xb9\xd0\x7e\x7b\x4a\x16\xef\xde\xcb\xca\xeb\x07\x0e\x45\xe2\x98\x06\x01\x95\x4a\x4a\x26\xde\x76\x82\x28\xee\x9d\x5e\xea\x66\x5b\xee\x3a\x1f\x12\x87\xe7\x2f\xf0\x8b\x00\xeb\xb5\xb2\x28\xbf\x94\xab\xb7\x37\xeb\x4f\xb7\xcb\x72\xfd\xe1\xf6\x63\x79\x25\xf2\x85\x20\x80\x75\xeb\x21\xa6\xd9\x4b\xe4\xd3\x23\x41\x04\xdc\xad\xae\x0f\xaa\xcf\xd2\xb2\xd2\xcd\xb6\xe0\xe3\xd5\x4f\xe4\x63\xc7\xd8\x9d\x41\x77\xa8\x4c\x20\x20\xf6\x95\x1f\xa2\x22\x40\xe4\xa7\x1a\x88\xfc\xb8\x96\xb1\xdc\xd4\xf8\x86\xfc\x0d\x0a\xfe\x81\x39\xbe\xbf\x44\x6a\xd9\x11\xc6\xef\x08\xef\xb5\xe6\x18\xeb\xde\xda\x3d\xb4\xef\x0c\x57\x98\x9d\x70\x33\x24\x8f\xd9\x89\x38\xbb\x18\x35\xd9\xc8\xff\x10\xca\x10\x7c\x80\xd7\xba\x0f\x81\x2b\x3c\xb6\xc6\xf2\x40\xda\x1b\xd7\xfc\x1f\x55\x1b\x02\x02\xa7\x3e\x38\xcc\xe9\xf7\x60\xf6\xb5\xb2\x76\x10\x89\xa9\x13\xa8\x7b\xa7\x93\xf1\x0e\x8f\x26\xb5\x63\x49\xe0\xe8\xfb\xa0\x19\x9d\x4a\x2d\x8d\xa3\x8a\x3c\x13\x28\x1c\x63\xf1\xd7\xb0\x07\x99\xf7\x51\x35\x7c\x89\x7c\x8e\x57\xd3\xc5\xbf\x1e\xe7\xda\x99\x84\x05\xd5\xe6\xfc\x15\x60\x30\xfa\x4f\x00\x00\x00\xff\xff\xc2\xfc\x39\x4f\xae\x02\x00\x00")

func scriptsSetDcgmExporterShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsSetDcgmExporterSh,
		"scripts/set-dcgm-exporter.sh",
	)
}

func scriptsSetDcgmExporterSh() (*asset, error) {
	bytes, err := scriptsSetDcgmExporterShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/set-dcgm-exporter.sh", size: 686, mode: os.FileMode(0644), modTime: time.Unix(1718793461, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xff, 0x76, 0xfc, 0x8c, 0xc5, 0xc9, 0xf9, 0x15, 0x9d, 0x23, 0x45, 0x58, 0x71, 0x99, 0xb1, 0xde, 0x56, 0x92, 0xbf, 0xf1, 0x97, 0x52, 0x93, 0xb1, 0x92, 0x1a, 0x2b, 0x73, 0x23, 0xa9, 0xda, 0x9b}}
	return a, nil
}

var _scriptsSetSystemctlServiceSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x52\xc1\x6e\x1a\x31\x10\xbd\xcf\x57\xbc\x2e\x2b\xd1\x1e\xc8\x96\x6b\xda\x8d\x44\xc1\x52\x91\x9a\xa4\x82\x10\x0e\x08\x21\x67\x99\x05\x2b\xc6\x46\xf6\x00\x41\x51\xfe\xbd\xda\x65\x9b\x03\xa1\x3d\x8d\x34\x7e\x7e\xf3\xe6\xcd\x6b\x7d\xca\x9e\x8c\xcb\x9e\x74\x5c\x13\x45\x16\x74\xf8\x54\x3c\xb6\x66\xcb\xa5\x36\xb6\xee\x8f\x8f\x51\x78\xd3\x17\x3b\xe6\xb0\x37\x05\x7f\xfe\x82\x57\x02\xac\x2f\xb4\x85\x7a\x54\xa3\xde\x70\x71\x77\x3f\x50\x8b\x9f\xf7\xb7\x2a\x4f\xd2\x6e\x42\x00\x17\x6b\x8f\xe4\xfc\xf5\x1a\xe9\x79\x2b\x21\x6a\x61\xc0\xa5\x71\x0c\x59\x33\xe2\x69\x08\x84\x37\x5b\xab\x85\x09\x18\xab\xd1\xe3\xb0\xaf\x16\x0f\xea\xf6\xf7\xaf\xde\x83\xca\xdb\xb3\x89\x33\x32\xa7\x01\xc7\x22\x98\xad\x18\xef\x72\xb5\xe7\xd0\x1b\x62\xc4\xd1\xef\x42\xc1\xb8\xf3\x4b\x46\xa5\x98\x03\xf5\x4a\xe1\x90\x3b\x96\x83\x0f\xcf\x57\xa2\xc3\x8a\x85\x68\xd6\xec\x33\xa7\xa9\x0f\xcf\xc6\xad\x06\x26\x70\x21\x3e\x1c\xf3\x76\x92\xbe\x9e\x0b\x7d\x4b\xda\xa4\x5e\xb8\x18\x8b\x0e\xf2\x2f\x44\xc6\x7b\x0e\xda\x74\x42\x23\xa3\xe3\xfc\x92\x69\xc4\xb1\xfe\xa4\xed\x41\x1f\x23\x4d\x22\x87\x3c\x78\x5f\x69\x18\xba\x28\xda\xda\x39\x4d\xb5\x13\x5e\xfe\x38\xe6\x9b\x9d\x15\xd3\xd9\x45\x0e\x8d\xd2\x36\x11\xd0\xc2\x34\x18\xb9\x6c\x11\xc4\x43\xa3\x34\x96\x6b\xe0\xc9\xfa\xf4\xdc\xb5\x04\x37\x48\x3e\xda\x9f\x55\x2c\x17\x65\x5f\x35\x73\xde\xaf\x79\x91\x32\x63\x29\xb2\x58\x47\x64\xd9\xd4\xff\xb3\xbd\x55\x07\xef\x6b\x6b\x9b\x5d\x3e\xe4\x0b\xe5\xce\x15\xd5\x4d\x71\x30\xb2\xae\x51\x7f\x99\xb0\xd5\xb2\x26\x53\x62\x86\x24\x6d\x25\xe8\x38\x46\x17\xf3\x6f\x15\xc8\xbd\xeb\x9c\x44\xbd\xe2\x6b\xa4\x5f\xf1\xfd\x7c\xdd\x9b\x3a\x9b\x2f\x46\xd0\xa5\xd2\x5c\x8c\x37\xaa\x04\xff\x09\x00\x00\xff\xff\x09\x57\x84\xfd\x1c\x03\x00\x00")

func scriptsSetSystemctlServiceShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsSetSystemctlServiceSh,
		"scripts/set-systemctl-service.sh",
	)
}

func scriptsSetSystemctlServiceSh() (*asset, error) {
	bytes, err := scriptsSetSystemctlServiceShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/set-systemctl-service.sh", size: 796, mode: os.FileMode(0644), modTime: time.Unix(1718790325, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xaa, 0xc2, 0xc, 0x4d, 0x9a, 0x76, 0x3a, 0xe5, 0x28, 0x53, 0xe2, 0xb0, 0xcd, 0x91, 0x64, 0xcf, 0xf7, 0xaf, 0x44, 0x7f, 0x8c, 0xe1, 0x2e, 0xb0, 0x32, 0x5c, 0x8e, 0x10, 0xca, 0x82, 0xb2, 0x14}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"scripts/check-step.sh":            scriptsCheckStepSh,
	"scripts/set-dcgm-exporter.sh":     scriptsSetDcgmExporterSh,
	"scripts/set-systemctl-service.sh": scriptsSetSystemctlServiceSh,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"scripts": {nil, map[string]*bintree{
		"check-step.sh":            {scriptsCheckStepSh, map[string]*bintree{}},
		"set-dcgm-exporter.sh":     {scriptsSetDcgmExporterSh, map[string]*bintree{}},
		"set-systemctl-service.sh": {scriptsSetSystemctlServiceSh, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
