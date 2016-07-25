// Code generated by go-bindata.
// sources:
// templates/init/django/.dockerignore
// templates/init/django/Dockerfile
// templates/init/django/docker-compose.yml
// templates/init/rails/.dockerignore
// templates/init/rails/Dockerfile
// templates/init/rails/docker-compose.yml
// templates/init/ruby/.dockerignore
// templates/init/ruby/Dockerfile
// templates/init/ruby/docker-compose.yml
// templates/init/sinatra/.dockerignore
// templates/init/sinatra/Dockerfile
// templates/init/sinatra/docker-compose.yml
// templates/init/unknown/.dockerignore
// templates/init/unknown/Dockerfile
// templates/init/unknown/docker-compose.yml
// templates/templates.go
// DO NOT EDIT!

package templates

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

var _initDjangoDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xce\xcf\x2b\xcb\xaf\xe0\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\x01\x13\x99\xe9\x79\xf9\x45\xa9\x5c\x80\x00\x00\x00\xff\xff\x57\x31\x5f\xce\x1d\x00\x00\x00")

func initDjangoDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerignore,
		"init/django/.dockerignore",
	)
}

func initDjangoDockerignore() (*asset, error) {
	bytes, err := initDjangoDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/.dockerignore", size: 29, mode: os.FileMode(420), modTime: time.Unix(1469130872, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initDjangoDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xcd\xae\xd3\x30\x10\x85\xf7\x79\x8a\x91\x60\xdb\x66\xd1\x27\x40\x25\x2c\x40\xb4\x51\x28\x48\x5d\x21\xe3\x4c\x52\x17\xc7\x63\xfc\x03\x8d\x50\xdf\x9d\xb1\xd3\xd0\xe6\xde\xbb\xb8\xbb\xcc\xf1\xcc\x99\x6f\x4e\x3e\x34\xfb\xcf\x20\xc9\xfc\xa6\x4b\xd9\x9e\x85\xe9\xa9\x28\xde\x80\x43\xab\x85\x44\xc0\x8b\x18\xac\x46\x61\x2d\x08\xd3\xce\xa5\x75\x74\x46\x19\x20\x10\x04\xa1\x34\x39\x08\x27\x04\x35\x88\x1e\x93\x36\x52\x74\x70\xeb\x61\xaf\xba\xd9\x7f\xac\xb6\x07\x50\x1e\x84\xf6\x04\xd1\x63\x0b\x3f\x46\xe8\xa3\x51\x92\x9c\x01\x65\xf2\xfc\x02\x02\xde\x93\xfc\x89\xae\x53\x1a\x8b\x6a\xf7\x0d\xde\xd5\xf5\x03\x4c\x96\x66\xdf\x25\x54\xa2\x17\x06\x70\xb0\x61\x84\x2f\xd5\xb6\xa9\x0e\xdf\x3f\x55\x47\x68\xa3\x53\xa6\x87\x41\x18\xa6\x5c\xdb\x91\xd7\x0d\x5c\xb4\x1e\xfe\x28\xad\xf9\x60\x1f\x75\x48\x28\x69\xd8\x39\x72\x79\xc7\x83\x41\x47\x39\x19\x49\x3c\x4b\x46\x8f\x99\x39\xf1\x79\x30\x88\x2d\xdf\xd4\x71\x10\x56\x59\x36\xf1\x41\x68\x5d\x6c\xf7\xf5\x91\x8d\x7f\x45\xe5\x70\x40\x13\xfc\x3a\x5c\x02\x94\xcc\x5f\x3e\x55\x8b\xe6\xeb\x2e\xcd\x6e\xe6\x61\x58\xad\xa2\xed\x9d\x68\x31\xc9\x2f\x3c\xbb\x67\xce\xaf\xa0\x93\xa4\x35\x67\xc4\x16\x41\xc9\x89\xef\xed\x5f\x8e\xf6\x5a\x4e\xd2\x04\xb7\x90\xe6\xae\x5b\xda\xd7\xb9\x65\xae\xa7\xf7\x7b\xac\xf9\xf9\x7f\x39\x81\x8f\xe1\x44\x66\xb3\xc8\xfe\x81\x83\x4f\x35\xa4\x8c\x8d\xf7\x0b\x12\x3c\xff\x91\x00\xd4\xe5\xef\xf4\xcb\xf3\x9e\x75\xf6\x2f\xfe\x05\x00\x00\xff\xff\x22\xcb\xe6\x65\xb5\x02\x00\x00")

func initDjangoDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerfile,
		"init/django/Dockerfile",
	)
}

func initDjangoDockerfile() (*asset, error) {
	bytes, err := initDjangoDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/Dockerfile", size: 693, mode: os.FileMode(420), modTime: time.Unix(1469130872, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initDjangoDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8d\xc1\xaa\xc2\x30\x10\x45\xf7\xfd\x8a\xfc\xc0\xcb\x8b\x36\x82\x04\xba\x92\xae\xdc\xa9\x1b\x57\x92\xb4\x43\x09\xa6\x99\x92\x4c\x6b\xfd\x7b\x13\x68\xbb\x10\xdc\xcd\xbd\xf7\x70\xe6\x05\x46\x15\x8c\x99\xd1\xba\x56\x31\x9e\x4e\xf0\x93\x0d\xe8\x7b\xf0\x94\x17\xc6\xfe\xd8\xb5\x3e\x5d\xea\xdb\xe3\x5c\xdf\x53\xe1\xb4\x01\x17\xd7\xa9\x41\x3f\xe1\xcc\x07\x0c\xc4\xa5\x2c\xf9\x10\x90\xb0\x41\x57\x91\x8b\xbf\x91\xf9\x5d\x51\x18\x21\xdb\xac\x7f\x6e\xb2\x56\x93\x36\x3a\xe6\x3e\xd3\x5b\x7f\x14\x4a\x0a\x21\x96\x94\x1c\x39\xee\x8a\x15\xcf\x98\xed\x75\x07\x6a\xf9\xf5\x3f\x60\xa4\x2e\x40\xfc\x16\x1d\x64\xb9\x2f\x3e\x01\x00\x00\xff\xff\x25\x21\x30\xfe\xf3\x00\x00\x00")

func initDjangoDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initDjangoDockerComposeYml,
		"init/django/docker-compose.yml",
	)
}

func initDjangoDockerComposeYml() (*asset, error) {
	bytes, err := initDjangoDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/django/docker-compose.yml", size: 243, mode: os.FileMode(420), modTime: time.Unix(1469130872, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x97\x7e\x4a\x92\xbe\x96\x5e\x71\x61\x4e\x66\x49\xaa\x31\x2a\x4f\x37\x2b\xbf\xb4\x28\x2f\x31\x87\x4b\x3f\x27\x3f\x5d\x5f\x8b\x4b\xbf\x24\xb7\x80\x0b\x10\x00\x00\xff\xff\xa0\x04\x95\x56\x4e\x00\x00\x00")

func initRailsDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerignore,
		"init/rails/.dockerignore",
	)
}

func initRailsDockerignore() (*asset, error) {
	bytes, err := initRailsDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/.dockerignore", size: 78, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x50\xbb\x4e\xc4\x30\x10\xec\xfd\x15\x2b\xd1\x5f\x7a\x5a\x24\xa8\xe0\x50\x24\x0a\xba\xf3\x39\xeb\x60\xe2\xd8\x96\x1f\x27\xf2\xf7\x6c\xd6\x09\x21\xa7\xa4\xca\x3c\xec\x99\xf1\x73\x7b\x7e\x05\xe5\xdd\xcd\xff\x34\x51\x1a\x9b\x84\x78\x20\x1c\x26\xf0\xce\x4e\x90\xbf\x10\xb4\xb1\x98\xc0\x21\x76\xd8\x81\xf6\x11\xae\xc5\x75\x16\xc1\xb8\x94\xa5\xb5\xe4\x2f\x4e\xf9\x71\x44\x97\xd9\x7f\x43\xd7\xf9\xd8\x28\xa9\x08\x58\xe3\xc8\xa9\x61\xf2\x05\x2e\xcb\xc1\x20\xd5\x20\x7b\xbc\xcc\x64\x84\x1e\xc7\x24\x9e\xce\xef\x9f\xf0\x82\xe3\x9c\x05\xfc\x35\x32\x84\x66\x61\x76\xf2\xc9\x7a\x35\xec\x64\x66\xa8\x06\xbb\x76\xe9\xec\xfa\xcf\x88\xf6\xe3\xed\xbe\xff\x3a\xf8\xbb\xa4\x7c\x3c\x58\xa6\x84\x39\x3d\x86\x88\xb4\x33\xfc\x15\x6a\xe5\x80\x4b\x61\x0e\x5a\x71\x55\xe9\x51\xb5\xe9\xb7\x2d\x15\x57\x2d\x94\xab\x35\x6a\xd3\x2a\xae\xda\x8c\x6b\x60\xd5\x36\xcc\xe5\x23\x85\x1c\x14\x5a\x47\xcc\xfd\x23\xd2\x10\xaf\xf9\x9f\x4e\xd7\x6b\x4f\x7c\x9b\xf8\x0d\x00\x00\xff\xff\x08\xae\x24\x4e\xf0\x01\x00\x00")

func initRailsDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerfile,
		"init/rails/Dockerfile",
	)
}

func initRailsDockerfile() (*asset, error) {
	bytes, err := initRailsDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/Dockerfile", size: 496, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRailsDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initRailsDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRailsDockerComposeYml,
		"init/rails/docker-compose.yml",
	)
}

func initRailsDockerComposeYml() (*asset, error) {
	bytes, err := initRailsDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/rails/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initRubyDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerignore,
		"init/ruby/.dockerignore",
	)
}

func initRubyDockerignore() (*asset, error) {
	bytes, err := initRubyDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xcd\x4e\xc0\x20\x10\x84\xef\x3c\xc5\x26\xde\xcb\x43\x98\xe8\x49\x6b\x9a\x78\xf0\x56\x0a\x4b\x6d\x0a\xbb\x84\x9f\x46\xde\x5e\x8a\x35\xb1\x72\x62\x67\xbf\x61\x86\xa7\x69\x7c\x01\xcd\x74\xf0\x97\x8c\x65\xa9\x42\x3c\xb4\x31\x54\x60\x72\x15\xf2\x27\x82\xdd\x1c\x26\x20\x44\x83\x06\x2c\x47\x58\x0a\x19\x87\xb0\x51\xca\xca\xb9\xc6\x17\xd2\xec\x3d\x52\xee\xfc\x81\x64\x38\x4a\xad\x74\x1b\xdc\x46\x8d\xb4\x50\xb9\xc0\x7c\x19\x83\xd2\xbb\x5a\x71\x3e\xc5\x08\x2b\xfa\x24\x1e\xc7\xb7\x0f\x78\x46\x7f\x66\x41\x3f\x52\x85\x20\x2f\xe5\xb6\x1e\x1c\xeb\xfd\xb6\xee\x4a\xab\xd1\xa9\x5b\x7a\xa7\xfe\x2a\x62\x7a\x7f\xfd\xdf\xff\xf7\xc3\x67\xf7\x88\x29\x03\xdb\x7e\x6f\xde\x9f\xe0\xa1\xbf\x23\xbe\x03\x00\x00\xff\xff\x8b\xae\xa0\xae\x2a\x01\x00\x00")

func initRubyDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerfile,
		"init/ruby/Dockerfile",
	)
}

func initRubyDockerfile() (*asset, error) {
	bytes, err := initRubyDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/Dockerfile", size: 298, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initRubyDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x8c\xc1\x0a\x83\x30\x0c\x86\xef\x3e\x45\xf0\x6e\xe9\x98\x87\x21\xf8\x30\x5a\x03\xca\xa2\x29\x69\x3a\xed\xdb\x2f\x85\x6d\xb7\xdd\xfa\xf5\xff\xf2\x9d\x38\x0f\x0d\xc0\x9c\x37\x5a\x06\x70\xf6\x0c\xbc\xef\xd3\x61\x80\x61\x65\x68\x71\xd9\x14\x16\x0e\x4f\x94\xce\xa6\xc8\x09\x5d\xd9\x09\xce\x4d\x57\x28\x9c\x05\x92\x4e\xa2\x39\x7e\x0f\x5b\x6b\xd0\x34\x23\xa5\x1a\x06\xe8\x6c\x38\x5e\x7c\xb9\xc8\xa2\xae\xef\xef\x2e\x0a\x2b\x07\xa6\x51\x29\xfd\x57\xae\x32\xaa\x64\x34\xa1\xfe\xfe\x62\x0f\x3f\xf4\xde\xfb\x0f\x99\x5b\xf1\xd6\xbc\x03\x00\x00\xff\xff\xae\x01\x4e\xf5\xc8\x00\x00\x00")

func initRubyDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initRubyDockerComposeYml,
		"init/ruby/docker-compose.yml",
	)
}

func initRubyDockerComposeYml() (*asset, error) {
	bytes, err := initRubyDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/ruby/docker-compose.yml", size: 200, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd7\x4b\x2a\xcd\x4b\xc9\x49\xe5\xd2\xd7\x4b\xce\xcf\x2b\xcb\xaf\x00\x32\x52\xf3\xca\x80\x64\x7a\x66\x09\x17\x20\x00\x00\xff\xff\xc9\x68\x92\x70\x1e\x00\x00\x00")

func initSinatraDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerignore,
		"init/sinatra/.dockerignore",
	)
}

