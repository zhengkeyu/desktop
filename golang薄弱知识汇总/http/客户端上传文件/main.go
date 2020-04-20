package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		if request.Method == "GET" {
			hash := md5.New()                                              //加密算法
			io.WriteString(hash, strconv.FormatInt(time.Now().Unix(), 10)) //把当前时间参与哈希的计算
			token := fmt.Sprintf("%x", hash.Sum(nil))                      //生成最终的哈希
			t, _ := template.ParseFiles("./golang薄弱知识汇总/http/上传文件/file.html")
			t.Execute(writer, token)
		} else {
			request.ParseMultipartForm(32 << 20)                 //设置读取文件的最大缓存
			file, handler, err := request.FormFile("uploadfile") //获取文件句柄
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Fprint(writer, "%v", handler.Header) //响应写入
			f, _ := os.Create("./golang薄弱知识汇总/http/上传文件/save.rar")
			defer f.Close()
			io.Copy(f, file) //写入
		}
	})
	http.ListenAndServe(":5000", nil)

}
