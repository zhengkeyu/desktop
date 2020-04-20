package main

import (
	"os"
	"fmt"
	"io"
	"log"
)

func main() {

	////文件信息
	//fileinfo,err := os.Stat("E:/郑克宇软件资料/培训日程/学习整理记录")//指定文件，返回对象以及错误
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(fileinfo.Name())//返回文件名字
	//	fmt.Println(fileinfo.IsDir())//返回是不是文件夹
	//	fmt.Println(fileinfo.Size())//返回文件大小
	//	fmt.Println(fileinfo.ModTime())//返回修改时间
	//	fmt.Println(fileinfo.Mode())//返回文件权限
	//}

	////文件路径操作
	//path,err := filepath.Abs("test.txt")//返回绝对路径
	//fmt.Println(path,err)
	//isabs := filepath.IsAbs("E:\\郑克宇软件资料\\培训日程\\学习整理记录")//返回路径是不是绝对路径
	//fmt.Println(isabs)

   // //创建文件夹
   //err := os.Mkdir("dir",os.ModePerm)//创建文件夹,os.ModePerm => 0777权限
	//if err != nil {
	//	fmt.Println(err)
	//}
   //err = os.MkdirAll("zheng/keyu",0777)//递归创建文件夹
	//if err !=nil {
	//	fmt.Println(err)
	//}

	////创建文件
	//file,err := os.Create("test.txt")//创建文件，有则覆盖,底层调用的openfile
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(file)
	//}

	//打开文件
	//file02,err:=os.Open("E:\\郑克宇软件资料\\培训日程\\学习整理记录")//打开文件，打开后只能用于读
    //fmt.Println(file02,err)

	//os.O_RDWR =>可读可写
	//os.O_CREATE =>存在则打开，不存在则创建再打开
	//os.O_RDONLY =>只能用于读
	//os.O_WRONLY =>只能用于写
	//os.O_APPEND =>追加模式，不会覆盖
	//file03,err := os.OpenFile("./test.txt",os.O_RDWR | os.O_CREATE,0777)
	//fmt.Println(file03,err)
	//file03.Close()//关闭文件，每次打开最后需关闭

	//删除文件
	//err := os.Remove("test.txt") //只能删除空文件夹或文件
	//err = os.RemoveAll("test.txt")//递归删除
	//fmt.Println(err)

	//读取文件
	//file,err := os.Open("test.txt")
	//if err !=nil {
	//	fmt.Println(err)
	//}
	//var sile = make([]byte,15)//接收读取到的信息
	//bytenum,err := file.Read(sile)//返回读取的字节数，以及可能出现的错误，再次读取会覆盖
	//if err != nil{
	//	fmt.Println(err)
	//}else {
	//	fmt.Println(string(sile),bytenum)
	//}

	//for循环读取一个文件
	mapfile,err := os.Open("E:\\郑克宇软件资料\\map.txt")
	defer mapfile.Close()//函数执行完需要关闭
	if err != nil {
		log.Panic(err)
	}
	var slic01 = make([]byte,1024)
	for ; ;  {
		byte,err := mapfile.Read(slic01)
		if byte ==0 || err == io.EOF  { //当byte==0 或者err==io.EOF时表示文件读取完毕了
			fmt.Println("文件读取完毕")
			break
		}
		fmt.Println(string(slic01[0:byte]))//读多少打印多少
	}

	//写入文件
	file,err := os.OpenFile("./test.txt",os.O_RDWR,os.ModePerm)
	if err != nil {
		log.Panic(err)
	}
	//file.WriteString("字符串")//直接往打开的对象里写入字符串
    byt,err := file.Write(slic01)//写入，参数为字节切片，可使用read()读取后再写入
	if err !=nil {
		log.Panic(err)
	}else{
		fmt.Printf("写入成功=>%d字节",byt)
	}
}
