package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/auth/sign", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    "123",
			Path:     "/",
			HttpOnly: true,
		}
		// 向响应的头部写入：Set-Cookie: session_id=123; Path=/; HttpOnly
		http.SetCookie(c.Writer, cookie)
		c.String(http.StatusOK, "Login successful")
	})

	//router.Use(BackList())
	router.GET("/home",BackList(), AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	router.Run()
}

// 黑名单
func BackList() gin.HandlerFunc {
	return func(c *gin.Context) {
		backList := []string{"127.0.0.1"}
		curIp := c.ClientIP()

		for _, v := range backList {
			if v == curIp {
				c.String(http.StatusForbidden, "you have been blacklisted")
				c.Abort()
			}
		}
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}

// 浏览器中依次访问：
//http://127.0.0.1:8080/auth/sign
//http://127.0.0.1:8080/home

// 或
//curl -X GET "http://127.0.0.1:8080/home" --cookie "session_id=123"
