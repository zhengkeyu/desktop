package main

import (
	"html/template"
	"log"
	"os"
)

type student struct {
	Name string
	Age  int
}

func main() {
	//创建模板名和模板 方式1
	//teml,err := template.New("test").Parse(`hello {{.}}`)
	//if err != nil{
	//	log.Panic(err)
	//}
	////往模板放数据
	//err = teml.Execute(os.Stdout,"this is test")
	//if err != nil{
	//	log.Panic(err)
	//}

	//创建模板名和模板 方式2
	//如果反射的是结构体，字段的首字母必须大写，否则找不到
	teml, err := template.New("test").Parse(
		`name => {{.Name}}
             age  => {{.Age}}
            `,
	)
	//也可以是指定一个模板文件
	//teml, err := template.New("test").ParseFiles("./file.txt")
	//if err != nil {
	//	log.Panic(err)
	//}
	//往模板放数据
	var obj = &student{
		Name: "小明",
		Age:  15,
	}
	//第一个参数为要输出到的目标
	err = teml.Execute(os.Stdout, obj)
	if err != nil {
		log.Panic(err)
	}
}
