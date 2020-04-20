package main

import (
	"strconv"
	"fmt"
)

func main() {
	var slic = make([]byte,0,1024)
	//把bool类型的数值转化为byte，并放进数组中
	slic = strconv.AppendBool(slic,true)
	//把int64类型的数值转化为byte，指定进制，并放进数组中
	slic = strconv.AppendInt(slic,1024,10)
	//把string类型的数值转化为byte，并放进数组中
	slic = strconv.AppendQuote(slic,"abc")
	fmt.Println(string(slic))

	//bool类型转化为string
	str := strconv.FormatBool(true)
	fmt.Println(str)
	//float类型转化成字符串
	str1 := strconv.FormatFloat(3.14,'f',-1,32)
	fmt.Println(str1)
	//strconv.Atoi("111")
	//strconv.Itoa(111)
	//把字符串转化为bool类型
	bo,_ := strconv.ParseBool("true")
	fmt.Println(bo)
}