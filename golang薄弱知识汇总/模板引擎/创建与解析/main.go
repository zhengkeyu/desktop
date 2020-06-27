package main

import (
	"html/template"
	"os"
)

func main() {
	////创建方式1
	//t := template.New("t3.html") //名字t.html必须在以下的ParseFiles路径中能找到
	//_, err := t.ParseFiles("./golang薄弱知识汇总/模板引擎/创建与解析/static/t3.html")
	////t.Parse() //针对字符串的解析
	//if err != nil {
	//	panic(err)
	//}
	//err = t.Execute(os.Stdout, "test value")
	//if err != nil {
	//	panic(err)
	//}

	//创建方式2
	t, err := template.ParseFiles("./golang薄弱知识汇总/模板引擎/创建与解析/static/t.html",
		"./golang薄弱知识汇总/模板引擎/创建与解析/static/t2.html",
	)
	//t, err := template.ParseGlob("./golang薄弱知识汇总/模板引擎/创建与解析/**/*") //解析路径下的所以模板
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(os.Stdout, "t.html", struct {
		//解析多个模板
		BookName string
		Text     string
	}{
		BookName: "三国演义",
		Text:     "东汉末年，诸侯割据，天下大乱，分久必合，合久必分!!!",
	})
	if err != nil {
		panic(err)
	}
}
