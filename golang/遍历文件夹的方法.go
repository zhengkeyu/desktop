package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	var path = "E:\\郑克宇软件资料\\培训日程"
	rangedir(path)
}
func rangedir(path string)  {
	slicfileinfo,err := ioutil.ReadDir(path)//遍历路径下的所有文件，放到切片里面 => []Fileinfo
	if err != nil {
		fmt.Println(err)
	}
	for _,r := range slicfileinfo{
		if r.IsDir() {//判断是否是文件夹，如是就再递归调用rangedir()
			fmt.Println("==============================")
			fmt.Println(r.Name())
			rangedir(path+"\\"+r.Name())
		}else {
			fmt.Println(r.Name())
		}
	}
}