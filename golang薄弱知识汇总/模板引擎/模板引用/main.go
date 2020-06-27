package main

import (
	"html/template"
	"os"
)

func main() {
    //模板可以引用其他模板，并可将渲染的数据传递给引用的模板
	//以下3个模板文件存在引用关系

	//读取模板文件，多个
	t, err := template.ParseFiles("./golang薄弱知识汇总/模板引擎/模板引用/static/index.html",
		"./golang薄弱知识汇总/模板引擎/模板引用/static/data1.html",
		"./golang薄弱知识汇总/模板引擎/模板引用/static/data2.html")
	if err != nil {
		panic(err)
	}

	//模板数据
	d := struct {
		Header string
		Data1  string
		Data2  string
	}{"zhengkeyu", "Golang", "厦门"}
	err = t.ExecuteTemplate(os.Stdout, "T", d)
	//“T”是指定渲染的模板
	//在没有define “T” 的时候可以用文件名 但是一旦define “T”了就必须使用"T"来找到模板
	if err != nil {
		panic(err)
	}
}
