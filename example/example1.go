package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-ego/autotls"
	"github.com/go-ego/ego"
)

func main() {
	r := ego.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