func initSinatraDockerignore() (*asset, error) {
	bytes, err := initSinatraDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8f\xc1\x6e\x84\x20\x10\x86\xef\x3c\xc5\x24\xbd\xcb\x43\x34\x69\x4f\xad\x8d\x49\x0f\xbd\x49\x61\x70\x89\x30\x43\x00\xcd\xfa\xf6\x8b\xac\x9b\xac\xeb\x49\xbe\xf9\x86\xff\xe7\x63\xe8\xbf\x40\x33\xad\x7c\x95\xd9\x91\x2a\x49\x09\xf1\x56\x49\xdc\x80\xc9\x6f\x50\x2e\x08\xd6\x79\xcc\x40\x88\x06\x0d\x58\x4e\xf0\xbf\x90\xf1\x08\x8e\x72\x51\xde\x57\x7f\x21\xcd\x21\x20\x95\xe6\xaf\x48\x86\x93\xd4\x4a\xd7\x83\x77\x54\x4d\x0b\x1b\x2f\x30\x1e\x8b\x51\xe9\x59\x4d\x38\xee\x30\xc1\x84\x21\x8b\xf7\xfe\xe7\x0f\x3e\x31\xec\x59\xd0\x3e\xa9\x62\x94\x07\x39\x8d\x3b\xcf\x7a\x3e\x8d\x1b\xa9\x35\x9a\x75\x4a\x6f\xd6\x33\x11\xc3\xef\xf7\x6b\xff\xc7\x83\xf7\xee\x09\x73\x01\xb6\xed\xbf\xee\xde\x83\xbb\x76\x8f\xb8\x05\x00\x00\xff\xff\x9c\x51\x49\xbe\x2d\x01\x00\x00")

func initSinatraDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerfile,
		"init/sinatra/Dockerfile",
	)
}

func initSinatraDockerfile() (*asset, error) {
	bytes, err := initSinatraDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/Dockerfile", size: 301, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initSinatraDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x4f\x4d\xb2\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\x03\x32\x73\x12\x93\x52\x73\x8a\x41\x82\x0a\x0a\xba\x0a\xc9\xf9\x79\x65\xf9\x15\x7a\x05\xf9\x45\x25\x7a\x26\x26\xc6\x7a\x05\x45\xf9\x25\xf9\xc9\xf9\x39\xb6\x25\x39\xc5\xb8\x95\x54\x54\xda\x96\x14\x95\xa6\x02\x15\x80\x44\xe1\x86\x59\x18\x58\x99\x18\x18\x18\x40\x79\x40\xb5\x20\xae\x21\x17\x20\x00\x00\xff\xff\xc0\xe1\x22\xef\x84\x00\x00\x00")

func initSinatraDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initSinatraDockerComposeYml,
		"init/sinatra/docker-compose.yml",
	)
}

func initSinatraDockerComposeYml() (*asset, error) {
	bytes, err := initSinatraDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/sinatra/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1464072726, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerignore = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\x4b\xcd\x2b\xe3\xd2\x4b\xcf\x2c\xe1\x02\x04\x00\x00\xff\xff\x9c\x10\x28\x7b\x0a\x00\x00\x00")

func initUnknownDockerignoreBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerignore,
		"init/unknown/.dockerignore",
	)
}

func initUnknownDockerignore() (*asset, error) {
	bytes, err := initUnknownDockerignoreBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/.dockerignore", size: 10, mode: os.FileMode(420), modTime: time.Unix(1465343086, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerfile = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\x28\x4d\x2a\xcd\x2b\x29\xb5\x32\x34\xd3\x33\x30\xe1\xe2\x72\xf6\x0f\x88\x54\xd0\x53\xd0\x4f\x2c\x28\xe0\x02\x04\x00\x00\xff\xff\xf1\xa3\x65\xfc\x1f\x00\x00\x00")

func initUnknownDockerfileBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerfile,
		"init/unknown/Dockerfile",
	)
}

func initUnknownDockerfile() (*asset, error) {
	bytes, err := initUnknownDockerfileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/Dockerfile", size: 31, mode: os.FileMode(420), modTime: time.Unix(1465343086, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _initUnknownDockerComposeYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4d\xcc\xcc\xb3\xe2\x52\x50\x48\x2a\xcd\xcc\x49\xb1\x52\xd0\xe3\x02\x04\x00\x00\xff\xff\xa6\xe1\xc1\x85\x11\x00\x00\x00")

func initUnknownDockerComposeYmlBytes() ([]byte, error) {
	return bindataRead(
		_initUnknownDockerComposeYml,
		"init/unknown/docker-compose.yml",
	)
}

