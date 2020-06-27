package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//context用于一个主协程的退出，控制其他衍生的子协程也退出
	ctx, cancel := context.WithCancel(context.Background())
    //调用cancel()时<-ctx.Done()都会收到信号来控制退出
	go woker(ctx, 1)

	go func(c context.Context) {
		go woker(c, 2)
		go woker(c, 3)
	}(ctx)

	time.Sleep(time.Second * 6)
	cancel()
	time.Sleep(time.Second * 10)
	fmt.Println("main quit")
	context.WithDeadline()
	context.WithTimeout()
}
//其他方法，类似context.WithCancel()，也是用于控制协程退出

//t := time.Now().Add(time.Second*3)
//ctx,_ := context.WithDeadline(context.Background(),t)   //指定一个具体的时间调用cancel()

//context.WithTimeout(context.Background(),time.Second*3) //指定过多少时间调用cancel()

func woker(c context.Context, name int) {
l:
	for i := 1; ; i++ {
		select {
		case <-c.Done():
			break l
		default:
			fmt.Printf("woker%v,num%v\n", name, i)
			time.Sleep(time.Second)
		}
	}
	fmt.Printf("woker%v,quit\n", name)
}
