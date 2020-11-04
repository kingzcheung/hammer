#  hammer - a simple tool to embed files in Go binary

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

In Your code,your need to import the generated package, and init hammer.

```go
import (
	_ "./hammer" // your package name
)

// ...
fs,_ := hammer.New("/public")
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

