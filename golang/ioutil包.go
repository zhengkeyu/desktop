package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func main() {
	//此方法简单，但是没有循序读取，文件太大，内存不够会卡死，不建议用
	byte,err := ioutil.ReadFile("./test.txt")
	if err !=nil{
		fmt.Println(err)
	}
	err = ioutil.WriteFile("./keyu.txt",byte,os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	//slic,err := ioutil.ReadDir("./BLC")//返回目标文件夹下的所有文件信息
}
