package main

import (
	"fmt"
)

func main() {
	var str = "zheng"
	//修改字符串里的字符

	//错误演示
	//str[0] = 'a'

	//正确演示
	by := []byte(str)
	by[0] = 'a'
	str = string(by)
	fmt.Println(str)

	//其他操作
	var str2 = "abcdhj"
	str2 = str2[:4]
	fmt.Println(str2)
}
