package box

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_zbox_write(t *testing.T) {
	zbox := NewZbox("../testdata/dist")
	zbox.SetNamePackage("hammer")
	err := zbox.write()
	if err != nil {
		t.Error(err)
	}
	err = zbox.codeGenerate()
	if err != nil {
		t.Error(err)
	}
}

func Test_zbox_write_file(t *testing.T) {
	zbox := NewZbox("../testdata/dist/index.html")
	zbox.SetNamePackage("hammer")
	zbox.SetForceRemove(true)
	err := zbox.write()
	if err != nil {
		t.Error(err)
	}
}

func Test_zbox_SetNamePackage(t *testing.T) {
	as := assert.New(t)
	zb := NewZbox("")
	zb.SetNamePackage("aIbC")
	as.Equal("a_ib_c", zb.namePackage)
	zb.SetNamePackage("Store")
	as.Equal("store", zb.namePackage)
	zb.SetNamePackage("Test_zbox_Get")
	as.Equal("test_zbox__get", zb.namePackage)
}
