package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type TestCard struct {
	Val int
}

//自定义类型
type TestCards []TestCard

func main() {
	num1 := Cardinit(10)
	num := sort.Reverse(num1) //此方法返回和参数相反的排序(升序或降序)
	sort.IsSorted(num)        //查询是否已经排序
	fmt.Println("排序前:", num)
	sort.Sort(num) //排序
	fmt.Println("排序后:", num)
}
func Cardinit(c int) TestCards {
	var res = make(TestCards, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < c; i++ {
		res = append(res, TestCard{rand.Intn(500)})
	}
	return res
}

//将需要排序的数据类型实现以下三个方法，再调用sort.Sort()就可实现排序

//获取长度
func (this TestCards) Len() int {
	return len(this)
}

//比大小
func (this TestCards) Less(i1, i2 int) bool {
	return this[i1].Val < this[i2].Val //把<换成>会按从大到小排序
}

//交换位置
func (this TestCards) Swap(i1, i2 int) {
	this[i1], this[i2] = this[i2], this[i1]
}
