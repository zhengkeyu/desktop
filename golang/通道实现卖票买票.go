package main

import (
	"time"
	"fmt"
)
func main() {
	var ch01 = make(chan int,5)
	var ch02 = make(chan bool)
	go bytick("客户1",ch01,ch02)
	go bytick("客户2",ch01,ch02)
	go bytick("客户3",ch01,ch02)
	saletick(12,ch01)
	ch02 <-false //几个客户几个阻塞，所有客户都执行完才会解除阻塞
	ch02 <-false
	ch02 <-false
}
//卖票
func saletick(tic int,ch chan int){
	for i:=1;i<=tic;i++  {
		ch <- i
	}
	close(ch)
}

//取票方法1：
//func bytick(who string,ch chan int,ch2 chan bool){
//	for value := range ch {
//		fmt.Printf("%s买票,第%d张\n",who,value)
//		time.Sleep(time.Second)
//	}
//	<- ch2
//}

//取票方法2：
func bytick(who string,ch chan int,ch2 chan bool){
	for {
		res := <-ch
		if res==0 {
			break
		}
		time.Sleep(time.Second)//睡一下保证取出的值都能打印出来
		fmt.Printf("%s买票,第%d张\n",who,res)
	}
	<- ch2
}