package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	g := r.Group("g1")

	g.GET("/g1get", func(c *gin.Context) { //访问g1/g1get1
		c.String(http.StatusOK, c.Request.URL.Path)
	})

	//以上的接口不会有这个中间键，以下的接口会有
	g.Use(func(c *gin.Context) {
		fmt.Println("这是一个中间键")
	})

	g.GET("/g1get2", func(c *gin.Context) { //访问g1/g1get1
        fmt.Println("luy")
		c.String(http.StatusOK, c.Request.URL.Path)
	})

	//多个handlefunc会逐个执行，结果会拼起来
    r.GET("/simple", func(c *gin.Context) {
    	fmt.Println("handlefunc run first")
		c.String(http.StatusOK,"hello")
	}, func(c *gin.Context) {
		fmt.Println("handlefunc run second")
		c.String(http.StatusOK,"world")
	})


	r.Run(":8081")
}
