package main

//在go中变量的大写表示公共的，包括结构体里面的字段也要大写才能调用，可以在别的包调用
import (
	or "fmt"    //给包起别名
	. "strconv" //省略包名
	_ "sync"    //声明包不使用，但是会执行包里面的init函数
)

//init()是比较特殊的函数，他会在自己程序内运行，也会在自己的包被调用时执行
func init() {
	or.Println("这是一个导包自动执行的初始化函数")
}
func main() {
	or.Println("a")
	Atoi("122")
}
