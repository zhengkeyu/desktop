package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Basic Auth是开放平台的两种认证方式，简单点说明就是每次请求API时都提供用户的username和password

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	//以下的账户允许访问
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"user1": "abc123",
		"user2": "abc456",
		"user3": "abc789",
	}))

	//请求该接口时需登录
	//试着访问看看
	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secret": "The secret ingredient to the BBQ sauce is stiring it in an old whiskey barrel.",
		})
	})

	r.Run(":8081") // 监听服务在 0.0.0.0:8080

}
