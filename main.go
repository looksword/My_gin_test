package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个新的 Gin 引擎
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, looksword!")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 Not Found")
	})

	// 运行 Gin 引擎
	r.Run(":9393")
}
