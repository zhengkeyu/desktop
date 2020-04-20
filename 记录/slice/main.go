package main

import "fmt"

//slice是对数组的引用
func main() {
	//引用一个对slice引用的slice的注意点
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("arr1 = %v\n", arr1)
	arr2 := arr1[2:6]
	fmt.Printf("arr2 = %v,len = %d,cap = %d\n", arr2, len(arr2), cap(arr2))
	//arr3是对arr2的引用，虽然arr2没有索引7，但是arr2是对arr1的引用,cap是9，所以7是可以取到的
	arr3 := arr2[3:7]
	fmt.Printf("arr3 = %v,len = %d,cap = %d\n", arr3, len(arr3), cap(arr3))

	//append使用中的注意点
	s1 := []int{1,2,3,4,5,6}
	s2 := s1[1:4]
	_ = append(s2,999)
	fmt.Println(s1)  //结果 = [1 2 3 4 999 6]
	//当对一个引用其他slice的slice使用append时，会对引用的源slice产生改变，因为切片不只有len属性还有cap属性
}
