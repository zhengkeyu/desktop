package main

import (
	"fmt"
	"os"
)

type user struct {
	Name string
	Age  int
}

func main() {
	//获取键盘输入,%c相当于指明输入的参数是个字符类型
	//var num1 int
	//fmt.Scanf("%c",&num1)
	//fmt.Println(num1)

	file, _ := os.Create("./test.txt")
	defer file.Close()
	//可以往一个io.Writer里写入格式化的数据
	fmt.Fprintf(file, "这是要格式化写入的数据%v", "hello")
	//占位符对打印结果位数的限制
	fmt.Printf("%06d", 100)  //000100
	fmt.Printf("%.3f", 3.14) //3.140
	//指定打印占位符
	fmt.Printf("%*s", 10, "hello") //10指定的是占位大小,结果 => "     hello"
	//对结构体字段名的打印
	o := user{
		Name: "张三",
		Age:  18,
	}
	fmt.Printf("%+v", o)
}
