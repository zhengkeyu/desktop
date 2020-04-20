package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main() {
	//复合类型有2种 => complex64,complex128
	var com complex64 = 2.1+9i //声明一个复合型
	fmt.Println(real(com),imag(com))
	fmt.Printf("%T,%T\n",real(com),imag(com))//拆开都为float类型

    //字符类型
	var res = '1'
	fmt.Printf("%T\n",res)//字符类型是int32

	//goto的使用
	//goto可以使用在任何地方，但是不能跨函数使用
    init:
	fmt.Println("main start")
	var num int
	fmt.Scan(&num)
	fmt.Println("main running")
	if num == 0{
		fmt.Println("main init")
		goto init
	}
	fmt.Println("main end")

	//获取命令行的输入 => go run main.go a b c
	count := os.Args//返回命令行输入的[]string
	fmt.Printf("%T",count) //[]string

	//随机数
	rand.Seed(time.Now().UnixNano())//用当前时间的纳秒做种子
	fmt.Println(rand.Intn(10))//[0,10)整数
	fmt.Println(rand.Float64())//[0,10)浮点数
	fmt.Println(rand.Int())//int64长度的随机数

	//切片和底层数组的关系
	var slic = [6]int{1,2,3,4,5,6}
	slic1 := slic[1:3]
	fmt.Println(slic1)
	slic2 := slic1[1:5]
	fmt.Println(slic2)

	//copy函数的使用
	//不会对切片的长度做出任何变化
	var slic = []int{1,3,5}
	var slic1 = []int{2,4,6,8,10}
	copy(slic,slic1)
	fmt.Println(slic,slic1)//复制切片，slic1 => 源切片，slic => 目标切片

	//以下2种声明方式，没办法添加键值
	//var m1 map[int]string
	//m1[1] = "aaa" //报错
	//var slic0011 []int
	//slic0011[0] = 1 //报错


}
