package hammer

import (
	"github.com/kingzcheung/hammer/box"
	"io/ioutil"
	"net/http"
)

func New(relativePath string) (box.ServerFileSystem, error) {
	return box.UzipFromNamespace(relativePath)
}

func Assets() (box.ServerFileSystem, error) {
	return box.UzipFromNamespace(box.DefaultName)
}

func Bytes(fs http.FileSystem, file string) ([]byte, error) {
	f, err := fs.Open(file)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func Strings(fs http.FileSystem, file string) (string, error) {
	f, err := Bytes(fs, file)
	return string(f), err
}
