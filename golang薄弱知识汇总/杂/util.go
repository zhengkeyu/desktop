package main

import "fmt"

func main() {
	//返回变量占用的字节大小
	//var num int8 = 10
	//fmt.Println(unsafe.Sizeof(num))

	//字符串的拼接换行
	//应该以+号结尾换行
	//var str = "a"+"b"+
	//	"c"
	//fmt.Println(str)

	//string类型其实底层是byte数组,需要改变字符串的单个字符的值需要把字符串转成byte数组
	var str = "abc"
	arr := []byte(str)
	arr[0] = 65
	str = string(arr)
	fmt.Println(str)
}
