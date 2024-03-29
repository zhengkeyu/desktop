package main

import "fmt"

func main() {
	arr := []int{23,12,3,56,11,78,21}
	ShellSort(arr)
	fmt.Println(arr)
}
func ShellSort(nums []int) []int{
	//外层步长控制
	for step := len(nums) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(nums); i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
	}
	return nums
}