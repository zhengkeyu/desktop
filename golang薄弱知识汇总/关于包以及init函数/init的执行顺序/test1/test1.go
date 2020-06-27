package test1

import (
	"fmt"
)

var t1 = 100
var data = func() int {
	r := 0
	for i := 1; i <= 10; i++ {
		r += i
	}
	return r
}()

const ct2 = 99

func init() {
	fmt.Printf("var t1=%vv,data=%v,ct2=%v\n", t1, data, ct2)
	fmt.Println("test1 init1")
}
func init() {
	fmt.Println("test1 init2")
}
func Test1() {
	fmt.Println("test1 main")
}
