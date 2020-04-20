package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	file,_ := os.Open("./test.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	file1,_ := os.Create("./demo.txt")
    writer := bufio.NewWriter(file1)
	defer file1.Close()
	for  {
		//每次读取一行
		//by,is,err := reader.ReadLine()
		//读取到换行结束
		str,err := reader.ReadString('\n')
		writer.Write([]byte(str))
		if err == io.EOF{
			break
		}
	}
    //writer只是写到内存中，调用此方法才能真正的写入到文件中
	writer.Flush()

}
