package main

import (
	"time"
	"fmt"
)

func main() {
ch1 := make(chan int)
ch2 := make(chan int)
go func() {
	time.Sleep(5*time.Second)
	ch1 <- 100
}()

go func() {
	time.Sleep(5*time.Second)
	ch2 <-200
}()
	select {
	case data := <-ch1:
		fmt.Println("one",data)
	case data := <-ch2:
		fmt.Println("two",data)
}
}
