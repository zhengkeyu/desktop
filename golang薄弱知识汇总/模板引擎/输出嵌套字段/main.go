package main

import (
	"html/template"
	"os"
)

type Student struct {
	Id     int
	Name   string
	Book   []string
	Friend map[int]Student
}

func main() {
	var stu = &Student{
		Id:   1,
		Name: "张三",
		Book: []string{"三国演义", "水浒传"},
		Friend: map[int]Student{1:
		{
			Id:     2,
			Name:   "赵六",
			Book:   []string{"红楼梦"},
			Friend: map[int]Student{},
		},
		},
	}
	t, _ := template.New("new tem").Parse(`
Id: {{.Id}}
Name:{{.Name}}{{range .Book}}
Book: {{.}}{{end}}{{with .Friend}}{{range .}}
{{.}}{{end}}{{end}}`)
	t.Execute(os.Stdout, stu)
}
