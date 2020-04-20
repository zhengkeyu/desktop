package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4}
	var point []*int
	for _, v := range arr {//for range 其实每次得到的值都是存在同一个内存地址
		point = append(point, &v)//所以这里其实append的都是同一个内存地址
	}
	for _, l := range point {
		fmt.Println(*l)
	}
}
