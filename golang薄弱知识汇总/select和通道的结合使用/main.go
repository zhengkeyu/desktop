package main

import (
	"fmt"
	"time"
)

func main() {
	//timer延时功能
	//value := time.NewTimer(time.Second)//2秒后忘value.C通道里存
	//value.Reset(time.Second) //重新设置timer秒数
	//value.Stop() //结束定时器
	//fmt.Println(<-value.C)

	//after延时功能
	//var ti = <-time.After(time.Second)
	//fmt.Println(ti)

	//ticker间隔执行
	//tick := time.NewTicker(time.Second)
	//for  {
	//	fmt.Println(<-tick.C)
	//}

	var ch01 = make(chan int)
	var ch02 = make(chan int)
	go get(ch02)
	go push(ch01)
	//用于通道的读取判断，哪个case里的通道发生读取则执行此case
	//注:
	//如果有default其他的case没有io会走default
	//如果没有default并且其他的case没有io,select会阻断
	select {
	case res01 := <-ch01: //select即使管道没有关闭，也不会一直阻塞，会往下走
		fmt.Println("有人放入了: ", res01)
	case ch02 <- 1:
		fmt.Println("有人取走了: ", 1)
	case <-time.After(time.Second): //延迟2秒执行，2秒前阻塞
		fmt.Println("没有符合的")
	default:
		fmt.Println("都取不到走这里")
	}
	time.Sleep(time.Second * 2)

}
func push(ch chan int) {
	ch <- 2
}
func get(ch chan int) {
	<-ch
}

//select超时
func TimeOut() {
	var ch make(chan int)
	go func() {
		time.Sleep(time.Second * 10)
		ch <- 100
	}()
w:
	for {
		select {
		case r := <-ch:
			fmt.Println("读取数据: ", r)
		case <-time.After(time.Second * 5)://5秒超时，退出
			fmt.Println("timeout break")
			break w
		}
	}
	fmt.Println("next...")
}
func TimeOut2(){
	var ch make(chan int)
	go func() {
		time.Sleep(time.Second * 10)
		ch <- 100
	}()
	timer := time.NewTimer(time.Second*5)
w:
	for {
		select {
		case r := <-ch:
			fmt.Println("读取数据: ", r)
		case <-timer://5秒超时，退出
			fmt.Println("timeout break")
			break w
		}
	}
	fmt.Println("next...")
}