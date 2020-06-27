package main

//执行init的顺序
//1:按导包顺序从第一条往最后一条
//2:包里面按照文件首字母a-z
//3:文件里按照从上到下
//注：执行init函数前会先执行文件内的全局变量
import (
	"desktop/golang薄弱知识汇总/关于包以及init函数/init的执行顺序/test1"
	"desktop/golang薄弱知识汇总/关于包以及init函数/init的执行顺序/test2"
)

//这种导入方法可以选择导入顺序
//import "Desktop/golang薄弱知识汇总/关于包以及init函数/init的执行顺序/test1"
//import "Desktop/golang薄弱知识汇总/关于包以及init函数/init的执行顺序/test2"

func main() {
	test1.Test1()
	test2.Test2()
}
