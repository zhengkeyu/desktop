package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
)

func main() {
	////执行删除文件
	//cmd := exec.Command("cmd.exe", "/c", "del .\\golang薄弱知识汇总\\执行终端命令\\test")
	//err := cmd.Run()
	//if err != nil {
	//	panic(err)
	//}

	//执行查看文件
	cmd2 := exec.Command("cmd.exe", "/c", "dir")
	var result = new(WriteString)
	cmd2.Stdout = result //得到的结果将会写入到result

	err := cmd2.Run()
	if err != nil {
		panic(err)
	}

	//将GBK转成UTF-8,否则中文乱码
	utf8str, _ := simplifiedchinese.GBK.NewDecoder().String(string(*result))
	fmt.Println(utf8str)
}

type WriteString string

//WriteString实现了Write所以也就实现了Writer接口因此可以赋值给cmd2.Stdout
func (w *WriteString) Write(p []byte) (n int, err error) {
	*w += WriteString(p)
	return len(p), nil
}
