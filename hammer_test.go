package hammer

import (
	// _ "github.com/kingzcheung/hammer/testdata/hammer"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

const indexText = `<!DOCTYPE html><html lang=en><head><meta charset=utf-8><meta http-equiv=X-UA-Compatible content="IE=edge"><meta name=viewport content="width=device-width,initial-scale=1"><link rel=icon href=/favicon.ico><title>Translation Management System</title><link href=/css/about.a8f98c3c.css rel=prefetch><link href=/js/about.a793be22.js rel=prefetch><link href=/css/app.e0ce1c24.css rel=preload as=style><link href=/css/chunk-vendors.05f5edda.css rel=preload as=style><link href=/js/app.3258337c.js rel=preload as=script><link href=/js/chunk-vendors.a608caf2.js rel=preload as=script><link href=/css/chunk-vendors.05f5edda.css rel=stylesheet><link href=/css/app.e0ce1c24.css rel=stylesheet></head><body><noscript><strong>We're sorry but web doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id=app></div><script src=/js/chunk-vendors.a608caf2.js></script><script src=/js/app.3258337c.js></script></body></html>`

func TestFind(t *testing.T) {
	as := assert.New(t)
	fs, err := New("dist")
	if err != nil {
		t.Error(err)
	}

	b, err := fs.Find("/index.html")
	b1, err := ioutil.ReadAll(strings.NewReader(indexText))
	if err != nil {
		t.Error(err)
	}
	as.Equal(b1, b)
}
