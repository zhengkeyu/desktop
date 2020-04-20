package main

import (
	"time"
	"fmt"
)

func main() {
	timer := time.NewTimer(time.Second*2)//返回Timer对象
	res :=  <-timer.C //timer.C属性是个通道，里面存有当前时间
	fmt.Println(res)

	//cha := time.After(time.Second*2)//指定时间，把当前时间放到通道里返回
	//res := <-cha //取出通道里的值
	//fmt.Println(res)
}
