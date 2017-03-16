package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()
	r.POST("/v1/graph/", func(c *gin.Context) {
		var nodes map[string]string
		err := c.BindJSON(&nodes)
		if err != nil {
			c.String(500, "%s", err)
			return
		}

		f, err := os.Open("/Users/jimyeh/.go/src/github.com/lemonlatte/conn2graphviz/connections.gif")
		if err != nil {
			c.String(500, "%s", err)
			return
		}

		buf := bytes.Buffer{}

		encoder := base64.NewEncoder(base64.StdEncoding, &buf)

		_, err = io.Copy(encoder, f)
		if err != nil {
			c.String(500, "%s", err)
			return
		}
		c.String(200, buf.String())
	})
	r.Static("/", "ui/public")

	r.Run(":8080")
}
