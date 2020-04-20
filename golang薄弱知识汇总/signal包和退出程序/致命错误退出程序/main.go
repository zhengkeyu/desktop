package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fmt.Println("run recover")
		err := recover()//只能捕捉并恢复panic()的错误
		fmt.Println(err)
	}()
	func() {
		log.Fatal("致命错误,不可recover...") //打印错误信息并调用os.Exit(1)
		//os.Exit(1)退出程序
	}()
	fmt.Println("main run next...")
}
