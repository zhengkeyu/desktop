package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	var buf = os.Args
	if len(buf) != 3{
		fmt.Println("复制：必须指定源文件和目标文件")
	}
	if buf[1] == buf[2]{
		fmt.Println("源文件和目标文件不能同名")
	}
	file,err := os.Open(buf[1])
	if err != nil{
		fmt.Println(err)
	}
	var slic = make([]byte,1024*4)
	srcfile,err := os.Create(buf[2])
	for  {
		filebyte,err := file.Read(slic)
		if err != nil{
			if err == io.EOF{
				fmt.Println("文件读取完毕")
				break
			}
			fmt.Println(err)
		}
		if err != nil{
			fmt.Println(err)
		}
		_,err = srcfile.Write(slic[:filebyte])
		if err != nil{
			fmt.Println(err)
		}
	}
}