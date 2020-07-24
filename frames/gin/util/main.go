package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	r := gin.New()

	//(1)设置请求时输出日志
	r.Use(gin.Logger())
	//或者可以设置自定义日志
    //r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//	// your custom format
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))

	//(2)日志文件和输出
	f, _ := os.Create("./frames/gin/util/gin.log")
	gin.DefaultErrorWriter = io.MultiWriter(f, os.Stdout) //设置os.Stdout同时会输出到控制台

	r.GET("/log", func(c *gin.Context) {

		//输出日志
		fmt.Fprintln(gin.DefaultErrorWriter, c.Request.URL.String())

	})

	r.Run(":8081")
}
