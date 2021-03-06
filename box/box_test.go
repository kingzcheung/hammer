package box

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

var pwd string

const zipData = "PK\x03\x04\x14\x00\x08\x00\x08\x00\x98k:Q\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\n\x00	\x00index.htmlUT\x05\x00\x01\x11Bo_\x8cS_o\xd30\x10\xff*\xc7^\xf6B\x92\xade\xac\x93lKh\xec\x01$\xc4\xa4\x0d\x01\x8f\x17\xfb\xd2\xb8ulc_R\xf5\xdb\xa3%\xa5k\xc74\xf5%\xd2\xc5\xbf\x7fw>\x8bw\x9f\xbf\xdf>\xfe\xbe\xbf\x83\x96;\xa7\xc4\xd3\x17\x1c\xfa\xa5$\xafDKh\x94\xe8\x88\x11t\x8b)\x13\xcb\x9e\x9bb\xb1\xfb\xd72\xc7\x82\xfe\xf4v\x90\xbf\x8a\x1f\x9f\x8a\xdb\xd0Ed[;\x02\x1d<\x93gy\xf6\xe5N\x92Y\xd2\xd9\x8e\xe2\xb1#9X\xda\xc4\x90\xf8\x19\xb5\xb1\x86[ih\xb0\x9a\x8a\xb1xo\xbde\x8b\xae\xc8\x1a\x1d\xc9\xcb3%\x9c\xf5kH\xe4\xa4\xd5\xc1C\x9b\xa8\x91U\x83\xc3SUZ\x1d\x94`\xcb\x8e\xd4cB\x9f\x1d\xb2\x0d\x1e\xbe\xa1\xc7%u\xe4\x19\x1e\xb6\x99\xa9\x13\xd5\x04\x9a\xc4&\x0d\x9ds\x85u\xe8\xb9\xc4Es\xb3\xd0s]\xea\x9cG\xa7\x98\xa8!\xd6\xed\x11~\xb5\x87_\xdf\xcck\x9a\xcd\xca\xd5\x1b\xe8Q=\xc6\x92.4]\xea\xd9\x87Cm\x17\xd0\x00f\x99y\xfbJ&\xdd\xf6~]\x0c\xe4MH\xb9\xbc\xb8j\xae\xc8\x18<\x8d\xbf\x9aL\xe7\xb3\xab\xc5|~\xad\x0f\x12\xee9:\xd9\xc8/I\xc7\x9e\xf8\xf1b\xa1\xb1\x99\x9dF?!\xf3\x184\xb7D\xff3_\x9d\xd1!\xbe\x9a\xb6\xb1\x0ef\xab\x84\x0f\xff\x02dN\xc1/\xd5O:O\x049\xa4\xb4\x85\xbag\xd8P\x0d&P\xf6\xe7\x0c\x9b\x90\xd6\x10S\x88\x94\xdc\x166\x96\xdb\xd03|\xc5\x01\x1fF\x11 \x8f\xb5#S\xc2\xbd#\xcc\xb4\xab\xc12p\x18\x97\xd4\xfa\x9eJQ\xed\xbcD\xf5lo\xec\x00\xd6H\x8cQ\x89\xca\xd8A\x89\xe9\x04r\xd2oOT\x89j\xdf\xc31\xe5\xc5\xcd\x1d\x00\xab\xa9\xfbj|\xac\x7f\x03\x00\x00\xff\xffPK\x07\x08\xa5\xbd\xb3)\xbc\x01\x00\x00\xbc\x03\x00\x00PK\x01\x02\x14\x03\x14\x00\x08\x00\x08\x00\x98k:Q\xa5\xbd\xb3)\xbc\x01\x00\x00\xbc\x03\x00\x00\n\x00	\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa4\x81\x00\x00\x00\x00index.htmlUT\x05\x00\x01\x11Bo_PK\x05\x06\x00\x00\x00\x00\x01\x00\x01\x00A\x00\x00\x00\xfd\x01\x00\x00\x00\x00"

const indexText = `<!DOCTYPE html><html lang=en><head><meta charset=utf-8><meta http-equiv=X-UA-Compatible content="IE=edge"><meta name=viewport content="width=device-width,initial-scale=1"><link rel=icon href=/favicon.ico><title>Translation Management System</title><link href=/css/about.a8f98c3c.css rel=prefetch><link href=/js/about.a793be22.js rel=prefetch><link href=/css/app.e0ce1c24.css rel=preload as=style><link href=/css/chunk-vendors.05f5edda.css rel=preload as=style><link href=/js/app.3258337c.js rel=preload as=script><link href=/js/chunk-vendors.a608caf2.js rel=preload as=script><link href=/css/chunk-vendors.05f5edda.css rel=stylesheet><link href=/css/app.e0ce1c24.css rel=stylesheet></head><body><noscript><strong>We're sorry but web doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id=app></div><script src=/js/chunk-vendors.a608caf2.js></script><script src=/js/app.3258337c.js></script></body></html>`

func Test_box_FindString(t *testing.T) {
	ReadFromZipData("testIndex", zipData)
	as := assert.New(t)
	fs, err := UzipFromNamespace("testIndex")
	if err != nil {
		t.Error(err)
	}
	s, err := fs.FindString("/index.html")
	if err != nil {
		t.Error(err)
	}
	s1, err := ioutil.ReadAll(strings.NewReader(indexText))
	if err != nil {
		t.Error(err)
	}
	as.Equal(string(s1), s)
}
