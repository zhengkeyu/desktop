package main

import "fmt"

const day = 70

var arr [day + 1]int

//缓存实现斐波那契数列
func main() {
	fmt.Println(FeiBo(day))
}
func FeiBo(i int) int {
	if arr[i] != 0 {
		return arr[i]
	}
	if i <= 2 {
		arr[i] = 1
		return 1
	}
	res := FeiBo(i-1) + FeiBo(i-2)
	arr[i] = res
	return res
}
