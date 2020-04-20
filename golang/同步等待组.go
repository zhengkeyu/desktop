package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg  = &sync.WaitGroup{}//声明一个指针类型的同步等待组
	wg.Add(2)//等待2个线程
	go print1(wg)
	go print2(wg)
	wg.Wait()//阻断，除非Add()==0
	fmt.Println("end")
}
func print1(wg *sync.WaitGroup){
	for i:=1;i<=100;i++  {
		fmt.Println(i)
	}
	wg.Done()//Add(2)-1,也可写成Add(-1)
}
func print2(wg *sync.WaitGroup){
	for i:=1;i<=100;i++  {
		fmt.Println(i)
	}
	wg.Add(-1)
}
