package hammer

import (
	"github.com/kingzcheung/hammer/box"
)

func New(relativePath string) (box.ServerFileSystem, error) {
	return box.UzipFromNamespace(relativePath)
}

func Assets() (box.ServerFileSystem, error) {
	return box.UzipFromNamespace(box.DefaultName)
}

func Find(fs box.ServerFileSystem, file string) ([]byte, error) {
	return fs.Find(file)
}

func FindString(fs box.ServerFileSystem, file string) (string, error) {
	return fs.FindString(file)
}
