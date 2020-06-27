package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/sgyy", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("./golang薄弱知识汇总/模板引擎/模板继承/static/base.html",
			"./golang薄弱知识汇总/模板引擎/模板继承/static/h1.html")
		if err != nil {
			panic(err)
		}
		err = t.ExecuteTemplate(writer, "h1.html", nil)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/shz", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("./golang薄弱知识汇总/模板引擎/模板继承/static/base.html",
			"./golang薄弱知识汇总/模板引擎/模板继承/static/h2.html")
		if err != nil {
			panic(err)
		}
		err = t.ExecuteTemplate(writer, "h2.html", nil)
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
