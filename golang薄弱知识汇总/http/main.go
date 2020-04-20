package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//底层库实现的类似beego的路由控制器
type myHandle struct {
}

func (*myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("router one"))
	if err != nil {
		log.Panic(err)
	}
}

type myHandleTwo struct {
}

func (this *myHandleTwo) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//如果超过1秒没有响应，服务器就认为响应超时，不再响应
	time.Sleep(time.Second)
	res.Write([]byte("router two"))
}

//最简单的路由控制器
func saybye(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("simple router"))
	if err != nil {
		log.Panic(err)
	}
}
func main() {
	//发送一个http请求
	//cli := &http.Client{}
	//req,err :=  http.NewRequest("GET","https://tieba.baidu.com/f?ie=utf-8&kw=golang",nil)
	//if err != nil{
	//	log.Panic(err)
	//}
	//res,err := cli.Do(req)
	//if err != nil{
	//	log.Panic(err)
	//}
	//by,_ := ioutil.ReadAll(res.Body)
	//file,_ := os.Create("./file.txt")
	//defer file.Close()
	//fmt.Println(string(by))

	//最简单的路由控制器
	//直接放匿名函数
	//http.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
	//	res.Write([]byte("success"))
	//})
	//放一个函数变量
	//http.HandleFunc("/", saybye)
	//http.ListenAndServe(":4000", nil)

	//底层库实现的类似beego的路由控制器
	var mux = http.NewServeMux()
	//根路由"/"包含了所有未注册的路由
	mux.Handle("/", &myHandle{})
	mux.Handle("/test", &myHandleTwo{})
	//http.ListenAndServe(":4000",mux)

	//也可以自己创建serve对象
	var serve = &http.Server{
		//设定监听ip，端口
		Addr: ":5000",
		//设定执行的对象
		Handler: mux,
		//设定服务器响应超时
		WriteTimeout: time.Second * 2,
	}
	//利用通道来主动退出服务器
	var ch = make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		err := serve.Close()
		if err != nil {
			fmt.Println("关闭服务器失败")
		}
	}()
	err := serve.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			fmt.Println("服务器主动关闭执行...")
		} else {
			fmt.Println("服务器其他原因关闭...")
		}
	}
	fmt.Println("服务器关闭成功")
}
