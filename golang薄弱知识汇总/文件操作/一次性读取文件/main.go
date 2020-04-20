package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	by,_ := ioutil.ReadFile("./test.txt")
	err := ioutil.WriteFile("./demo.txt",by,0666)
	if err != nil{
		fmt.Println(err)
	}
}
