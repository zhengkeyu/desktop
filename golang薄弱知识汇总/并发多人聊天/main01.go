package main

import (
	"net"
	"fmt"
	"io"
)

func main() {
	conn,err := net.Dial("tcp","localhost:6000")
    if err != nil{
    	fmt.Println(err)
	}
    go func() {
		for  {
			var serarray = make([]byte,1024)
			num,err := conn.Read(serarray)
			if err == io.EOF{
				fmt.Println(err)
				break
			}
			if err != nil{
				fmt.Println(err)
				break
			}
			fmt.Println(string(serarray[:num]))
		}
	}()
	var str string
	for  {
		fmt.Scan(&str)
		_,err := conn.Write([]byte(str))
		if err != nil{
			fmt.Println(err)
			break
		}
		if str == "exit"{
			break
		}
	}

}
