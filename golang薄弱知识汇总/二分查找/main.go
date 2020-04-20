package main

import "fmt"

func main() {
	var arr = []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(GetIndex(arr,0,9,1))
}
//取得开始和结束的索引，不断判断中间索引的值
func GetIndex(arr []int,left int,right int,value int)int{
	if len(arr) == 1{
		if arr[0] == value{
			return left
		}else{
			return -1
		}
	}
	if left >right{
		fmt.Println("起始索引不能小于结束索引")
		return -1
	}
	cen := (left+right)/2
	if value > arr[cen]{
		return GetIndex(arr,cen+1,right,value)
	}else if value < arr[cen]{
		return GetIndex(arr,left,cen -1,value)
	}else{
		return cen
	}
	//return -2
}