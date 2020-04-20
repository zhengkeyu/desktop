package main

import (
	"fmt"
	"sort"
)

func main() {
	//Search()方法可以在一个已经升序排序的切片中查找对应条件的元素，有则返回元素的索引，没有则返回不存在的索引
	var arr = []int{1,2,3}
	c := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 11 //自己实现的条件
	})
	fmt.Println(c)
}
