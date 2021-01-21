#  hammer

 hammer is a simple tool to embed files in Go binary.


see [https://github.com/golang/go/issues/43407](https://github.com/golang/go/issues/43407)

## Installation

To install command line tool

```
go get -u github.com/kingzcheung/hammer/hammer
```

To get go  library

```
go get -u github.com/kingzcheung/hammer
```

### Usage

In CLI:

```bash
hammer path/to/your/project/public
```

In Your code,your need to import the generated package,  init hammer and serve.

```go
import (
	_ "./hammer" // your package name
)

// ...
fs,_ := hammer.New("public")
http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(fs)))

```

In gin:

```go
route.Use(func(c *gin.Context) {
		fs, _ := hammer.Assets()
		fileserver := http.StripPrefix("/", http.FileServer(fs))
		if fs.Exist(c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	})
```

Or read the content of a single file:

> But does not support read directory!

```go
// dist
// ├── css
// │   └── about.a8f98c3c.css
// ├── favicon.ico
// ├── fonts
// │   ├── element-icons.535877f5.woff
// │   └── element-icons.732389de.ttf
// ├── hammer
// │   └── hammer.go
// ├── img
// │   ├── excel.fc9b920c.jpg
// │   ├── gujia.846046c3.png
// │   └── no_pic.bffd1360.png
// ├── index.html
// └── js
//    ├── about.a793be22.js
//    └── about.a793be22.js.map

fs,err := hammer.New("dist")
if err != nil {
  panic(err)
}
fs.Find("/index.html") //[]byte
fs.FindString("/js/about.a793be22.js")//string
```

#### Hammer single file

```
$ hammer /path/to/single/file.html
```

You only  read the content by filename:

```go
fs.Find("/file.html")
```