func initUnknownDockerComposeYml() (*asset, error) {
	bytes, err := initUnknownDockerComposeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init/unknown/docker-compose.yml", size: 17, mode: os.FileMode(420), modTime: time.Unix(1465343086, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9b\x5b\x6f\x1b\xc7\x92\xc7\x9f\xc5\x4f\x31\x47\xc0\x39\x90\x16\x3a\xd2\xdc\x2f\x02\xf2\x72\xec\x2c\x90\x87\xcd\x01\x4e\x92\x87\xc5\x7a\x11\xcc\xa5\x47\xcb\x8d\x44\x7a\x49\x2a\xb1\x62\xe4\xbb\x6f\xff\xaa\x6a\x44\x5a\x1c\x52\x32\x25\xc1\xde\x8b\x81\xb1\xc8\x9e\xee\xea\xaa\xee\xaa\x7f\x5d\xba\x79\x71\x11\xbc\x99\x77\x2e\xb8\x72\x33\xb7\xa8\x57\xae\x0b\x9a\xbb\xe0\x6a\xfe\xd7\x66\x3a\xeb\xea\x55\x7d\x3e\xf1\x1d\x96\xf3\xdb\x45\xeb\x96\x97\x7c\x5e\xb9\x9b\xf7\xd7\xbe\xdf\xf2\x62\x3a\x9b\xae\x2e\xba\xff\xac\x67\x57\xf3\x8b\xf3\x6e\xde\xfe\xe2\x16\xd3\xab\xd9\x7c\xe1\x76\x77\x7b\x2b\xbd\xfa\xe9\xf5\x9e\x3e\x4a\xe9\xaf\xed\xfc\xe6\xfd\x7c\xe9\xce\xef\x6e\xae\x47\xfa\x2e\xea\xe9\xf5\xf2\xd1\x59\xb5\xd7\xde\x49\xb5\xcb\xd3\xe6\xbc\x6d\xee\x1e\x9f\x92\x4e\xfb\x67\xa4\xc7\x93\x26\x5c\x4e\x67\xf5\x6a\x51\x3f\x3a\xe7\xd0\x6f\xef\xb4\x43\xa7\x27\xcd\x7c\x3b\xfb\x65\x36\xff\x6d\xf6\xe8\xcc\x43\xbf\xbd\x33\x0f\x9d\x1e\x9b\xf9\xfe\xd3\xf9\xd5\x9c\x37\x6f\xff\x1e\x7c\xff\xf7\x1f\x83\x6f\xdf\x7e\xf7\xe3\x9f\x26\x93\xf7\x75\xfb\x4b\x7d\xe5\xd6\xfd\x27\x93\xa9\x27\xb4\x58\x05\x27\x93\xa3\xe3\xe6\xce\xb7\x1c\xfb\x0f\x50\x5f\xb8\xe5\xf2\xe2\xea\xf7\xe9\x7b\x1a\xfa\x9b\x15\x7f\xa6\x73\xfd\xff\x62\x3a\xbf\x5d\x4d\xaf\xf9\x32\x97\x01\xef\xeb\xd5\x7f\x5c\xc0\x39\x1f\x68\x58\xae\x16\xd3\xd9\x95\xbc\x5b\x4d\x6f\xdc\xf1\xe4\x74\x32\xe9\x6f\x67\x6d\x60\x16\xf1\x0f\x57\x77\x27\x7c\x08\xfe\xed\xdf\x99\xf6\x2c\x98\xd5\x37\x2e\xd0\x61\xa7\xc1\xc9\xd0\xea\x16\x8b\xf9\xe2\x34\xf8\x38\x39\xba\xfa\x5d\xbe\x05\x97\xdf\x04\x70\x75\xfe\xbd\xfb\x0d\x22\x6e\x71\x22\x6c\xf3\xfd\x6f\xb7\x7d\xef\xbf\x43\xf6\xf4\x74\x72\x34\xed\x65\xc0\x9f\xbe\x09\x66\xd3\x6b\x48\x1c\x2d\xdc\xea\x76\x31\xe3\xeb\x59\xe0\x45\x3a\xff\x16\xea\xfd\xc9\x31\x84\x82\x3f\xff\xd7\x65\xf0\xe7\x5f\x8f\x95\x13\x99\xcb\xd3\xf8\x63\x32\x39\xfa\xb5\x5e\x04\xcd\x6d\x1f\xe8\x3c\x3a\xc9\xe4\xe8\x67\x65\xe7\x9b\x60\x3a\x3f\x7f\x33\x7f\x7f\x77\xf2\x17\xdf\xe7\xcc\xf3\xe6\x47\xb5\xd7\xdf\x0e\x9c\x9e\xbf\xb9\xf6\xfb\x74\xe2\xc5\x7f\x21\x7e\x20\xa3\xf4\x77\x10\xf2\x1d\x95\x6f\x6b\xf4\x6c\x9d\xff\x0d\xd6\x4f\x4e\xcf\xe8\x31\xf1\xef\x56\x77\xef\x5d\x50\x2f\x97\x6e\xc5\x92\xdf\xb6\x2b\xa8\x88\x7c\xb6\x1f\x7e\x9a\x59\x3f\x0f\x82\xf9\xf2\xfc\x9f\xfd\xb6\x7e\xe7\xbf\xdc\x8f\xb3\x2d\x1c\xda\x37\x28\xc8\x1e\xfa\x7f\xba\x8d\x93\xa3\xe5\xf4\x77\xf9\x3e\x9d\xad\xf2\x74\x72\x74\x03\x44\x06\xf7\x44\xff\xc5\x7f\x95\xc6\x1f\xbd\x86\x04\xa8\xc9\x39\x9f\x98\x47\x54\xe5\xa4\x9f\x3e\x9c\xeb\x34\xf8\xde\x4f\x71\x72\x6a\x33\x30\xa7\x49\xd9\x4f\xcf\x99\xdd\x0f\xde\x3d\xf6\x07\xcf\x8e\x1f\x2b\xdc\x7c\x3a\x14\x46\xf7\x0e\x85\x57\x3f\x74\x83\xf3\x4f\x09\x20\xda\x63\x04\x10\xce\xd3\xb8\x17\x74\x8b\x82\x49\xbf\x9b\xc8\x77\xcb\xb7\xd3\x85\x27\xd1\xcc\xe7\xd7\x9b\xa3\xeb\xeb\xe5\x23\x92\xdf\x2d\x55\x70\x8f\x2f\x75\xeb\x3e\xfe\xb1\x31\xda\x54\x02\x2d\xff\x19\xa8\x79\x2b\x1e\xe4\xed\x06\x66\x79\x25\x57\xad\x38\x39\x7e\xf7\x21\xea\xdf\x7d\x28\x9b\x77\x1f\xc2\xd2\x3f\xa1\x3d\xd5\xbb\x0f\xb9\xf3\xed\xd6\xd6\xfb\x3e\x5d\xfc\xee\x43\xea\xfb\xb5\xbe\xbd\xf5\xdf\x63\x3e\xfb\xa7\xf6\x9f\x5d\xb8\xf1\xbe\xd3\x77\x2e\xd9\x68\xa3\x7f\xeb\x69\x45\x7e\x3e\xdf\x5e\x79\xfa\xce\x3f\x85\x7f\x7a\xff\xa4\x99\xa7\xe3\xff\x66\xbe\x4f\x19\x6e\xf0\x61\x73\xf3\x64\xc5\xbb\x0f\x89\x1f\x9f\xf5\xca\x43\xd4\x6d\xf6\x3b\x1e\xf0\x68\x5c\x62\xb3\x97\x31\x1c\x1a\xac\x6a\x03\xc7\xbc\x01\xee\x58\xb9\x33\xff\xea\x78\xa7\x8b\x3f\xf6\xaf\x4f\xef\xd5\x7d\x9c\x02\x4c\xfc\x93\x58\xea\x26\x13\x62\xaa\xf7\x78\xb8\x57\x86\xc7\x70\xe7\x1e\x2e\xc4\xe0\x3d\xb5\x07\xca\xf3\x11\xb3\xba\x0c\xf6\x49\x11\x60\x3e\x97\x41\x5c\x9d\x05\xd8\xc1\xe5\xa6\x99\x9c\xa4\x71\x78\x2a\xed\x68\xf7\xa5\x6a\xff\x4f\xb3\xe9\x87\x93\x28\xcd\xab\x28\x09\xcb\x22\x3e\x0b\xc2\x53\x0f\x6c\x35\xb3\xff\x45\x64\xfd\x28\x02\x5e\x06\x26\x27\xac\x5d\xca\xff\x7f\xdc\x6f\x40\x7d\xb6\x57\x73\x71\x46\x07\xe9\x6d\xe9\x75\xaa\x8a\x54\x2f\x6b\xff\xae\xf3\xfa\x97\xf8\x77\x91\x7f\x4a\xaf\x77\x7d\xa1\x7a\x58\xd6\xda\x2f\x47\x97\x3d\xdd\x3c\xf7\x7f\xfd\xf7\xd8\xbf\x4f\x7d\x5b\x9c\xa9\x0e\xf3\xb9\x49\xbd\x1e\xf2\xce\xcf\x93\xfa\x27\x43\xe7\x23\xd5\xf9\xd4\xf7\xc9\xbc\xde\x47\x7e\x5c\xeb\x9f\xdc\xb7\xf5\xe8\xbe\xff\x5b\xfa\x7e\x19\xf4\x3d\x5f\x95\xff\xdc\x44\xca\x4f\xe7\xdb\x1c\xf3\x79\xfe\x1a\x3f\x77\x53\xea\xdf\xd6\x8f\xeb\x23\xfd\x8b\xcd\xe4\x7e\x5c\xea\xfb\x24\x3c\x9e\x87\x7e\xb0\x2d\x3f\xbe\xad\x74\x9e\x3a\x57\x9b\xeb\xfc\xf7\xca\xa9\x8c\xd8\x1a\xf6\x05\xbf\xcc\x81\x8d\xa5\x7e\xde\xba\xd1\xf7\xa9\xa7\xd5\x86\xba\x9e\x91\xef\x53\x23\xa7\xa7\x93\x23\x63\xa7\x6b\x0c\x9f\xd8\x5d\xed\xfb\x17\x3c\xa9\xf6\x89\x2a\x9d\x9f\xf5\x0c\x7d\x5b\x1d\x29\x6f\x49\xa5\xe3\x58\x3f\xda\x93\x4c\xf7\x25\xf2\x34\x2a\xf6\x20\xd7\x75\x82\x4e\x81\xfc\x8d\xce\x07\x9e\x34\xb5\xf2\x5f\xf4\xca\x4b\xe3\xfb\x86\x85\xae\x1d\xe3\x4b\x64\xcf\x95\x2e\x7b\xc4\x1a\x87\x7e\x7c\xd2\x2b\x4f\x0e\x19\x12\xdd\xa3\xca\xcf\x51\x19\xf6\xe4\xec\x77\x6c\xfb\x11\xeb\xd3\x19\x3f\xb4\x95\x95\xea\x48\xe6\xbf\x47\xb5\xae\x47\x5e\xab\x8e\x84\x9d\xf6\xed\xa0\x91\xe9\x7e\xb2\xd7\x55\x6e\xba\xd2\xab\x8e\x64\xac\x81\xed\x7f\x98\xab\x6c\x4d\xa8\xb2\xc1\x77\xdc\x2b\x0d\x64\x62\x4f\x42\xa7\x63\xe1\x3d\x63\x2f\xd0\x99\x81\xff\x44\xf7\xb3\x44\x07\x23\xdb\x9b\x5c\x71\x12\x1d\x45\x5f\x3b\xe3\x8d\x36\xf4\x92\xf5\xe9\x9d\xee\x75\xdd\x29\xbe\xa2\xd3\xd8\x0b\xfb\x86\xbe\xf2\x2e\xf7\xed\x5d\xa9\xfb\x54\xc4\x6a\x03\xe8\x6b\x99\xe8\x5c\xf0\xc1\x3b\xf6\x37\xf5\x4f\xd2\xaa\x5e\xb1\xbe\x65\xaf\xfa\xc8\x7b\xf4\x93\xb1\xd8\x14\xfb\x8b\xbe\x20\x4f\xc7\xbe\x46\xaa\x17\x19\x3c\x57\xba\xe7\xf4\x87\x7e\x6e\x76\x93\xb7\xba\xbe\xac\x29\xf2\x60\x23\xec\x3b\x3e\xc1\x65\xba\x7e\xd8\x5c\x64\x7b\x94\xd4\x2a\x2b\x7b\x57\xa5\x6a\x1b\xf8\x04\x6c\x82\xf5\x63\xcf\xb0\x25\xf4\x29\x76\x6a\xf7\x60\x82\x33\x7d\xce\x6c\x5d\xd8\x23\xd7\xa9\x1d\xc2\x0b\xbe\x05\x1b\x62\x7f\x90\x15\xfb\xcb\x0b\xd3\x79\xf4\x30\x54\x3d\xa9\x4d\x97\xe5\x1d\xeb\x9d\xab\x3c\x8c\x45\x7f\x5c\xaf\x74\xe1\xa9\x74\xaa\xa7\x59\xad\x76\x8b\x3f\x44\x67\x1b\x3f\xb6\x32\x9b\x17\x7d\xc3\x5e\x6b\xdd\xcb\xa6\x52\x3d\xa5\xbd\x2e\x0c\x9f\x1a\xb5\x83\xde\xfc\x25\xeb\xc3\xda\x97\x91\xee\x85\x8b\xd4\x86\xd1\xc3\x06\x3b\x2d\x55\x07\xe4\x3d\xfb\xd9\x2b\xcf\xf0\x0e\x1e\xb2\xc6\xa2\xd3\xd8\x7b\xac\xf2\x82\x95\xac\x3f\xb8\xc9\xde\xb1\xf6\xc8\xd2\xa7\xea\xe7\xfb\x44\xf1\x04\x1d\x42\x26\xd6\x89\x39\xc2\x6c\xdb\x57\xc7\xb1\x8e\x91\x35\x47\xd7\x33\xb3\xb7\xfd\xbe\x1a\x8c\x7f\xbe\xa7\x86\xca\x96\x9f\x5e\xbf\xda\xef\xa4\xe9\x71\x88\x8b\xde\x60\xfd\x35\x1c\xf4\x26\xfb\xe6\x9d\xf3\x2a\xf9\x8a\xdc\xf3\x1b\xcd\x5f\xff\xf5\xe6\xfa\x20\x27\x8d\x13\x40\x29\x5b\x1c\x80\x37\x84\x36\x5e\x3b\xe9\xd4\x9c\x74\xdf\xa9\x93\x06\x04\x70\x56\x28\x18\xb4\x01\x95\x72\x30\xac\x5a\x01\x5f\x1c\x7d\xab\x00\x1b\x35\x1a\x3c\xd2\x0e\x40\xe2\xf8\xe0\x01\x20\x05\xc4\x68\x07\xc8\xf3\x46\xe7\xc0\xd8\x00\x9b\xdc\x9c\x30\x3c\x40\x0b\x20\x69\xcc\x70\x0a\x33\x5e\x94\x5f\x9c\x60\x66\x81\x46\xa5\x0e\x09\x3e\x00\x3d\x40\x05\xa3\xc1\xf8\x7b\x03\x10\xc0\x1a\x07\xc5\x3c\xd2\x96\x5a\xb0\x90\xab\x41\x01\xca\x18\x8c\x00\x1a\x7d\x6b\x05\x7b\x82\x0b\x01\xfe\x5e\x1d\x03\x46\x8f\x3c\x12\x54\x17\x0a\x1e\xc8\x0b\x18\x25\x06\x0a\x80\x23\x8e\x33\x6c\x15\xac\x6a\x0b\x62\x00\x11\xe4\xaa\xcc\x39\x31\x46\xd6\x28\xd2\x35\x6d\x0c\x0c\xe8\x07\x0f\x99\x39\x1f\x82\x9c\xce\xc0\x08\x10\x62\x1f\x9b\x58\x65\x1d\x9c\x3a\x00\xcd\xda\x24\x16\x6c\x01\x6e\xf4\x8d\x58\x7b\xff\x2e\xac\x95\x0e\x8e\x53\x64\x6f\x15\xcc\x9c\xd3\xfd\x65\x2d\x09\x6a\xaa\x52\x81\x14\xc0\x41\x06\x71\xc4\x95\x3a\x0a\xe4\xc3\xa9\x01\x68\x80\x3c\x0e\x02\xbd\x00\x8c\x49\x0c\xf2\x54\x81\x34\x36\xc7\x10\x46\x23\x20\x95\x29\x1f\xe8\x19\xeb\x0e\xc0\x3d\x21\xa1\x58\x6b\xfa\xf3\xa1\x6a\x4d\x6b\x0b\xb0\xb6\xeb\x42\xfb\x81\x6b\x4d\xea\x10\xf8\xda\x12\xea\x35\x40\x6c\x4c\xa4\x21\xd5\x48\xbf\x34\x98\xfd\x83\x92\xe7\x8b\x24\xc9\xd8\x21\x81\x52\x5c\x2b\x66\x48\x02\x6c\x41\x14\x4e\x75\xb3\xcf\x58\x32\x0d\xad\x24\xd6\xe0\x18\x9d\x6c\x6b\x0d\xd8\xd1\xe9\xa2\x56\x27\xcf\xdc\x15\x38\xe4\xd4\xa6\x04\xdb\x9c\xda\x54\xe6\x34\xf0\x22\xd8\x01\x77\xe8\xcf\xdc\xe0\x2a\x78\x04\x5f\x12\x10\x14\x3a\x2f\xb6\x0e\x8e\x10\x14\x8a\xbd\x44\x16\x70\x5a\x40\x4d\x00\x2f\xc1\xe9\x10\xd4\x34\xfa\x8e\x71\xb1\x05\x33\x92\xb4\x1b\x56\x3e\xb4\xb3\xda\x12\x83\x2a\x53\x9c\x80\xa7\x1d\x76\xb6\xb5\x09\x87\x99\xd8\x16\x99\xb5\x75\x8d\x94\xc8\xb7\xed\x6a\x6b\xfc\x53\x4d\x6a\x17\xff\x2f\x6a\x4d\xa3\x22\x98\x1d\x15\xe5\xe7\x9a\x51\x1a\x16\x71\x11\xe7\xaf\x60\x46\x07\x67\xec\x64\x04\x64\x7a\x38\x5f\x94\x05\xc7\x35\x04\x03\xae\x55\x27\x8d\x93\x40\x79\x3b\xab\x08\x61\x18\x44\xd9\x28\x64\x5d\x6a\x45\x0a\x1a\x7c\xc7\xe1\xe0\x90\x31\x26\x71\x34\x8d\x3a\x0f\x17\xab\xb3\xc5\x68\x22\x53\x74\xc9\x0e\x0b\xe5\xa1\x33\x43\xc3\x49\x10\x89\xe3\x6c\x30\x46\xa2\x65\x67\xce\x5e\xb2\x9f\x44\x9d\x79\x61\xd1\xb0\x18\x79\xa7\xc6\x8f\x3c\x89\x55\x0b\xc8\x28\xab\x46\xb3\xc0\xa2\xb4\xac\xd9\xd3\x89\x73\x0d\x04\x5a\x0b\x14\xc4\x29\xf7\x2a\x2f\x86\x49\x95\x00\x67\x48\xd0\x03\x2f\x45\xae\x7c\x13\xdd\x13\x89\x47\x91\x06\x3c\xad\x65\xe6\x38\x41\x02\x29\xaa\x08\x95\x39\x7d\x02\x0a\x67\xd9\xa7\x04\x23\x95\x82\x04\x59\x03\x40\x32\x64\xdb\x04\x47\x80\x53\x6f\xd9\xa1\x64\x91\x96\xc9\x93\x39\x39\xcb\x2e\x68\x03\x60\x58\x2f\x64\x96\x4c\xab\xd1\x79\x91\x8b\x0c\x97\xbf\xac\x4b\xd3\x6a\x20\x02\x80\x95\x96\xa1\x92\x69\xb2\x37\x04\x47\xc8\x4d\xa6\x51\xf5\xba\x0e\x04\x73\xcc\x23\x59\x7b\x66\x59\x5f\xa6\xa0\x58\x5a\x96\x93\xdb\x3a\x30\x3f\xe0\xca\xfe\xd3\xc7\xd9\x3c\xf4\x41\x0f\xc8\xc0\xd1\x23\xfa\x52\x05\x11\xfd\xa9\x15\x74\xd1\x39\xd6\x8f\xb9\x08\x34\x00\x50\xd9\x2f\x74\xc5\x82\x3f\xf6\xbc\x68\x75\xcf\x33\xcb\xbe\x4b\xab\x6e\x24\x04\xa4\xb9\xd2\x61\x9f\x12\xcb\x50\xd1\x4d\x82\x15\xf4\x40\x2a\x2d\x56\xd9\x20\xd3\x47\x47\x59\xf7\xcc\xaa\x32\xe8\x45\x63\x9f\x01\xd2\xd8\x74\x3b\xb3\xcf\xb2\x7f\x8d\x65\x9c\xb5\x06\x54\x00\xb8\x04\xb8\xb5\x06\x9d\xb9\xe9\x33\x6b\x4e\x00\xc6\x5a\x33\x77\x9c\x68\x75\x07\x1d\x23\xf8\x92\x2c\xb4\xb0\xaa\x52\xa7\xef\x71\x46\xe8\x19\x8e\x87\x3d\xaf\x4c\x3f\xd0\x01\xe8\xe2\x28\x90\x1f\xbd\x65\x4d\xc2\x6e\x1b\xe0\xd1\x0d\xf8\x61\x3f\x87\x40\x77\x1d\x70\xed\x02\xf8\xc3\x93\xbd\x07\x44\x1e\x82\xfb\xbe\x54\xef\xc1\xd0\x03\x70\xfd\xb5\x12\xbd\x6d\xde\x0d\xd2\xd3\x2a\xff\x7a\x30\xfd\x99\x69\xde\x10\x7a\xa4\x9d\xa6\x0d\xce\xc2\x1c\x90\x5a\xea\x52\x16\x32\xa1\xa1\x84\x2c\x20\x1f\xef\xa9\x8f\x52\x6f\x03\x01\x40\x5b\x6a\x16\xa4\x1b\x99\x7d\x27\x6d\x20\x85\x91\x34\xb0\x5e\x23\x3e\x7f\xdb\xa1\x66\x62\x75\x15\x49\x7d\x32\x45\x8c\x30\x5b\x9f\x3f\xc4\xd6\x06\x1a\xf3\x90\x8e\x0d\x7d\x52\xeb\x17\xdb\xdf\x81\x26\x28\x40\x1d\x8f\x76\x3e\x83\xc0\x82\xba\x56\x07\xe4\xc1\x4a\xc5\xc3\x58\x08\x04\x82\x50\x0b\x89\x06\x74\x48\x35\x55\x2b\xad\x46\x16\x19\xd2\x56\xf6\x79\x78\x48\x71\x90\x41\xea\xcf\x86\xae\x62\x75\x96\xfe\xc5\x23\xa1\x17\x1e\x03\xda\xd4\x63\x40\x58\x90\xeb\xf1\xd0\xeb\xb9\x19\xce\x28\xa9\x87\x56\xfa\x94\xfc\x66\x94\xd0\x01\x36\xfb\xba\xd9\xcd\x6e\x79\xcc\x82\xa3\x24\xfe\xd2\x16\x7c\xdb\xdc\xfd\x8f\xca\x6d\x76\x2a\x74\xa5\x75\x19\xf2\x9d\xc2\x0e\x03\x76\x29\xf4\x03\x99\x0f\xd4\xe5\x07\x54\x36\xd4\x78\xeb\xe2\xcb\x88\x02\x3f\x18\xfd\x64\xdd\x1d\xe7\xfd\x65\xd5\x76\x84\x7f\x53\xd8\x24\xfc\x6a\xf4\xf5\xe0\x24\x42\x8e\x92\x7b\xd3\x52\x67\xc1\xf3\x70\xec\x97\x2a\x1c\x12\xd4\x01\xd9\xb1\x1d\xbd\xa1\xa5\x04\xc0\x52\x79\x2a\x55\xab\x09\x82\xaa\x5a\x03\xf6\xde\x8e\x32\x70\x2b\x12\xd4\x75\xe6\x6a\x2c\xa8\xaf\xec\xf8\x09\x5a\x99\x55\xb3\x08\xdc\x70\x63\x1c\xb1\xe4\xb1\x1e\x1b\x10\xdc\xe7\x06\xf9\x24\x16\x1c\xc5\x14\x76\x44\x05\xbf\x54\x43\xa5\x32\x58\x58\x80\x67\xc7\xe3\xa9\x95\xea\x25\xd8\x8c\xd4\xbd\x10\xf8\x15\xe6\x4e\x08\x86\x49\x62\x70\x83\x24\x20\x43\x72\x11\x9b\x9b\xe1\x28\x80\x80\x58\x8e\x31\x0b\x75\x33\xf0\xcf\x11\x0f\x19\x3f\x95\x40\x02\x52\xac\x92\x87\x0a\x1a\x6e\x50\xaa\x8b\xb1\x06\xcf\x04\xbd\x58\x2f\xd5\x3a\x39\x32\x8b\xd4\x7a\x93\x52\x2b\x11\xb5\xf1\x4f\xd5\xb5\xb3\x63\x44\x2a\xbb\x72\x64\x1a\x1a\xcd\x50\xe5\x95\x79\x2d\x08\x65\x0d\x87\xe0\x38\xb1\x64\x46\x8e\x0f\x1b\x0d\xd8\xe3\x8d\xc4\x81\x40\x39\xec\x75\x4f\xa0\x2f\x95\xc8\x7c\x5d\xf9\x84\xbf\xd2\x5c\x24\x09\x8f\xb3\xe3\x36\xc6\xb3\x26\xce\x2a\xbd\xb4\x39\xab\xa0\xd4\x96\x4c\xca\x35\x84\x46\xd1\x48\x8e\x5f\x9c\xa2\x1d\x7b\x47\x1b\x73\xf5\x76\xc4\x25\xa8\x44\x40\x9d\xe8\xbe\xf2\x59\x8e\xc0\x2a\x0d\x57\x3a\x4b\x0a\x39\xb6\x95\x63\x46\xbb\x2a\xc1\x11\xa9\x54\x4f\x12\x4d\x3a\xc2\x64\x1b\xe9\xd0\x71\x39\x02\x1a\xdc\x7c\xbd\x3b\xa8\xfe\xc4\x5a\x9e\x8b\x73\x0f\x42\xea\x4f\x6f\xee\xed\x83\xb8\xcf\x0a\xa8\xc7\x58\x7e\x79\x78\x1b\x09\xa7\xe3\xea\x8b\x97\x48\xee\x65\x7f\x89\x43\x93\x56\x73\x6e\xcc\xb8\xb4\x9b\x0d\x14\xfb\x81\x16\x67\xc5\x6f\x89\x5e\x43\x3b\x05\xad\x14\xde\x30\x77\x22\x47\x72\x3a\xc6\x60\x7a\xa1\x39\x64\x4e\x46\x51\x61\x60\x49\x8a\xf9\x9d\xaa\x32\xa6\x84\x59\x01\x7b\x98\x20\x10\x83\x69\x52\x90\x17\xb5\x8d\xf5\x34\x15\x28\xc0\x3c\xc9\x15\x31\x75\xf2\x78\x8a\x8d\x32\x87\x9d\x88\xcb\xe9\x7d\xab\xb5\x01\xf8\x24\xaf\x15\x48\xb4\x93\x72\x20\x0f\x47\x0f\x1c\x50\x9b\x88\xec\xa4\x95\x1c\x9a\x0c\x42\x4e\x7e\x9d\x46\xd3\xd4\x23\x42\x3b\xa8\xa0\x2f\x9f\xa5\x96\xd2\x59\x1e\x5d\xae\x6b\x09\x72\xd8\xe3\x54\x46\x39\xf9\x76\xca\x2b\x26\x0f\xd4\x93\x69\x50\xaf\xc1\x34\xa9\x9b\x44\xc6\x33\xee\x80\x9a\x10\xb2\x49\x81\xb5\x35\x28\x2b\x95\x56\x6d\x35\x0c\x1e\x6a\x21\x72\xc8\xd1\xe8\xc1\x85\x9c\xd8\x47\xba\xa6\xbd\xf1\x44\x7f\x02\x21\x8a\xb2\x02\x9f\x76\x73\xa2\xb7\xd3\x7b\xe0\x26\xb4\x5b\x01\x9c\xc0\x02\x1f\xdc\xf6\x08\xed\xe4\x1d\x7e\xa9\x3d\x51\xa3\x6a\xda\x71\x08\xa9\xed\xe0\x43\x72\xf2\xcc\xd6\xe9\xb1\x60\xe9\xd9\xc1\xff\x08\xa5\x07\x70\xf2\xa4\xd0\x7f\x84\xcc\xe7\x83\xcb\x2b\x07\xfe\xbb\x84\x19\xa0\x26\xfc\xd2\x71\xd4\x0f\x7a\xbd\xf9\xff\x5a\xe8\x3f\x22\xf6\x61\xca\x3c\x42\x68\xad\xcb\xa3\x17\xd1\xb7\x35\x79\x84\xc6\x53\x15\x79\xb7\x1c\x2f\xaa\xc7\x3b\x04\xf9\x5a\x92\x81\x4f\x56\xe1\xf9\xf9\x40\x64\xef\xd2\x8d\x7c\x20\x7f\x90\x0f\xa4\x7a\xd2\x3f\xe4\x03\x80\x33\xce\x10\xc7\x43\x7c\x4b\x3c\x8b\xda\x03\xc6\x8d\x7d\x96\x82\x74\xa8\x57\x80\x70\xa8\xa9\x01\x72\x3b\x38\xc9\xdc\x62\x3a\x2b\xa6\xd6\x16\x87\x4a\x99\xcb\x0a\xf2\xf0\x01\xaf\xce\xae\x0d\xe2\x74\x88\x13\x29\xf9\x50\x5c\x96\xd8\xba\xd5\x32\x92\x5c\xef\x1b\x4e\xf3\xcc\xe9\xe1\x10\x32\x3b\xed\xeb\xec\x2a\x6d\x6b\xd7\x06\xe5\x30\xc4\xe9\x3a\xe1\x34\xe4\xba\x97\x5d\x67\x2c\xad\x50\x2e\xd7\xc8\xca\x75\x11\x57\x6e\x01\x84\x4a\x97\x71\x5c\x67\x22\xd7\xc0\x99\x49\x8e\x53\x68\x3c\x1c\x5b\xde\x12\x5b\xf1\x1f\xfe\xe4\x6a\x53\xaf\x65\xbd\xd2\x6e\x07\x74\x76\x03\xa0\xb3\x6b\x66\x45\xa1\x39\x0c\x8e\x5a\x4e\x0f\x2d\xc8\x48\xcc\xbc\x59\x2f\xe6\x92\x42\x79\xa5\xfd\x70\xc0\xc4\xec\xf4\x69\x87\x2b\x75\x76\x6b\xa1\xb7\x78\xdd\x59\x7e\x20\xd7\x24\x23\xcd\x5b\x18\x2f\xce\xdd\xd9\xe1\x85\x5d\x1f\x0b\xed\x9a\x1b\x7b\x9a\xd8\x75\x4a\x39\x4c\x70\xda\x06\x4f\xf4\x27\x8f\x93\x3c\x30\x53\x19\x24\xa6\xcf\x75\x5f\x38\x80\x21\x07\x70\x96\x0f\xc8\x4d\x92\x7e\x7d\x9d\x8e\x36\x78\x66\x8f\x28\x57\xe2\x94\x25\xf7\xe8\x74\x8d\x79\x27\xd7\x32\x93\xf5\x61\x05\xba\x4a\x29\x71\xec\x7a\x15\xc1\x0c\xeb\x32\xe8\x8e\x5c\xd1\x1c\xcf\x0d\xb6\x8c\xe7\x05\x80\xf0\xd3\x0c\x61\xfb\x77\x36\x8f\x60\xe0\xe7\xe4\x09\xbb\xd8\x7f\x15\xfc\x1b\xc9\x16\x92\x30\xfa\x9a\xd0\xef\xff\xcb\xef\xff\x5b\xcb\xef\x3b\xb6\xf9\x05\xac\x75\x2c\x0c\xdf\xfd\xb3\xb7\x47\x6c\xf7\xf3\x83\xf1\xfd\x82\xbd\x8a\x1d\x7f\xd5\xa5\xf8\x9f\xf4\x77\x7f\x2f\xf7\x73\x9c\x3d\x3f\xb7\x41\xf7\x42\xbb\x19\x39\xe6\x43\x88\x7d\x62\xbb\x89\x28\x76\x3b\xae\x9b\x23\x2c\x1f\xa6\x97\x23\x84\xd6\x3a\x39\xfa\xeb\xca\x6d\x75\x1c\xa1\xf1\x54\x55\xdc\x2d\xc7\x8b\xaa\xe1\x0e\x41\x06\x0d\xfc\xec\x70\x3a\x4b\xd2\x24\x2c\x5f\x47\x01\x0f\x0e\xa7\x09\x03\x09\xd9\x86\xfb\x32\xdc\x67\x90\x5f\xc3\x74\x6b\x37\x82\x5a\x52\xf3\x68\xec\xee\x06\x61\xb4\xfc\xfa\xc6\xea\x54\xa8\x27\x2e\x08\x5a\x52\x0f\xb2\x3b\x19\x02\xdb\xa1\x5e\x72\xe5\x6f\x6a\xea\x1c\xdb\xbd\x9e\x5d\x2a\x4d\x98\xca\xaf\x26\x72\xbb\x33\x12\xf5\x4f\x53\xe9\xc3\xc3\xa2\x2d\x32\xdb\xea\xbc\x2f\x2c\xda\x1a\x7e\x90\x26\xbf\x56\x58\x34\x26\xc1\x10\x16\x7d\x76\x54\xf4\x9a\x4a\xfc\xcc\xa8\x88\x42\x44\x6a\x51\x0f\x0f\x3f\xa0\xf8\xef\x00\x00\x00\xff\xff\x1e\x9a\x69\x4f\x00\x40\x00\x00")

func templatesGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesGo,
		"templates.go",
	)
}

func templatesGo() (*asset, error) {
	bytes, err := templatesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.go", size: 32768, mode: os.FileMode(420), modTime: time.Unix(1469489334, 0)}
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
	"init/django/.dockerignore": initDjangoDockerignore,
	"init/django/Dockerfile": initDjangoDockerfile,
	"init/django/docker-compose.yml": initDjangoDockerComposeYml,
	"init/rails/.dockerignore": initRailsDockerignore,
	"init/rails/Dockerfile": initRailsDockerfile,
	"init/rails/docker-compose.yml": initRailsDockerComposeYml,
	"init/ruby/.dockerignore": initRubyDockerignore,
	"init/ruby/Dockerfile": initRubyDockerfile,
	"init/ruby/docker-compose.yml": initRubyDockerComposeYml,
	"init/sinatra/.dockerignore": initSinatraDockerignore,
	"init/sinatra/Dockerfile": initSinatraDockerfile,
	"init/sinatra/docker-compose.yml": initSinatraDockerComposeYml,
	"init/unknown/.dockerignore": initUnknownDockerignore,
	"init/unknown/Dockerfile": initUnknownDockerfile,
	"init/unknown/docker-compose.yml": initUnknownDockerComposeYml,
	"templates.go": templatesGo,
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
	"init": &bintree{nil, map[string]*bintree{
		"django": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initDjangoDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initDjangoDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initDjangoDockerComposeYml, map[string]*bintree{}},
		}},
		"rails": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRailsDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRailsDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRailsDockerComposeYml, map[string]*bintree{}},
		}},
		"ruby": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initRubyDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initRubyDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initRubyDockerComposeYml, map[string]*bintree{}},
		}},
		"sinatra": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initSinatraDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initSinatraDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initSinatraDockerComposeYml, map[string]*bintree{}},
		}},
		"unknown": &bintree{nil, map[string]*bintree{
			".dockerignore": &bintree{initUnknownDockerignore, map[string]*bintree{}},
			"Dockerfile": &bintree{initUnknownDockerfile, map[string]*bintree{}},
			"docker-compose.yml": &bintree{initUnknownDockerComposeYml, map[string]*bintree{}},
		}},
	}},
	"templates.go": &bintree{templatesGo, map[string]*bintree{}},
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

