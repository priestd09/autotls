# autotls

[![Build Status](https://travis-ci.org/go-ego/autotls.svg?branch=master)](https://travis-ci.org/ego-ego/autotls) 
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ego/autotls)](https://goreportcard.com/report/github.com/go-ego/autotls)
[![GoDoc](https://godoc.org/github.com/go-ego/autotls?status.svg)](https://godoc.org/github.com/go-ego/autotls)
[![Join the chat at https://gitter.im/go-ego/autotls](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/go-ego/autotls?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Support Let's Encrypt for a Go server application, based on [gin-cors](https://github.com/gin-gonic/autotls).

## example

example for 1-line LetsEncrypt HTTPS servers.

[embedmd]:# (example/example1.go go)
```go
package main

import (
	"log"

	"github.com/go-ego/autotls"
	"github.com/go-ego/ego"
)

func main() {
	r := ego.Default()

	// Ping handler
	r.GET("/ping", func(c *ego.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
```

example for custom autocert manager.

[embedmd]:# (example/example2.go go)
```go
package main

import (
	"log"

	"github.com/go-ego/autotls"
	"github.com/go-ego/ego"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	r := ego.Default()

	// Ping handler
	r.GET("/ping", func(c *ego.Context) {
		c.String(200, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
```
