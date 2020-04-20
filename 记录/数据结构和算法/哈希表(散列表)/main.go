package main

import "fmt"

//哈希表结构
type hashTable struct {
	table [7]link
}

//链表
type link struct {
	data *user
}

//存放的数据机构
type user struct {
	id   int
	name string
	next *user
}
//哈希表也叫散列表，占用很少的内存来存储数据，避免了频繁的操作数据库，提高了性能
//把数据根据某种逻辑分散存储在不同的表中，使用时根据之前设定好的逻辑去找到对应的表提取数据，如果数据都存在一张表上，遍历需要耗费的资源大，这就是分表的好处
func main() {
	hashtable := newHashTable()
	//根据id分类添加到指定的表

	//会被添加到0号表
    hashtable.push(&user{id:7,name:"刘备"})
	hashtable.push(&user{id:14,name:"张飞"})
    //会被添加到1号表
	hashtable.push(&user{id:8,name:"曹操"})
	hashtable.push(&user{id:15,name:"张辽"})
    //会被添加到6号表
	hashtable.push(&user{id:13,name:"孙权"})
	hashtable.push(&user{id:20,name:"周瑜"})

	hashtable.printTable()
}

//新哈希表
func newHashTable() *hashTable {
	return &hashTable{table: [7]link{}}
}

//添加数据
func (this *hashTable) push(u *user) {
	id := insertWhere(u.id)
	if id <= 6 {
		if this.table[id].data == nil {
			this.table[id].data = u
		} else {
			for {
				if this.table[id].data.next == nil {
					this.table[id].data.next = u
					break
				}
				this.table[id].data.next = this.table[id].data.next.next
			}
		}
	}
}
//判断插入哪张表
func insertWhere(id int) int {
	return id % 7
}
//遍历哈希表
func (this *hashTable)printTable(){
	for k := range this.table{
		if this.table[k].data != nil{
			for  {
				fmt.Printf("所在表: %v\n",k)
				fmt.Printf("id: %v\n",this.table[k].data.id)
				fmt.Printf("name: %v\n",this.table[k].data.name)
				fmt.Println()
				if this.table[k].data.next == nil{
					break
				}else{
					this.table[k].data.next = this.table[k].data.next.next
				}
			}
		}
	}
}