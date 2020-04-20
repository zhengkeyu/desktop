package main

import (
	"fmt"
	"unsafe"
)

type Test1 struct {
	Id   int
	Name string
}
type Test2 struct {
	Id   int
	Name string
}

func main() {
	var obj1 = Test1{Id: 1, Name: "zhangsan"}
	var obj2 = *(*Test2)(unsafe.Pointer(&obj1))
	fmt.Printf("obj1 = %T,obj2 = %T", obj1, obj2)
}
