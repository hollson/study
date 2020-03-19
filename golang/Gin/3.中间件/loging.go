package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

//gin.Default()下自带的两个中间件。
func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)

	router:=gin.Default()
	//router := gin.New()
	//router.Use(gin.Logger(), gin.Recovery())
	router.GET("/log", func(c *gin.Context) {
		name := c.DefaultQuery("name", "bob")
		panic("middleware test")
		c.String(200, "%s", name)
	})
	router.Run()
}
