package main

import "fmt"

func main() {
	arr := []int{34, 12, 6, 23, 66, 2, 99}//12,34,
	for i := 1; i < len(arr); i++ {
		insertIndex := i - 1
		insertVal := arr[i]
		for {
			//如果比到前面都没有比他大的数，说明这个数就是最大的
			if insertIndex == -1 {
				arr[0] = insertVal
				break
			}
			if insertVal > arr[insertIndex] {
				arr[insertIndex+1] = insertVal
				break
			}
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
	}
	fmt.Println(arr)
}
