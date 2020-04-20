package main

import (
	"runtime"
	"fmt"
)

func init()  {//init()函数类似main函数，但是每个包可以有一个，起到初始化作用
	cpu := runtime.NumCPU() //获取cpu线程数
	fmt.Println(cpu)
	runtime.GOMAXPROCS(runtime.NumCPU())//设置运行的最大线程
}
func main() {

}
