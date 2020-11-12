package box

import (
	"archive/zip"
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const DefaultNameSpace = "assets"

var ErrContOpenDir = errors.New("can not open dir")
var zipMap = map[string]string{}

type file struct {
	os.FileInfo
	data []byte
}

func GetZipData() map[string]string {
	return zipMap
}

type ServerFileSystem interface {
	http.FileSystem
	Find(path string) ([]byte, error)
	FindString(path string) (string, error)
	Exist(path string) bool
}

type box struct {
	files     map[string]file
	namespace string
}

func (b *box) Open(name string) (http.File, error) {
	f, ok := b.files[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	if f.IsDir() {
		return nil, ErrContOpenDir
	}
	return &httpFile{file: f, reader: bytes.NewReader(f.data)}, nil
}
func (b *box) Exist(path string) bool {
	for filename, _ := range b.files {
		if filename == path {
			return true
		}
	}
	return false
}

func (b *box) Bytes(file string) ([]byte, error) {
	f, err := b.Open(file)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func (b *box) Find(path string) ([]byte, error) {
	return b.Bytes(path)
}

func (b *box) Strings(file string) (string, error) {
	b1, err := b.Bytes(file)
	if err != nil {
		return "", err
	}
	return string(b1), nil
}

func (b *box) FindString(path string) (string, error) {
	return b.Strings(path)
}

func ReadFromZipData(ns string, data string) {
	zipMap[ns] = data
}

func UzipFromNamespace(ns string) (ServerFileSystem, error) {

	zdata, ok := zipMap[ns]
	if !ok {
		return nil, errors.New("zip data is null")
	}
	return uzip(zdata)
}

func uzip(data string) (ServerFileSystem, error) {
	dlen := int64(len(data))
	read, err := zip.NewReader(strings.NewReader(data), dlen)
	if err != nil {
		return nil, err
	}

	bobj := new(box)

	files := make(map[string]file, len(read.File))
	for _, r := range read.File {
		rc, err := r.Open()
		if err != nil {
			return nil, err
		}
		fobj := file{}
		fobj.FileInfo = r.FileInfo()
		d, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, err
		}
		fobj.data = d
		files["/"+r.Name] = fobj
		_ = rc.Close()
	}
	bobj.files = files

	return bobj, nil
}
