package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	//监听信号，当程序关闭时(手动关闭)会发送信号到通道s,否则阻断
	s := make(chan os.Signal)
	signal.Notify(s,os.Interrupt)
	d := <- s
	fmt.Println(d)
}
