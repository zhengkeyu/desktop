package main

import (
	"errors"
	"fmt"
)

func main() {
	TestPanic()
	fmt.Println("main run hear")
}
//在一个函数中调用painc程序会崩掉，但是如果调用recover程序就能恢复
func TestPanic(){
	defer func() {
		v := recover()
		//处理错误
		if err,ok := v.(error);ok{
			fmt.Println("err = ",err)
		}else {//可以再panic
			fmt.Println("panic again")
			panic(err)
		}
	}()
	panic(errors.New("this is an error"))
	//panic(123)
}
