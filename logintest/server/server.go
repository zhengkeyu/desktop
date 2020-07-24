package server

import "github.com/gin-gonic/gin"

func RunGin() {
	r := gin.New()
	r.POST("/login", UserLogin)           //登录
	r.POST("/refreshtoken", RefreshToken) //刷新token
	r.POST("/register", UserRegister)     //注册

	r.Run(":8081")
}
