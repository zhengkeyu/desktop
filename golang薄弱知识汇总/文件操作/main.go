package main

import (
	"os"
	"fmt"
)

func main() {
	//创建一个文件，如果文件存在则会清空文件内容，其实内部调用的是openfile()
	file,err := os.Create("./zky.txt")
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(file)
	}
	//以只读打开文件
	//file1,err := os.Open("./zke.txt")
	//os.OpenFile() //功能强大，参数较复杂
	//关闭文件，打开文件最后必须关闭
	file.Close()

	//写入
	//往file里写入信息，返回写入的字节数与错误
	num,err := file.Write([]byte("abc"))
	//往file里写入信息(参数只能字符串类型)，返回写入的字节数与错误
	num1,err := file.WriteString("abc")
	//往file里写入信息，2表示指定从哪个位置开始写,返回写入的字节数与错误
	num2,err := file.WriteAt([]byte("abc"),2)

	//读取
	//把读取的内容存到slic字节切片里面，返回读取的字节数和错误
	var slic = make([]byte,0,1024)
	num3,err := file.Read(slic)
	//把读取的内容存到slic字节切片里面，2表示从哪里开始读，返回读取的字节数和错误
	num4,err := file.ReadAt(slic,2)

	//删除文件
	os.Remove("./zky.txt")

	//标准设备文件的使用
	os.Stdout.Close()//默认打开的，关闭后，fmt.println()将无法在屏幕输出内容
	os.Stdout.WriteString("abc")//这个方法可以往屏幕写入内容，类似fmt.println()，且os.Stdout.Close()关闭后一样能写入
	//判断文件
	fileinfo,err := os.Stat("../demo")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(fileinfo.IsDir())//返回对象文件是不是文件夹
		fmt.Println(fileinfo.Name())//返回文件名字
		fmt.Println(fileinfo.Mode())//返回文件权限
		fmt.Println(fileinfo.Size())//返回文件大小
		//等等的属性
	}
}