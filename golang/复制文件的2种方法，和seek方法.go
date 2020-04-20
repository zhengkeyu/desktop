package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	var onepath = "E:\\郑克宇软件资料\\map.txt"
	var twopath = "./test.txt"
	//err := copyfile(onepath,twopath)
	//if err != nil{
	//	log.Panic(err)
	//}
	err := copyfiletwo(onepath,twopath)
	if err !=nil {
		fmt.Println(err)
	}
//设置光标位置的方法
file,_ := os.Open("./test.txt")
file.Seek(3,0)
//whence属性:
//0相对于文件开头偏移
//1相对于当前光标位置偏移
//2相对于文件结尾偏移
//offset =>偏移量
}
func copyfile(p, a string )error  {
	mapfile,err := os.Open(p)
	if err != nil {
		return err
	}
	file,err := os.OpenFile(a,os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		return err
	}
	var slic01 = make([]byte,10)
	for ; ;  {
		byte,err := mapfile.Read(slic01)//读取文件
		if byte ==0 || err == io.EOF  { //读取文件，当byte==0 或者err==io.EOF时表示文件读取完毕了
			fmt.Println("文件读取完毕")
			break
		}else if err !=nil {
			return err
		}
		fmt.Println("读取文件=>",string(slic01[0:byte]))//读多少打印多少
		byt,err := file.Write(slic01[0:byte])//写入，读多少写多少
		if err !=nil {
			return err
		}else{
			fmt.Printf("写入文件=>%d字节\n",byt)
		}
	}
	//文件打开需关闭
	mapfile.Close()
	file.Close()
	return nil
}
func copyfiletwo(a,b string)error{
	file,err := os.Open(a)
	if err !=nil {
		fmt.Println("源文件打开失败")
		return err
	}
	defer file.Close()
	file02,err := os.OpenFile(b,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err !=nil {
		fmt.Println("目标文件打开失败")
		return err
	}
	defer file02.Close()
	byte,err := io.Copy(file02,file)//io包下的方法，可以直接复制文件
	if err !=nil{
		fmt.Println("复制失败")
		return err
	}else {
		fmt.Printf("复制成功%d字节",byte)
	}
	return nil
}