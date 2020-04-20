package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Format("2006-01-02 15:04:05") //返回当前时间的字符串
	fmt.Println(t)
}
