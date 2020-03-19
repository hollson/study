package main

import (
	"github.com/gin-gonic/gin"
	locale_en "github.com/go-playground/locales/en"
	locale_zh "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
	trans_en "gopkg.in/go-playground/validator.v9/translations/en"
	trans_zh "gopkg.in/go-playground/validator.v9/translations/zh"
)


//binding换成validate
type User struct {
	Name     string `form:"name" validate:"required"`
	Password string `form:"password" validate:"required,min=6,max=12"`
}

var (
	Uni      *ut.UniversalTranslator //翻译器
	Validate *validator.Validate     //验证器
)

//验证多语言化
func main() {
	Validate = validator.New()
	zh := locale_zh.New()
	en := locale_en.New()
	Uni = ut.New(zh, en)
	router := gin.Default()
	router.GET("/trans", func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := Uni.GetTranslator("locale")
		switch locale {
		case "zh":
			trans_zh.RegisterDefaultTranslations(Validate, trans)
		default:
			trans_en.RegisterDefaultTranslations(Validate, trans)
		}
		var u User
		if err := c.ShouldBind(&u); err != nil {
			c.String(http.StatusInternalServerError, "%v", err)
			c.Abort()
			return
		}
		if err := Validate.Struct(u); err != nil {
			errs := err.(validator.ValidationErrors)
			sliceErrs := []string{}
			for _, e := range errs {
				sliceErrs = append(sliceErrs, e.Translate(trans))
			}
			c.String(http.StatusBadRequest, "%v", sliceErrs)
			c.Abort()
			return
		}
	})
	router.Run()
}

// curl -X GET "http://127.0.0.1:8080/trans?name=123&password=123"