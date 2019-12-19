package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	println(router)
	// 重定向
	router.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://baidu.com")
		c.Abort()
	})

	// 只返回状态码
	router.GET("/code", func(c *gin.Context) {
		c.Status(http.StatusAccepted)
	})

	// string渲染
	router.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s", "world")
	})

	// json渲染
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Json rendering successful",
		})
	})

	// html渲染
	router.GET("/html", func(c *gin.Context) {
		//https://blog.csdn.net/a976134036/article/details/78867297
		c.HTML(http.StatusOK, "home", "xxx")
	})

	// xml渲染
	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "Xml rendering successful",
		})
	})
	router.Run()
}
