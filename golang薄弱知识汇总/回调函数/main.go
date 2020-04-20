package main

import "fmt"

//定义一种函数类型
type testfunc func(int,int)int

//定义2种函数
func Add(a,b int)int{
	return a+b
}

func Mul(a,b int)int{
	if a >b{
		return a-b
	}else if a == b{
		return 0
	}else{
		return b -a
	}
}
//定义一个回调函数,回调函数根据所给的函数实现不同的功能，多态的体现
func NumberResult(a,b int,fun testfunc )int{
	return fun(a,b)
}

func main() {
res := NumberResult(10,2,Mul)
fmt.Println(res)
}