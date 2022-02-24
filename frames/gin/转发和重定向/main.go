package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	//转发
	//转发是服务费的内部跳转，不会产出新的请求
	r.GET("/forward", func(c *gin.Context) {
		c.Request.URL.Path = "/forward/text"
		r.HandleContext(c)
	})
	r.GET("/forward/text", func(c *gin.Context) {
		c.String(http.StatusOK,"<h1>这是一个标题。</h1>" )
	})

	//重定向
	//重定向是通知浏览器再发生一次请求
	r.GET("/redirect", func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path)
		c.Redirect(http.StatusMovedPermanently,"http://www.youtoball.com/app/try/go"+c.Request.URL.Path)

		fmt.Println("next...")
	})

	r.Run(":8081")
}
