package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/normal", Normal)
	router.GET("/handle", Handle)
	router.GET("/any", Any)
	router.Run() //默认端口8080
	//router.Run(":8088") //指定端口
}

// 由gin封装的方法,如：GET、POST、PUT、DELETE、PATCH、OPTION、HEAD
func Normal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "get",
	})
}

// curl -X GET "http://127.0.0.1:8080/normal"

func Handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "handle",
	})
}

//curl -X POST "http://127.0.0.1:8080/handle"

func Any(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": c.Request.Method,
	})
}

//curl -X DELETE "http://127.0.0.1:8080/any"
