package main

import "fmt"

type TestInterface interface {
	Run()
}
type TestStruct struct {
}

func (this *TestStruct) Run() {
	fmt.Println("func run")
}
func NewTest() *TestStruct {
	return nil
}
func main() {
	//结构体空指针赋值给对应的实现接口，将不再为nil
	var v2 TestInterface
	v1 := NewTest()
	fmt.Printf("v1 = %v,IsNil = %v,type = %T\n",v1,v1 == nil,v1)
	v2 = NewTest()
	fmt.Printf("v2 = %v,IsNil = %v,type = %T\n",v2,v2 == nil,v2)
}