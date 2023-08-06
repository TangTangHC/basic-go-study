package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello go")
	})
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "这是参数路由, %s", name)
	})
	server.GET("/order", func(ctx *gin.Context) {
		oId := ctx.Query("id")
		ctx.String(http.StatusOK, "这是查询参数, %s", oId)
	})
	server.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
