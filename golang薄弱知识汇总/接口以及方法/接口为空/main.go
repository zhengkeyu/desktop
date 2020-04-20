package main

import "fmt"

type TestInter interface {
	Fun()
}
type TeestStruct struct {
}

func (t *TeestStruct) Fun() {
}
func main() {
	//接口其实包含2个指针，一个指向值，一个指向对应的类型，只有2个都为nil，接口才为nil
	var testInter TestInter
	fmt.Printf("值:%v,类型:%T,是否为nil:%v\n",testInter,testInter,testInter==nil)
	testInter = TestFun(0) //这边把对应的值赋值为nil，类型赋值为*main.TeestStruct，所以接口现在不为nil
	fmt.Printf("值:%v,类型:%T,是否为nil:%v\n",testInter,testInter,testInter==nil)
}
func TestFun(v int) *TeestStruct {
	if v == 0 {
		return nil
	} else {
		return &TeestStruct{}
	}
}
