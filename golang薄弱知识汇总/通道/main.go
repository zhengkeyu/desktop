package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个双向通道
	//var ch = make(chan<- int) //只能存的通道
	//var ch = make(<-chan int) //只能取的通道
	//var ch = make(chan int ,5)//有缓存的通道
	var ch = make(chan int)
	//取
	go get(ch)
	//存
	go push(ch)
	//主程序睡
	time.Sleep(time.Second)
}
//单向通道，只能取
func get(ch <-chan int)  {
	for k := range ch{ //fo range对通道的读取
		fmt.Println(k)
	}
}
//单向通道，只能存
func push(ch chan<- int)  {
	for i := 0;i < 5;i++ {
		ch <- i
	}
	close(ch)
}