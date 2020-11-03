package box

import (
	test "github.com/kingzcheung/hammer/testdata/hammer"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var pwd string

func Test_newBox(t *testing.T) {
	_, err := uzip(test.ZipData)
	if err != nil {
		t.Error(err)
	}
	as := assert.New(t)
	sourceData, err := ioutil.ReadFile(`../testdata/dist/css/about.a8f98c3c.css`)
	if err != nil {
		t.Error(err)
	}
	as.Equal(sourceData, zipdata["css/about.a8f98c3c.css"])
}

func TestUzipFromNamespace(t *testing.T) {
	ReadFromZipData(test.ZipData)
	fs, err := UzipFromNamespace(DefaultName)
	if err != nil {
		t.Error(err)
	}
	f, err := fs.Open("/index.html")
	if err != nil {
		t.Error(err)
	}
	b, _ := ioutil.ReadFile("../testdata/dist/index.html")
	b1, _ := ioutil.ReadAll(f)
	as := assert.New(t)
	as.Equal(b, b1)
}
