package cal_test

import (
	"fmt"
	"testing"
)
//实现不用main函数也可以测试函数,可以有多个测试用例函数
//测试用例，文件名必须以_test.go结尾，测试用例函数必须以Test+首字母大写开头
//cd到该文件夹下运行go test(运行错误时才输出日志) 或go test -v(-v代表运行正确错误都输出日志),默认测试所有_test.go结尾文件。
//也可以指定测试测试哪个文件go test -v 文件名(如果该测试有引用其他文件的话需要加上该文件名?.go)或再进一步确定指定测试那个函数go test -v -test.run 函数名
func TestCal(t *testing.T){
	_,err := fmt.Println("hello")
	if err != nil{
		//格式化输出错误信息，并退出程序
		t.Fatalf("fmt.Println()运行错误%v",err)
	}else{
		t.Logf("输出日志")
	}
}

func TestTwo(t *testing.T){
	_,err := fmt.Print("hello")
	if err != nil{
		//格式化输出错误信息，并退出程序
		t.Fatalf("fmt.Print()运行错误%v",err)
	}else{
		t.Logf("输出日志")
	}
}