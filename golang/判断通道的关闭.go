package main

import "fmt"

func main() {
var cha =make(chan int)
go input(cha)//放
 output(cha)//取
}
func input(c chan int){
	for i:=0;i<=50;i++  {
		fmt.Println("放入：",i)
		c <- i
	}
close(c)//通道的关闭，最后必须关闭
}
func output(c chan int)  {
	for   {
		if value,ok := <-c;!ok {//当通道是关闭时，ok==false,value==数据的初始化值。当通道是打开时，ok==true,value==通道里的值。
			break
		}else {
			fmt.Println("取出：",value)
		}
	}
}