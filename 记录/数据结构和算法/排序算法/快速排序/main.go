package main

import "fmt"

func main() {
	arr := []int{34, 11, 56, 2, 89, 5, 66, 15}
	Sort (0,len(arr)-1,arr)
	fmt.Println(arr)
}

//快速排序使用递归
func Sort(left int, right int, array []int) {
	l := left
	r := right
	privot := array[(left+right)/2]
	for ; l < r; {
		for ; array[l] < privot; {
			l++
		}
		for ; array[r] > privot; {
			r--
		}
		if l >= r {
			break
		}
		//交换
		array[l], array[r] = array[r], array[l]
		//优化
		if array[l] == privot {
			r--
		}
		if array[r] == privot {
			l++
		}
	}
	//如果l == r,在移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		Sort(left, r, array)
	}
	//向右递归
	if right > l {
		Sort(l, right, array)
	}
}
