package main

import "fmt"

func main() {
arr := []int{31,22,4,12,67,14,56}
sort(arr,0,len(arr)-1)
fmt.Println(arr)
}
func sort(arr []int,begin, end int){
	b := begin
	e := end
	center := (end - begin)/2
	if b >= e{
		return
	}
	for b < e {
		for arr[b] < arr[center] {
			b++
		}
		for arr[e] > arr[center] {
			e--
		}
		arr[e],arr[b] = arr[b],arr[e]
	}
	sort(arr,begin,b)
	sort(arr,e,end)
}