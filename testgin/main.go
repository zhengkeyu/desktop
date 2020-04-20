package main

import (
	"desktop/testgin/src"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1/topic") //路由分组
	v1.GET("", src.GetTopicList)
	v1.GET("/:flag", src.GetTopicDetail)

	v1.Use(src.MustLogin()) //中间键，判断是否登录
	v1.POST("", src.NewTopic)
	v1.DELETE("/topic_id", src.DelTopic)

	//注册验证器函数
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("fun", src.Fun)
	}
	//
	err := router.Run() //默认127.0.0.1:8080
	if err != nil {
		panic(err)
	}

	//和router.Run()效果差不多
	//server := http.Server{
	//	Addr:              ":8080",
	//	Handler:           router,
	//}
	//server.ListenAndServe()
}
