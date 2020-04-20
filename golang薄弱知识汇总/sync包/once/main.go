package main

import (
	"fmt"
	"sync"
)

func main() {
	//one.Do()方法可以确保让函数只执行一次,可用于for循环内
	var one = &sync.Once{}
	for i := 1; i <= 5; i++ {
		one.Do(TestDo)
		fmt.Println("abc")
	}
}
func TestDo()  {
	fmt.Println("hello")
}
