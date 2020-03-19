package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"reflect"
	"time"
)

//https://github.com/go-playground/validator
type Account struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=12"`
	Confirm  string `json:"confirm" binding:"eqfield=Password"` //字段名，不是标签名
	Age      int    `json:"age" binding:"required,gt=0"`
}

func main() {
	router := gin.Default()
	defer router.Run()
	router.POST("/invalid", TagInvalid)

	// 注册自定义验证方法
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	router.GET("/bookable", CustomInvalid)
}

// 1. 标签验证
func TagInvalid(c *gin.Context) {
	acc := new(Account)
	if err := c.BindJSON(acc); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
	}
	data, _ := json.Marshal(acc)
	c.String(http.StatusOK, string(data))
}

//curl -X POST "http://127.0.0.1:8080/invalid" -d '{"name":"tom","password":"123456","confirm":"123456","age":22}'

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}

// 2.自定义验证
func CustomInvalid(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

