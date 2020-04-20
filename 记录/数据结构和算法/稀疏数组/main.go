package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	res := SparseArithmetic()

	//生成原数组模板
	arr := make([][]int,res[0][0])
	for i := 0;i<res[0][0];i++ {
		for j := 0;j< res[0][1];j++ {
			arr[i] = append(arr[i],0)
		}
	}

	//根据稀疏数组还原原数组
	for k,v := range res {
		if k == 0{
			continue
		}
		arr[v[0]][v[1]] = v[2]
	}
	fmt.Println(arr)
}

//创建原数组并生成稀疏数组返回
func SparseArithmetic()[][3]int{
	//创造一个稀疏数组
	arr := make([][]int,10)
	for i:=0;i<10;i++{
		for j := 0;j<10;j++ {
			arr[i] = append(arr[i],0)
		}
	}
	rand.Seed(time.Now().UnixNano())
	for k :=range arr{
		for i := 0;i<=rand.Intn(5);i++ {
			r := rand.Intn(len(arr[0]))
			arr[k][r] = rand.Intn(100)
		}
	}
	arr[0][0] = 10
	fmt.Println(arr)

	//对稀疏数组进行统计
	sparse := make([][3]int,0)
	count := 0
	for k,v := range arr {
		for j,l := range v {
			if l != 0{
				res := [3]int{k,j,l}
				sparse = append(sparse,res)
				count++
			}
		}
	}
	final := make([][3]int,0)
	ar := [3]int{len(arr),len(arr[0]),count}
	final = append(final,ar)
	final = append(final,sparse...)
	fmt.Println(final)
	return final
}