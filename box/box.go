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

const DefaultName = "default"

var ErrContOpenDir = errors.New("can not open dir")
var zipdata = map[string]string{}

type file struct {
	os.FileInfo
	data []byte
}

type box struct {
	files map[string]file
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

func (b *box) Bytes(file string) ([]byte, error) {
	f, err := b.Open(file)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func (b *box) Strings(file string) (string, error) {
	b1, err := b.Bytes(file)
	if err != nil {
		return "", err
	}
	return string(b1), nil
}

func ReadFromZipData(data string) {
	zipdata[DefaultName] = data
}

func UzipFromNamespace(ns string) (http.FileSystem, error) {
	zdata, ok := zipdata[ns]
	if !ok {
		return nil, errors.New("zip data is null")
	}
	return uzip(zdata)
}

func uzip(data string) (http.FileSystem, error) {
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
