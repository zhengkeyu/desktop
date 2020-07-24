package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Format("2006-01-02 15:04:05") //返回当前时间的字符串
	fmt.Println(t)

	//根据时间戳转化成时间字符串
	time.Unix(1595814960,0).Format("2006-01-02 15:04:05")

	//时间字符串转成时间戳
	loc,_ := time.LoadLocation("Asia/Shanghai")
	t2,_ := time.ParseInLocation("2006-01-02 15:04:05","2020-07-20 13:35:00",loc)
	fmt.Println(t,t2)
}
