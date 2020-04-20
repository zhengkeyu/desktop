package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	file,err := os.Open("./test.txt")
	if err !=nil{
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)//创建一个reader对象
	a,b,c := reader.ReadLine() //读取一次缓存
	fmt.Println(a,b,c)
	a,b,c = reader.ReadLine()
	fmt.Println(a,b,c)
	str,err := reader.ReadString('/')//读取到字符'/'位置，返回字符串
	fmt.Println(str,err)
	byt,err := reader.ReadBytes('=')//读取到字符'='位置，返回[]byte
	fmt.Println(string(byt),err)
}
