package main

import "fmt"

func main() {
	//选择排序类似冒泡排序的优化版本
	//打乱的切片
	arr := []int{5, 1, 23, 44, 14, 17, 89, 45}
	for i := 0; i < len(arr)-1; i++ {
		max := arr[i]
		maxIndex := i
		for k := i+1; k < len(arr); k++ {
			if max > arr[k] {
				max = arr[k]
				maxIndex = k
			}
		}
		//如果没进上面的if就不用交换了
		if maxIndex != i{
			arr[i],arr[maxIndex] = arr[maxIndex],arr[i]
		}
	}
	fmt.Println(arr)
}
