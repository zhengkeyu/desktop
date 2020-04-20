package main

import (
	"fmt"
	"sort"
)

func main() {
	//针对[]int的排序
	var arr = []int{32,11,5,78,22}
	sort.Ints(arr)//此函数字段将[]int转为内建类型并排序
	fmt.Println(arr)//排好序了
	fmt.Println(sort.SearchInts(arr,111))//针对[]int的查找

	//针对[]float64的排序(跟上面的一样)
	var floArr = []float64{99.1,22.4,3.14,5.88}
	sort.Float64s(floArr)
	fmt.Println(floArr)
	fmt.Println(sort.SearchFloat64s(floArr,3.14))
	//如果是要降序
	var floArrTwo = []float64{99.1,22.4,3.14,5.88}
	sort.Sort(sort.Reverse(sort.Float64Slice(floArrTwo)))//如果是要降序需调用sort.Reverse(),并需要强转类型sort.Float64Slice()
	fmt.Println(floArrTwo)

	//针对[]string的排序
	var strArr = []string{"d1","c1","a1","b1"}
	sort.Strings(strArr)
	fmt.Println(strArr)
}
