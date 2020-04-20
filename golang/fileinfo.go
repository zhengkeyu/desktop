package main

import (
	"os"
	"fmt"
)

func main() {
	fileinfo,err := os.Stat("E:/郑克宇软件资料/培训日程/学习整理记录")//指定文件，返回对象以及错误
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(fileinfo.Name())//返回文件名字
		fmt.Println(fileinfo.IsDir())//返回是不是文件夹
		fmt.Println(fileinfo.Size())//返回文件大小
		fmt.Println(fileinfo.ModTime())//返回修改时间
		fmt.Println(fileinfo.Mode())//返回文件权限
	}
}
