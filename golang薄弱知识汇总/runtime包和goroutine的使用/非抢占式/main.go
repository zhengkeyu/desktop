package main

import (
	"fmt"
	"runtime"
	"time"
)
//非抢占式goroutine,是由goroutine来主动交出控制权，不受系统调度
func main() {
	runtime.GOMAXPROCS(1) //设置只使用一个cpu(核心数)
	var arr = [10]int{}
	for i := 0; i <= 10; i++ {
		if i == 0 {
			go func(i int) { //占用唯一的cpu，其他的协程没有执行的机会
				for {
					arr[i]++
				}
			}(i)
		} else {
			go func(i int) { //io或者打印...,会让出协程的控制权
				for {
					fmt.Println(i)
				}
			}(i)
		}
	}

	//唯一的cpu核心被占用，执行不到以下代码
	time.Sleep(time.Millisecond)
	fmt.Println("main end...")
}
