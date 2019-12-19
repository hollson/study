package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// gin基于https://github.com/julienschmidt/httprouter路由器
func main() {
	router1()
}

func TestHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

/// 1.gin启动服务
func router1() {
	router := gin.Default()
	router.GET("/", TestHandler)
	router.Run()
}

/// 2.由http托管服务
func router2() {
	host := gin.Default()
	host.GET("/", TestHandler)
	http.ListenAndServe(":8080", host)
}

/// 3.http.Server托管服务
func router3() {
	host := gin.Default()
	s := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	host.GET("/", TestHandler)
	s.ListenAndServe()
}
