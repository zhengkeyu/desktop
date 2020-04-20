package main

import (
	"fmt"
	"runtime"
)

func main() {
	//指定cpu以多少核来执行协程
	res := runtime.GOMAXPROCS(8)
	fmt.Println(res)
	go func() {
		for i := 0;i <= 5;i++ {
			fmt.Println("goroutine: ",i)
		}
	}()
	go func() {
		for i := 0;i <= 5;i++ {
			if i == 4{
				//结束此协程
				runtime.Goexit()
			}
			fmt.Println("goroutine2: ",i)
		}
	}()

	for i := 0;i <= 10;i++ {
		//让此协程让出执行权,并随机再分配
		runtime.Gosched()
		fmt.Println("main: ",i)
	}
}