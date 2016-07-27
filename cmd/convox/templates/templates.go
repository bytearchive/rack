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

	info := bindataFileInfo{name: "init/django/.dockerignore", size: 29, mode: os.FileMode(420), modTime: time.Unix(1469642945, 0)}
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

	info := bindataFileInfo{name: "init/django/Dockerfile", size: 693, mode: os.FileMode(420), modTime: time.Unix(1469642945, 0)}
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

	info := bindataFileInfo{name: "init/django/docker-compose.yml", size: 243, mode: os.FileMode(420), modTime: time.Unix(1469642945, 0)}
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

	info := bindataFileInfo{name: "init/rails/.dockerignore", size: 78, mode: os.FileMode(420), modTime: time.Unix(1465598672, 0)}
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

	info := bindataFileInfo{name: "init/rails/Dockerfile", size: 496, mode: os.FileMode(420), modTime: time.Unix(1465162563, 0)}
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

	info := bindataFileInfo{name: "init/rails/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1465162563, 0)}
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

	info := bindataFileInfo{name: "init/ruby/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1465598672, 0)}
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

	info := bindataFileInfo{name: "init/ruby/Dockerfile", size: 298, mode: os.FileMode(420), modTime: time.Unix(1465162563, 0)}
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

	info := bindataFileInfo{name: "init/ruby/docker-compose.yml", size: 200, mode: os.FileMode(420), modTime: time.Unix(1465162563, 0)}
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

	info := bindataFileInfo{name: "init/sinatra/.dockerignore", size: 30, mode: os.FileMode(420), modTime: time.Unix(1465598672, 0)}
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

	info := bindataFileInfo{name: "init/sinatra/Dockerfile", size: 301, mode: os.FileMode(420), modTime: time.Unix(1465162563, 0)}
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

	info := bindataFileInfo{name: "init/sinatra/docker-compose.yml", size: 132, mode: os.FileMode(420), modTime: time.Unix(1465162563, 0)}
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

	info := bindataFileInfo{name: "init/unknown/.dockerignore", size: 10, mode: os.FileMode(420), modTime: time.Unix(1465598780, 0)}
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

	info := bindataFileInfo{name: "init/unknown/Dockerfile", size: 31, mode: os.FileMode(420), modTime: time.Unix(1465598780, 0)}
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

	info := bindataFileInfo{name: "init/unknown/docker-compose.yml", size: 17, mode: os.FileMode(420), modTime: time.Unix(1465598780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9b\x59\x6f\x23\xc7\x11\xc7\x9f\xc5\x4f\x31\x16\x60\x43\x0a\x64\x69\x6e\xce\x08\xf0\x8b\x77\x1d\xc0\x0f\xb1\x01\x1f\x0f\x41\x36\x30\xe6\xe8\x51\x18\x4b\xe4\x86\xa4\xec\x95\x17\xfe\xee\xe9\x5f\x55\x8d\xc8\x15\x87\x94\x44\x49\xf0\xe6\x58\x60\x56\x64\x4f\x77\x75\x55\x77\xd5\xbf\x8e\x6e\x9e\x9d\x05\xaf\x66\xad\x0b\x2e\xdc\xd4\xcd\xab\xa5\x6b\x83\xfa\x26\xb8\x98\x7d\x5e\x4f\xa6\x6d\xb5\xac\x4e\x47\xbe\xc3\x62\x76\x3d\x6f\xdc\xe2\x9c\xcf\x4b\x77\xf5\xf6\xd2\xf7\x5b\x9c\x4d\xa6\x93\xe5\x59\xfb\xcf\x6a\x7a\x31\x3b\x3b\x6d\x67\xcd\xcf\x6e\x3e\xb9\x98\xce\xe6\x6e\x7b\xb7\xd7\xd2\xab\x9b\x5c\xee\xe8\xa3\x94\x3e\x6f\x66\x57\x6f\x67\x0b\x77\x7a\x73\x75\x39\xd0\x77\x5e\x4d\x2e\x17\xf7\xce\xaa\xbd\x76\x4e\xaa\x5d\x1e\x36\xe7\x75\x7d\x73\xff\x94\x74\xda\x3d\x23\x3d\x1e\x34\xe1\x62\x32\xad\x96\xf3\xea\xde\x39\xfb\x7e\x3b\xa7\xed\x3b\x3d\x68\xe6\xeb\xe9\xcf\xd3\xd9\xaf\xd3\x7b\x67\xee\xfb\xed\x9c\xb9\xef\x74\xdf\xcc\xb7\x9f\x4e\x2f\x66\xbc\x79\xfd\x6d\xf0\xcd\xb7\x3f\x04\x5f\xbd\xfe\xfa\x87\x4f\x46\xa3\xb7\x55\xf3\x73\x75\xe1\x56\xfd\x47\xa3\x89\x27\x34\x5f\x06\x47\xa3\x83\xc3\xfa\xc6\xb7\x1c\xfa\x0f\x50\x9f\xbb\xc5\xe2\xec\xe2\xb7\xc9\x5b\x1a\xba\xab\x25\x7f\x26\x33\xfd\xff\x6c\x32\xbb\x5e\x4e\x2e\xf9\x32\x93\x01\x6f\xab\xe5\x3f\xce\xe0\x9c\x0f\x34\x2c\x96\xf3\xc9\xf4\x42\xde\x2d\x27\x57\xee\x70\x74\x3c\x1a\x75\xd7\xd3\x26\x30\x8b\xf8\xce\x55\xed\x11\x1f\x82\xbf\xfd\x9d\x69\x4f\x82\x69\x75\xe5\x02\x1d\x76\x1c\x1c\xf5\xad\x6e\x3e\x9f\xcd\x8f\x83\xf7\xa3\x83\x8b\xdf\xe4\x5b\x70\xfe\x45\x00\x57\xa7\xdf\xb8\x5f\x21\xe2\xe6\x47\xc2\x36\xdf\xbf\xbc\xee\x3a\xff\x1d\xb2\xc7\xc7\xa3\x83\x49\x27\x03\x3e\xf9\x22\x98\x4e\x2e\x21\x71\x30\x77\xcb\xeb\xf9\x94\xaf\x27\x81\x17\xe9\xf4\x2b\xa8\x77\x47\x87\x10\x0a\x3e\xfd\xd7\x79\xf0\xe9\x2f\x87\xca\x89\xcc\xe5\x69\xfc\x3e\x1a\x1d\xfc\x52\xcd\x83\xfa\xba\x0b\x74\x1e\x9d\x64\x74\xf0\x93\xb2\xf3\x45\x30\x99\x9d\xbe\x9a\xbd\xbd\x39\xfa\xcc\xf7\x39\xf1\xbc\xf9\x51\xcd\xe5\x57\x3d\xa7\xa7\xaf\x2e\xfd\x3e\x1d\x79\xf1\x9f\x89\x1f\xc8\x28\xfd\x2d\x84\x7c\x47\xe5\xdb\x1a\x3d\x5b\xa7\x5f\xc2\xfa\xd1\xf1\x09\x3d\x46\xfe\xdd\xf2\xe6\xad\x0b\xaa\xc5\xc2\x2d\x59\xf2\xeb\x66\x09\x15\x91\xcf\xf6\xc3\x4f\x33\xed\x66\x41\x30\x5b\x9c\xfe\xd9\x6f\xeb\xd7\xfe\xcb\xed\x38\xdb\xc2\xbe\x7d\x8d\x82\xec\xa1\xff\xa7\xdb\x38\x3a\x58\x4c\x7e\x93\xef\x93\xe9\x32\x4f\x47\x07\x57\x40\x64\x70\x4b\xf4\x2f\xfe\xab\x34\xfe\xe0\x35\x24\x40\x4d\x4e\xf9\xc4\x3c\xa2\x2a\x47\xdd\xe4\xee\x5c\xc7\xc1\x37\x7e\x8a\xa3\x63\x9b\x81\x39\x4d\xca\x6e\x72\xca\xec\x7e\xf0\xf6\xb1\xdf\x7b\x76\xfc\x58\xe1\xe6\xc3\xa1\x30\xba\x73\x28\xbc\xfa\xa1\x6b\x9c\x7f\x48\x00\xd1\xee\x23\x80\x70\x9e\xc6\xad\xa0\x1b\x14\x4c\xfa\xed\x44\xbe\x5e\xbc\x9e\xcc\x3d\x89\x7a\x36\xbb\x5c\x1f\x5d\x5d\x2e\xee\x91\xfc\x66\xa1\x82\x7b\x7c\xa9\x1a\xf7\xfe\xf7\xb5\xd1\xa6\x12\x68\xf9\x4f\x40\xcd\x6b\xf1\x20\xaf\xd7\x30\xcb\x2b\xb9\x6a\xc5\xd1\xe1\x9b\x77\x51\xf7\xe6\x5d\x51\xbf\x79\x17\x16\xfe\x09\xed\x29\xdf\xbc\xcb\x9d\x6f\xb7\xb6\xce\xf7\x69\xe3\x37\xef\x52\xdf\xaf\xf1\xed\x8d\xff\x1e\xf3\xd9\x3f\x95\xff\xec\xc2\xb5\xf7\xad\xbe\x73\xc9\x5a\x1b\xfd\x1b\x4f\x2b\xf2\xf3\xf9\xf6\xd2\xd3\x77\xfe\x19\xfb\xa7\xf3\x4f\x9a\x79\x3a\xfe\x6f\xe6\xfb\x14\xe1\x1a\x1f\x36\x37\x4f\x36\x7e\xf3\x2e\xf1\xe3\xb3\x4e\x79\x88\xda\xf5\x7e\x87\x3d\x1e\x0d\x4b\x6c\xf6\x32\x84\x43\xbd\x55\xad\xe1\x98\x37\xc0\x2d\x2b\x77\xe2\x5f\x1d\x6e\x75\xf1\x87\xfe\xf5\xf1\xad\xba\x0f\x53\x80\x89\x3f\x89\xa5\xae\x33\x21\xa6\x7a\x8b\x87\x3b\x65\xb8\x0f\x77\x6e\xe1\x42\x0c\xde\x53\xbb\xa3\x3c\xef\x31\xab\xf3\x60\x97\x14\x01\xe6\x73\x1e\xc4\xe5\x49\x80\x1d\x9c\xaf\x9b\xc9\x51\x1a\x87\xc7\xd2\x8e\x76\x9f\xab\xf6\xff\x38\x9d\xbc\x3b\x8a\xd2\xbc\xcc\xd3\xb8\x4c\xb3\x93\x20\x3c\xf6\xc0\x56\x31\xfb\x67\x22\xeb\x7b\x11\xf0\x3c\x30\x39\x61\xed\x5c\xfe\xff\xfd\x76\x03\xaa\x93\x9d\x9a\x8b\x33\xda\x4b\x6f\x0b\xaf\x53\x65\xa4\x7a\x59\xf9\x77\xad\xd7\xbf\xc4\xbf\x8b\xfc\x53\x78\xbd\xeb\xc6\xaa\x87\x45\xa5\xfd\x72\x74\xd9\xd3\xcd\x73\xff\xd7\x7f\x8f\xfd\xfb\xd4\xb7\xc5\x99\xea\x30\x9f\xeb\xd4\xeb\x21\xef\xfc\x3c\xa9\x7f\x32\x74\x3e\x52\x9d\x4f\x7d\x9f\xcc\xeb\x7d\xe4\xc7\x35\xfe\xc9\x7d\x5b\x87\xee\xfb\xbf\x85\xef\x97\x41\xdf\xf3\x55\xfa\xcf\x75\xa4\xfc\xb4\xbe\xcd\x31\x9f\xe7\xaf\xf6\x73\xd7\x85\xfe\x6d\xfc\xb8\x2e\xd2\xbf\xd8\x4c\xee\xc7\xa5\xbe\x4f\xc2\xe3\x79\xe8\x7a\xdb\xf2\xe3\x9b\x52\xe7\xa9\x72\xb5\xb9\xd6\x7f\x2f\x9d\xca\x88\xad\x61\x5f\xf0\xcb\x1c\xd8\x58\xea\xe7\xad\x6a\x7d\x9f\x7a\x5a\x4d\xa8\xeb\x19\xf9\x3e\x15\x72\x7a\x3a\x39\x32\xb6\xba\xc6\xf0\x89\xdd\x55\xbe\xff\x98\x27\xd5\x3e\x51\xa9\xf3\xb3\x9e\xa1\x6f\xab\x22\xe5\x2d\x29\x75\x1c\xeb\x47\x7b\x92\xe9\xbe\x44\x9e\x46\xc9\x1e\xe4\xba\x4e\xd0\x19\x23\x7f\xad\xf3\x81\x27\x75\xa5\xfc\x8f\x3b\xe5\xa5\xf6\x7d\xc3\xb1\xae\x1d\xe3\x0b\x64\xcf\x95\x2e\x7b\xc4\x1a\x87\x7e\x7c\xd2\x29\x4f\x0e\x19\x12\xdd\xa3\xd2\xcf\x51\x1a\xf6\xe4\xec\x77\x6c\xfb\x11\xeb\xd3\x1a\x3f\xb4\x15\xa5\xea\x48\xe6\xbf\x47\x95\xae\x47\x5e\xa9\x8e\x84\xad\xf6\x6d\xa1\x91\xe9\x7e\xb2\xd7\x65\x6e\xba\xd2\xa9\x8e\x64\xac\x81\xed\x7f\x98\xab\x6c\x75\xa8\xb2\xc1\x77\xdc\x29\x0d\x64\x62\x4f\x42\xa7\x63\xe1\x3d\x63\x2f\xd0\x99\x9e\xff\x44\xf7\xb3\x40\x07\x23\xdb\x9b\x5c\x71\x12\x1d\x45\x5f\x5b\xe3\x8d\x36\xf4\x92\xf5\xe9\x9c\xee\x75\xd5\x2a\xbe\xa2\xd3\xd8\x0b\xfb\x86\xbe\xf2\x2e\xf7\xed\x6d\xa1\xfb\x34\x8e\xd5\x06\xd0\xd7\x22\xd1\xb9\xe0\x83\x77\xec\x6f\xea\x9f\xa4\x51\xbd\x62\x7d\x8b\x4e\xf5\x91\xf7\xe8\x27\x63\xb1\x29\xf6\x17\x7d\x41\x9e\x96\x7d\x8d\x54\x2f\x32\x78\x2e\x75\xcf\xe9\x0f\xfd\xdc\xec\x26\x6f\x74\x7d\x59\x53\xe4\xc1\x46\xd8\x77\x7c\x82\xcb\x74\xfd\xb0\xb9\xc8\xf6\x28\xa9\x54\x56\xf6\xae\x4c\xd5\x36\xf0\x09\xd8\x04\xeb\xc7\x9e\x61\x4b\xe8\x53\xec\xd4\xee\xc1\x04\x67\xfa\x9c\xd9\xba\xb0\x47\xae\x55\x3b\x84\x17\x7c\x0b\x36\xc4\xfe\x20\x2b\xf6\x97\x8f\x4d\xe7\xd1\xc3\x50\xf5\xa4\x32\x5d\x96\x77\xac\x77\xae\xf2\x30\x16\xfd\x71\x9d\xd2\x85\xa7\xc2\xa9\x9e\x66\x95\xda\x2d\xfe\x10\x9d\xad\xfd\xd8\xd2\x6c\x5e\xf4\x0d\x7b\xad\x74\x2f\xeb\x52\xf5\x94\xf6\x6a\x6c\xf8\x54\xab\x1d\x74\xe6\x2f\x59\x1f\xd6\xbe\x88\x74\x2f\x5c\xa4\x36\x8c\x1e\xd6\xd8\x69\xa1\x3a\x20\xef\xd9\xcf\x4e\x79\x86\x77\xf0\x90\x35\x16\x9d\xc6\xde\x63\x95\x17\xac\x64\xfd\xc1\x4d\xf6\x8e\xb5\x47\x96\x2e\x55\x3f\xdf\x25\x8a\x27\xe8\x10\x32\xb1\x4e\xcc\x11\x66\x9b\xbe\x3a\x8e\x75\x8c\xac\x39\xba\x9e\x99\xbd\xed\xf6\xd5\x60\xfc\xd3\x3d\x35\x54\x36\xfc\xf4\xea\xd5\x6e\x27\x4d\x8f\x7d\x5c\xf4\x1a\xeb\x2f\xe1\xa0\xd7\xd9\x37\xef\x9c\x97\xc9\x47\xe4\x9e\x5f\x69\xfe\xfa\xd7\xab\xcb\xbd\x9c\x34\x4e\x00\xa5\x6c\x70\x00\xde\x10\x9a\x78\xe5\xa4\x53\x73\xd2\x5d\xab\x4e\x1a\x10\xc0\x59\xa1\x60\xd0\x06\x54\x8a\xde\xb0\x2a\x05\x7c\x71\xf4\x8d\x02\x6c\x54\x6b\xf0\x48\x3b\x00\x89\xe3\x83\x07\x80\x14\x10\xa3\x1d\x20\xcf\x6b\x9d\x03\x63\x03\x6c\x72\x73\xc2\xf0\x00\x2d\x80\xa4\x36\xc3\x19\x9b\xf1\xa2\xfc\xe2\x04\x33\x0b\x34\x4a\x75\x48\xf0\x01\xe8\x01\x2a\x18\x0d\xc6\xdf\x19\x80\x00\xd6\x38\x28\xe6\x91\xb6\xd4\x82\x85\x5c\x0d\x0a\x50\xc6\x60\x04\xd0\xe8\x5b\x29\xd8\x13\x5c\x08\xf0\x77\xea\x18\x30\x7a\xe4\x91\xa0\x7a\xac\xe0\x81\xbc\x80\x51\x62\xa0\x00\x38\xe2\x38\xc3\x46\xc1\xaa\xb2\x20\x06\x10\x41\xae\xd2\x9c\x13\x63\x64\x8d\x22\x5d\xd3\xda\xc0\x80\x7e\xf0\x90\x99\xf3\x21\xc8\x69\x0d\x8c\x00\x21\xf6\xb1\x8e\x55\xd6\xde\xa9\x03\xd0\xac\x4d\x62\xc1\x16\xe0\x46\xdf\x88\xb5\xf7\xef\xc2\x4a\xe9\xe0\x38\x45\xf6\x46\xc1\xcc\x39\xdd\x5f\xd6\x92\xa0\xa6\x2c\x14\x48\x01\x1c\x64\x10\x47\x5c\xaa\xa3\x40\x3e\x9c\x1a\x80\x06\xc8\xe3\x20\xd0\x0b\xc0\x98\xc4\x20\x4f\x15\x48\x63\x73\x0c\x61\x34\x00\x52\x99\xf2\x81\x9e\xb1\xee\x00\xdc\x03\x12\x8a\x95\xa6\x3f\x1d\xaa\x56\xb4\x36\x00\x6b\xb3\x2e\xb4\x1b\xb8\x56\xa4\xf6\x81\xaf\x0d\xa1\x5e\x02\xc4\x86\x44\xea\x53\x8d\xf4\x8f\x06\xb3\xef\x28\x79\x3e\x4b\x92\x8c\x1d\x12\x28\xc5\x95\x62\x86\x24\xc0\x16\x44\xe1\x54\xd7\xfb\x0c\x25\xd3\xd0\x4a\x62\x0d\x8e\xd1\xc9\xa6\xd2\x80\x1d\x9d\x1e\x57\xea\xe4\x99\xbb\x04\x87\x9c\xda\x94\x60\x9b\x53\x9b\xca\x9c\x06\x5e\x04\x3b\xe0\x0e\xfd\x99\x1b\x5c\x05\x8f\xe0\x4b\x02\x82\xb1\xce\x8b\xad\x83\x23\x04\x85\x62\x2f\x91\x05\x9c\x16\x50\x13\xc0\x4b\x70\xda\x07\x35\xb5\xbe\x63\x5c\x6c\xc1\x8c\x24\xed\x86\x95\x77\xed\xac\xb2\xc4\xa0\xcc\x14\x27\xe0\x69\x8b\x9d\x6d\x6c\xc2\x7e\x26\xb6\x41\x66\x65\x5d\x03\x25\xf2\x4d\xbb\xda\x18\xff\x50\x93\xda\xc6\xff\xb3\x5a\xd3\xa0\x08\x66\x47\xe3\xe2\xb1\x66\x94\x65\x65\x91\x8f\xe3\x17\x30\xa3\xbd\x33\x76\x32\x02\x32\x3d\x9c\x2f\xca\x82\xe3\xea\x83\x01\xd7\xa8\x93\xc6\x49\xa0\xbc\xad\x55\x84\x30\x0c\xa2\x6c\x14\xb2\x2a\xb4\x22\x05\x0d\xbe\xe3\x70\x70\xc8\x18\x93\x38\x9a\x5a\x9d\x87\x8b\xd5\xd9\x62\x34\x91\x29\xba\x64\x87\x63\xe5\xa1\x35\x43\xc3\x49\x10\x89\xe3\x6c\x30\x46\xa2\x65\x67\xce\x5e\xb2\x9f\x44\x9d\xf9\xd8\xa2\x61\x31\xf2\x56\x8d\x1f\x79\x12\xab\x16\x90\x51\x96\xb5\x66\x81\xe3\xc2\xb2\x66\x4f\x27\xce\x35\x10\x68\x2c\x50\x10\xa7\xdc\xa9\xbc\x18\x26\x55\x02\x9c\x21\x41\x0f\xbc\x8c\x73\xe5\x9b\xe8\x9e\x48\x3c\x8a\x34\xe0\x69\x2c\x33\xc7\x09\x12\x48\x51\x45\x28\xcd\xe9\x13\x50\x38\xcb\x3e\x25\x18\x29\x15\x24\xc8\x1a\x00\x92\x3e\xdb\x26\x38\x02\x9c\x3a\xcb\x0e\x25\x8b\xb4\x4c\x9e\xcc\xc9\x59\x76\x41\x1b\x00\xc3\x7a\x21\xb3\x64\x5a\xb5\xce\x8b\x5c\x64\xb8\xfc\x65\x5d\xea\x46\x03\x11\x00\xac\xb0\x0c\x95\x4c\x93\xbd\x21\x38\x42\x6e\x32\x8d\xb2\xd3\x75\x20\x98\x63\x1e\xc9\xda\x33\xcb\xfa\x32\x05\xc5\xc2\xb2\x9c\xdc\xd6\x81\xf9\x01\x57\xf6\x9f\x3e\xce\xe6\xa1\x0f\x7a\x40\x06\x8e\x1e\xd1\x97\x2a\x88\xe8\x4f\xa5\xa0\x8b\xce\xb1\x7e\xcc\x45\xa0\x01\x80\xca\x7e\xa1\x2b\x16\xfc\xb1\xe7\xe3\x46\xf7\x3c\xb3\xec\xbb\xb0\xea\x46\x42\x40\x9a\x2b\x1d\xf6\x29\xb1\x0c\x15\xdd\x24\x58\x41\x0f\xa4\xd2\x62\x95\x0d\x32\x7d\x74\x94\x75\xcf\xac\x2a\x83\x5e\xd4\xf6\x19\x20\x8d\x4d\xb7\x33\xfb\x2c\xfb\x57\x5b\xc6\x59\x69\x40\x05\x80\x4b\x80\x5b\x69\xd0\x99\x9b\x3e\xb3\xe6\x04\x60\xac\x35\x73\xc7\x89\x56\x77\xd0\x31\x82\x2f\xc9\x42\xc7\x56\x55\x6a\xf5\x3d\xce\x08\x3d\xc3\xf1\xb0\xe7\xa5\xe9\x07\x3a\x00\x5d\x1c\x05\xf2\xa3\xb7\xac\x49\xd8\x6e\x02\x3c\xba\x01\x3f\xec\x67\x1f\xe8\xae\x02\xae\x6d\x00\xbf\x7f\xb2\x77\x87\xc8\x5d\x70\xdf\x95\xea\xdd\x19\xba\x07\xae\xbf\x54\xa2\xb7\xc9\xbb\x41\x7a\x5a\xe6\x8f\xc6\xf4\x28\x8f\xb3\x3c\x79\x01\x4c\x7f\x62\x9a\xd7\x87\x1e\x69\xab\x69\x83\xb3\x30\x07\xa4\x96\xba\x94\x85\x4c\x68\x28\x21\x0b\xc8\xc7\x7b\xea\xa3\xd4\xdb\x40\x00\xd0\x96\x9a\x05\xe9\x46\x66\xdf\x49\x1b\x48\x61\x24\x0d\xac\x56\x88\xcf\xdf\xa6\xaf\x99\x58\x5d\x45\x52\x9f\x4c\x11\x23\xcc\x56\xe7\x0f\xb1\xb5\x81\xc6\x3c\xa4\x63\x7d\x9f\xd4\xfa\xc5\xf6\xb7\xa7\x09\x0a\x50\xc7\xa3\x9d\xcf\x20\xb0\xa0\xae\xd5\x01\x79\xb0\x52\xf1\x30\x16\x02\x81\x20\xd4\x42\xa2\x1e\x1d\x52\x4d\xd5\x0a\xab\x91\x45\x86\xb4\xa5\x7d\xee\x1f\x52\x1c\x64\x90\xfa\xb3\xa1\xab\x58\x9d\xa5\x7f\xf1\x40\xe8\x85\xc7\x80\x36\xf5\x18\x10\x16\xe4\xba\x3f\xf4\x7a\x6a\x86\x33\x48\xea\xae\x95\x3e\x24\xbf\x19\x24\xb4\x87\xcd\xbe\x6c\x76\xb3\x5d\x1e\xb3\xe0\x28\x89\xff\x68\x0b\xbe\xae\x6f\xfe\xa3\x72\x9b\xad\x0a\x5d\x6a\x5d\x86\x7c\x67\x6c\x87\x01\xdb\x14\xfa\x8e\xcc\x7b\xea\xf2\x1d\x2a\x6b\x6a\xbc\x71\xf1\x65\x40\x81\xef\x8c\x7e\xb0\xee\x0e\xf3\xfe\xbc\x6a\x3b\xc0\xbf\x29\x6c\x12\xfe\xd1\x59\xc4\xad\xfc\x7b\x27\x11\x72\x94\xdc\x99\x96\x3a\x0b\x9e\xfb\x63\xbf\x54\xe1\x90\xa0\x0e\xc8\x8e\xed\xe8\x0d\x2d\x25\x00\x96\xca\x53\xa1\x5a\x4d\x10\x54\x56\x1a\xb0\x77\x76\x94\x81\x5b\x91\xa0\xae\x35\x57\x63\x41\x7d\x69\xc7\x4f\xd0\xca\xac\x9a\x45\xe0\x86\x1b\xe3\x88\x25\x8f\xf5\xd8\x80\xe0\x3e\x37\xc8\x27\xb1\xe0\x28\x66\x6c\x47\x54\xf0\x4b\x35\x54\x2a\x83\x63\x0b\xf0\xec\x78\x3c\xb5\x52\xbd\x04\x9b\x91\xba\x17\x02\xbf\xb1\xb9\x13\x82\x61\x92\x18\xdc\x20\x09\x48\x9f\x5c\xc4\xe6\x66\x38\x0a\x20\x20\x96\x63\xcc\xb1\xba\x19\xf8\xe7\x88\x87\x8c\x9f\x4a\x20\x01\x29\x56\xc9\x43\x05\x0d\x37\x28\xd5\xc5\x58\x83\x67\x82\x5e\xac\x97\x6a\x9d\x1c\x99\x45\x6a\xbd\x49\xa1\x95\x88\xca\xf8\xa7\xea\xda\xda\x31\x22\x95\x5d\x39\x32\x0d\x8d\x66\xa8\xf2\xca\xbc\x16\x84\xb2\x86\x7d\x70\x9c\x58\x32\x23\xc7\x87\xb5\x06\xec\xf1\x5a\xe2\x40\xa0\x1c\x76\xba\x27\xd0\x97\x4a\x64\xbe\xaa\x7c\xc2\x5f\x61\x2e\x92\x84\xc7\xd9\x71\x1b\xe3\x59\x13\x67\x95\x5e\xda\x9c\x55\x50\x2a\x4b\x26\xe5\x1a\x42\xad\x68\x24\xc7\x2f\x4e\xd1\x8e\xbd\xa3\x8d\xb9\x3a\x3b\xe2\x12\x54\x22\xa0\x4e\x74\x5f\xf9\x2c\x47\x60\xa5\x86\x2b\xad\x25\x85\x1c\xdb\xca\x31\xa3\x5d\x95\xe0\x88\x54\xaa\x27\x89\x26\x1d\x61\xb2\x89\x74\xe8\xb8\x1c\x01\xf5\x6e\xbe\xda\x1e\x54\x7f\x60\x2d\x4f\xc5\xb9\x3b\x21\xf5\x87\x37\xf7\x76\x41\xdc\xa3\x02\xea\x21\x96\x9f\x1f\xde\x06\xc2\xe9\xb8\x7c\x7c\x89\xe4\xa5\x9c\xf1\x73\x1c\x9a\x34\x9a\x73\x63\xc6\x85\xdd\x6c\xa0\xd8\x0f\xb4\x38\x2b\x7e\x4b\xf4\x1a\xda\x29\x68\xa9\xf0\x86\xb9\x13\x39\x92\xd3\x31\x06\xd3\x0b\xcd\x21\x73\x32\x8a\x0a\x03\x4b\x52\xcc\x6f\x55\x95\x31\x25\xcc\x0a\xd8\xc3\x04\x81\x18\x4c\x93\x82\xbc\xa8\x6d\xac\xa7\xa9\x40\x01\xe6\x49\xae\x88\xa9\x93\xc7\x53\x6c\x94\x39\xec\x44\x5c\x4e\xef\x1b\xad\x0d\xc0\x27\x79\xad\x40\xa2\x9d\x94\x03\x79\x38\x7a\xe0\x80\xda\x44\x64\x27\xad\xe4\xd0\x64\x10\x72\xf2\xeb\x34\x9a\xa6\x1e\x11\xda\x41\x05\x7d\xf9\x2c\xb5\x94\xd6\xf2\xe8\x62\x55\x4b\x90\xc3\x1e\xa7\x32\xca\xc9\xb7\x53\x5e\x31\x79\xa0\x9e\x4c\x83\x7a\x0d\xa6\x49\xdd\x24\x32\x9e\x71\x07\xd4\x84\x90\x4d\x0a\xac\x8d\x41\x59\xa1\xb4\x2a\xab\x61\xf0\x50\x0b\x91\x43\x8e\x5a\x0f\x2e\xe4\xc4\x3e\xd2\x35\xed\x8c\x27\xfa\x13\x08\x51\x94\x15\xf8\xb4\x9b\x13\x9d\x9d\xde\x03\x37\xa1\xdd\x0a\xe0\x04\x16\xf8\xe0\xb6\x47\x68\x27\xef\xf0\x4b\xed\x89\x1a\x55\xdd\x0c\x43\x48\x65\x07\x1f\x92\x93\x67\xb6\x4e\xf7\x05\x4b\x4f\x0e\xfe\x07\x28\xdd\x81\x93\x07\x85\xfe\x03\x64\x1e\x0f\x2e\x2f\x1c\xf8\x6f\x13\xa6\x87\x9a\xf0\xf1\x71\xd4\xf3\x42\xcd\xf7\x7a\xbd\xf9\x7f\x2d\xf4\x1f\x10\x7b\x3f\x65\x1e\x20\xb4\xd2\xe5\xc1\x8b\xe8\x9b\x9a\x3c\x40\xe3\xa1\x8a\xbc\x5d\x8e\x67\xd5\xe3\x2d\x82\x7c\x2c\xc9\xc0\x07\xab\xf0\xf4\x7c\x20\xb2\x77\xe9\x5a\x3e\x90\xdf\xc9\x07\x52\x3d\xe9\xef\xf3\x01\xc0\x19\x67\x88\xe3\x21\xbe\x25\x9e\x45\xed\x01\xe3\xda\x3e\x4b\x41\x3a\xd4\x2b\x40\x38\xd4\xd4\x00\xb9\xe9\x9d\x64\x6e\x31\x9d\x15\x53\x2b\x8b\x43\xa5\xcc\x65\x05\x79\xf8\x80\x57\x67\xd7\x06\x71\x3a\xc4\x89\x94\x7c\x28\x2e\x4b\x6c\xdd\x68\x19\x49\xae\xf7\xf5\xa7\x79\xe6\xf4\x70\x08\x99\x9d\xf6\xb5\x76\x95\xb6\xb1\x6b\x83\x72\x18\xe2\x74\x9d\x70\x1a\x72\xdd\xcb\xae\x33\x16\x56\x28\x97\x6b\x64\xc5\xaa\x88\x2b\xb7\x00\x42\xa5\xcb\x38\xae\x33\x91\x6b\xe0\xcc\x24\xc7\x19\x6b\x3c\x1c\x5b\xde\x12\x5b\xf1\x1f\xfe\xe4\x6a\x53\xa7\x65\xbd\xc2\x6e\x07\xb4\x76\x03\xa0\xb5\x6b\x66\xe3\xb1\xe6\x30\x38\x6a\x39\x3d\xb4\x20\x23\x31\xf3\x66\xbd\x98\x4b\x0a\xe5\xa5\xf6\xc3\x01\x13\xb3\xd3\xa7\xe9\xaf\xd4\xd9\xad\x85\xce\xe2\x75\x67\xf9\x81\x5c\x93\x8c\x34\x6f\x61\xbc\x38\x77\x67\x87\x17\x76\x7d\x2c\xb4\x6b\x6e\xec\x69\x62\xd7\x29\xe5\x30\xc1\x69\x1b\x3c\xd1\x9f\x3c\x4e\xf2\xc0\x4c\x65\x90\x98\x3e\xd7\x7d\xe1\x00\x86\x1c\xc0\x59\x3e\x20\x37\x49\xba\xd5\x75\x3a\xda\xe0\x99\x3d\xa2\x5c\x89\x53\x96\xdc\xa3\xd5\x35\xe6\x9d\x5c\xcb\x4c\x56\x87\x15\xe8\x2a\xa5\xc4\xa1\xeb\x55\x04\x33\xac\x4b\xaf\x3b\x72\x45\x73\x38\x37\xd8\x30\x9e\x67\x00\xc2\x0f\x33\x84\xcd\xdf\xd9\xdc\x83\x81\x8f\xc9\x13\xb6\xb1\xff\x22\xf8\x37\x90\x2d\x24\x61\xf4\x31\xb9\xf0\xff\x97\xdf\xff\x5b\xcb\xef\x5b\xb6\xf9\x19\xac\x75\x28\x0c\xdf\xfe\xb3\xb7\x7b\x6c\xf7\xf1\xc1\xf8\x6e\xc1\x5e\xc4\x8e\x3f\xea\x52\xfc\x8f\xfa\xbb\xbf\xe7\xfb\x39\xce\x8e\x9f\xdb\xa0\x7b\xa1\xdd\x8c\x1c\xf2\x21\xc4\x3e\xb1\xdd\x44\x14\xbb\x1d\xd6\xcd\x01\x96\xf7\xd3\xcb\x01\x42\x2b\x9d\x1c\xfc\x75\xe5\xa6\x3a\x0e\xd0\x78\xa8\x2a\x6e\x97\xe3\x59\xd5\x70\x8b\x20\xbd\x06\xee\x15\x4e\x8f\x8b\xf0\x45\x14\x70\xef\x70\x9a\x30\x90\x90\xad\xbf\x2f\xc3\x7d\x06\xf9\x35\x4c\xbb\x72\x23\xa8\x25\x35\x8f\xda\xee\x6e\x10\x46\xcb\xaf\x6f\xac\x4e\x85\x7a\xe2\x82\xa0\x25\xf5\x20\xbb\x93\x21\xb0\x1d\xea\x25\x57\xfe\xa6\xa6\xce\xb1\xdd\xeb\xd9\xa6\xd2\x84\xa9\xfc\x6a\x22\xb7\x3b\x23\x51\xf7\x30\x95\xde\x3f\x2c\xda\x20\xb3\xa9\xce\xbb\xc2\xa2\x8d\xe1\x7b\x69\xf2\x4b\x85\x45\x43\x12\xf4\x61\xd1\xe3\xa3\xa2\x17\x54\xe2\x27\x46\x45\x14\x22\x52\x8b\x7a\x78\xf8\x01\xc5\xbf\x03\x00\x00\xff\xff\x00\x8f\x53\xf6\x00\x40\x00\x00")

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

	info := bindataFileInfo{name: "templates.go", size: 32768, mode: os.FileMode(420), modTime: time.Unix(1469649918, 0)}
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

