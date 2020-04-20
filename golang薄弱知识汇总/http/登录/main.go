package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/login",Login)
	http.ListenAndServe(":5000",nil)

}
func Login(res http.ResponseWriter,req *http.Request){
	req.ParseForm()
	if req.Method == "GET"{//如果是get请求说明不是提交账号密码，所以响应登录页面
		t,_ := template.ParseFiles("golang薄弱知识汇总/http/登录/login.html")
		t.Execute(res,nil)
	}else{
		fmt.Println("username => ",req.Form["username"])
		fmt.Println("password => ",req.Form["password"])
	}
}