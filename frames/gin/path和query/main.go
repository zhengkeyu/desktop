package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// 会匹配/user/? 或 /user/?/，但不会匹配 /?/ 或 /?
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, ":name = %v", name)
	})

	// 会匹配/user/?/ 或 /user/?/send,不会匹配/user/?
	// 如果路径是/user/?/会被上面的优先匹配到,因为先声明
	//If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, ":name = %v,*action = %v", name, action)
	})

	router.GET("/query", func(c *gin.Context) {
		name := c.Query("name") //获取query没有则返回空字符串
		//name := c.DefaultQuery("name","defaultName") //获取query没有则给默认值
		c.String(http.StatusOK, name)
	})

	router.POST("form/json", func(c *gin.Context) {
		////以下只针对表单的
		//c.PostForm("name")
		//c.DefaultPostForm("name")

		////绑定到结构体可以是json或者form
		//d := TestForm{}
		//err :=c.ShouldBind(&d)
	})

	router.POST("fileform", func(c *gin.Context) {
		f, _ := c.FormFile("file")                 //获取上传文件
		_ = c.SaveUploadedFile(f, "./"+f.Filename) //保存上传文件至硬盘

		//多个文件上传
		//fs,_ := c.MultipartForm()
		//for _,v := range fs.File{
		//
		//}
	})

	router.Run(":8081")
}

type TestForm struct {
	Id   int
	Name string
	Arr  []int
}
