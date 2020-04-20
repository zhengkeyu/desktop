package main

import "fmt"

//链表
type table struct {
	name string
	next *table
}
func main(){
	var obj = &table{"first",nil}
	obj.push("zhangsan")
	obj.push("lisi")
	obj.push("wangwu")
	obj.push("zhaoliu")

	obj.delete("first") //删
	obj.update("zhaoliu","liubei")//改
	obj.printTable()//查
}
//增
func (this *table)push(n string){
	for  {
		if this.next == nil{
			break
		}
		this = this.next
	}
	this.next = &table{name:n}
}
//删
func (this *table)delete(n string){
	for  {
		if this.next.name == n{
			this.next = this.next.next
			break
		}
		if this.next.next == nil{
			fmt.Println("没有这个信息")
			break
		}
		this = this.next
	}
}
//改
func (this *table)update(src,dst string){
	for  {
		if this.name == src{
			this.name = dst
			break
		}
		if this.next == nil{
			fmt.Println("没有这个信息")
		}
		this = this.next
	}
}
//查
func (this *table)printTable(){
	for  {
		fmt.Println(this.name)
		if this.next == nil{
			break
		}
		this = this.next
	}
}