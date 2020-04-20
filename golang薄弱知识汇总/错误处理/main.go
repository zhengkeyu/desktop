package main

import (
	"fmt"
	"log"
)

func main() {
	//创建错误对象的2种方法
	//err := errors.New("除数不能为0")
	//err := fmt.Errorf("%s","除数不能为0")
	test(2,0)
	test1()
}
//Panic和recover的使用
func test(a,b int)(int){
	defer func (){
		recover()//恢复程序，只能放在有defer的函数中使用，有错误则返回错误信息，没有则返回nil
	}()
	fmt.Println("in test()")
	if b == 0{
		log.Panic("除数不能为0")//直接报错中断主函数
		return -1
	}else{
		return a/b
	}
}
func test1(){
	fmt.Println("this is a function")
}