package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()

	//第一次请求后如果有css或者js文件会再向服务端进行再请求
	g.Static("/static","./frames/gin/模板/static") //设置静态文件目录

	//设置模板函数g.SetFuncMap()对应 template.Funcs()

	g.LoadHTMLFiles("./frames/gin/模板/h.html") //模板解析

	g.GET("/index", func(c *gin.Context) { //模板渲染,并响应
		c.HTML(http.StatusOK, "h.html", nil)
	})

	_ = g.Run() //启动服务，默认8080
}
