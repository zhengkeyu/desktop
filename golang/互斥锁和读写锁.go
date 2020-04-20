package main

import (
	"sync"
	"fmt"
)

var count = 100
var wg = new(sync.WaitGroup)
var rw = new(sync.RWMutex)
func main() {
	//互斥锁
	var lock = new(sync.Mutex)
	lock.Lock()//锁
	lock.Unlock()//解锁
	//读写锁
	//读写互斥,读读不互斥,写写互斥。  互斥=>(只允许一个线程操作), 不互斥=>(可以多个线程操作)
	var rw = sync.RWMutex{}
	rw.RLock()//读锁定
	rw.RUnlock()//读解锁
	rw.Lock()//写锁定
	rw.Unlock()//写解锁
	wg.Add(1)
    go gettick("客户1")
	go gettick("客户2")
	go gettick("客户3")
	go viewtick()
	go viewtick()
	wg.Wait()
}

func gettick(who string){
	for  {
		if count==0{
			wg.Add(-1)
		}
		rw.Lock()
		count--
		fmt.Printf("%s抢票1张,余票%d\n",who,count)
		rw.Unlock()
	}
}
func viewtick(){
	for  {
		if count==0 {
			fmt.Println("查询票已经卖完了")
			break
		}
		rw.RLock()
		fmt.Printf("余票%d\n",count)
		rw.RUnlock()
	}
}

