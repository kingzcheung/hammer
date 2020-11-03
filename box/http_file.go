package box

import (
	"bytes"
	"os"
)

type httpFile struct {
	file
	reader *bytes.Reader
}

func (f *httpFile) Read(p []byte) (n int, err error) {
	return f.reader.Read(p)
}

func (f *httpFile) Close() error {
	return nil
}
func (f *httpFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
func (f *httpFile) Seek(offset int64, whence int) (int64, error) {
	return f.reader.Seek(offset, whence)
}
func (f *httpFile) Stat() (os.FileInfo, error) {
	return f, nil
}
