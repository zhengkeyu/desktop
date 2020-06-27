package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

type Book struct {
	Name  string
	Price int
}
type Info struct {
	Address     string
	PhoneNumber string
	Age         int
}
type TestObj struct {
	Name       string
	PersonInfo *Info
	Books      []Book
	Joins      map[string]string
	Xss        string
	XssSafe    string
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		t := template.New("test.html")

		//模板函数的使用
		f := func(b []Book) int { //定义
			return len(b)
		}
		t.Funcs(template.FuncMap{ //添加
			"myf": f,
			"safe": func(s string) template.HTML {
				return template.HTML(s)
			},
		})

		//修改模板引擎标识符 {{}}
		//t.Delims("{[","]}") //将语法{{}}改为{[]}

		_, err := t.ParseFiles("./golang薄弱知识汇总/模板引擎/模板语法/test.html")
		if err != nil {
			panic(err)
		}

		err = t.Execute(writer, createTestObj())
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func createTestObj() TestObj {
	rand.Seed(time.Now().UnixNano())
	r := TestObj{
		Name: "郑克宇",
		Books: []Book{
			{"三国演义", 300},
			{"水浒传", 400},
			{"红楼梦", 250},
			{"红楼梦", 380},
			{"Golang入门到精通", 1000},
		},
		PersonInfo: &Info{
			Address:     "厦门",
			PhoneNumber: "13950******",
			Age:         rand.Intn(35) + 15,
		},
		Joins: map[string]string{
			"Golang协会": "会长",
			"中国历史协会":   "初级会员",
			"户外旅行社团":   "中级会员",
		},
		//html/template对比text/template 多了转义的功能，防止xss攻击，此script不会执行
		Xss:     "<script>alert(123);</script>",
		XssSafe: "<a href=\"https://www.baidu.com\">百度一下</a>", //模板中使用safe函数可防止转义
	}
	return r
}
