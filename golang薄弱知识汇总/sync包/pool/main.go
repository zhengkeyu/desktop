package main

import (
	"bytes"
	"sync"
)

type testPool struct {
	Id int
	Name string
	Arr []int
}
//临时对象池会重用内存地址而非开辟新空间，因此可以减少内存的使用率并可增加效率
func main() {
	////实例化对象池
	//myPool := sync.Pool{
	//	New: func() interface{} {
	//		return new(testPool)//指定返回的数据类型
	//	},
	//}
	////获取对象池里的对象没有则会new一个
	//obj1 := myPool.Get().(*testPool)//获得的对象需要类型断言
	////给字段赋值
	//obj1.Name = "zheng"
	//obj1.Id = 1
	//obj1.Arr = []int{1,2,3}
	////对象不用需放回对象池中
	//myPool.Put(obj1)


	//对象池对缓存的应用
	var bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	buffer := bufferPool.Get().(*bytes.Buffer)

	buffer.Reset()//bytes.Buffer不用需要重置
	bufferPool.Put(buffer)
}
