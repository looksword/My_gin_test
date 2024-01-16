package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个新的 Gin 引擎
	r := gin.Default()

	// 添加一个路由，当有人访问根路径时，会返回 "Hello, World!"
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// 运行 Gin 引擎
	r.Run()
}
