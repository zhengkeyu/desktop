package main

import (
	"fmt"
	"time"
)

func main() {
	//设定时间和执行函数
	t := time.AfterFunc(time.Second*2, func() {
		fmt.Println("执行...")
	})
	//t.Reset(time.Second*5)//重新设置时间
	//t.Stop() //停止计时器
	time.Sleep(time.Second * 10)
}
