package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name  string `form:"name"`
	Birth string `form:"birth",time_format:"2006-01-02"`
}

func main() {
	router := gin.Default()
	router.Any("/param", UrlParamHandler)
	router.Any("/obj", BindObjHandler)
	router.Any("/body", BodyHandler)
	router.Any("/bodyrebind", BodyRebindHandler)
	router.Run()
}

// 1.url参数
func UrlParamHandler(c *gin.Context) {
	name := c.Query("name")
	sex := c.DefaultQuery("sex", "male")
	c.String(http.StatusOK, "name=%s,sex=%s", name, sex)
}

//curl -X  GET "http://127.0.0.1:8080/param?name=hollson&sex=male"

// 2.绑定对象，json
func BindObjHandler(c *gin.Context) {
	p := &Person{}
	if err := c.ShouldBind(p); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"name":  p.Name,
		"birth": p.Birth,
	})
}

//curl -X GET "http://127.0.0.1:8080/obj?name=hollson&birth=1985-12-30"
//curl -X POST "http://127.0.0.1:8080/obj" -d "name=hollson&birth=1985-12-30"
//curl -H "Content-Type:application/json" -X POST "http://127.0.0.1:8080/obj" -d '{"name":"hollson","birth":"1988-12-01"}'

// 3.获取Body
func BodyHandler(c *gin.Context) {
	body := c.Request.Body
	bts, err := ioutil.ReadAll(body)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
	}
	//name := c.PostForm("name")  //已经拿不到了
	c.String(http.StatusOK, string(bts))
}

// curl -X POST "http://127.0.0.1:8080/body" -d '{"name":"tom","birth":"1988-12-12"}'
// curl -H "Content-Type:application/json" -X POST "http://127.0.0.1:8080/body" -d '{"name":"hollson","birth":"1988-12-01"}'

// 4.Body二次绑定
func BodyRebindHandler(c *gin.Context) {
	body := c.Request.Body
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
	}
	//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	name := c.PostForm("name")
	birth := c.DefaultPostForm("birth", "2000-01-01")
	c.String(http.StatusOK, "%s，%s，%s", string(bodyBytes), name, birth)
}

// curl -X POST "http://127.0.0.1:8080/bodyrebind" -d 'name="lucy"'
// curl -X POST "http://127.0.0.1:8080/bodyrebind" -d 'name="Bob"&birth=2019-01-01'
