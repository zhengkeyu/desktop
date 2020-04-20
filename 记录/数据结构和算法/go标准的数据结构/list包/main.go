package main

import (
	"container/list"
	"fmt"
)

//链表数据结构
//type List struct {
//	root Element // 链表的根元素
//	len  int     // 链表的⻓度
//}

//type Element struct {
//	next, prev *Element    // 上⼀个元素和下⼀个元素
//	list       *List       // 元素所在链表
//	Value      interface{} // 元素
//}
func main() {
	list := list.New() //初始化链表
	fmt.Println(list)

	//对应的方法

	//type Element
	//func (e *Element) Next() *Element
	//func (e *Element) Prev() *Element

	//type List
	//func New() *List
	//func (l *List) Back() *Element // 最后⼀个元素
	//func (l *List) Front() *Element // 第⼀个元素
	//func (l *List) Init() *List // 链表初始化
	//func (l *List) InsertAfter(v interface{}, mark *Element) *Element // 在某个元素前插⼊
	//func (l *List) InsertBefore(v interface{}, mark *Element) *Element // 在某个元素后插⼊
	//func (l *List) Len() int // 在链表⻓度

	//func (l *List) MoveAfter(e, mark *Element) // 把e元素移动到mark之后
	//func (l *List) MoveBefore(e, mark *Element) // 把e元素移动到mark之前
	//func (l *List) MoveToBack(e *Element) // 把e元素移动到队列最后
	//func (l *List) MoveToFront(e *Element) // 把e元素移动到队列最头部
	//func (l *List) PushBack(v interface{}) *Element // 在队列最后插⼊元素
	//func (l *List) PushBackList(other *List) // 在队列最后插⼊接上新队列
	//func (l *List) PushFront(v interface{}) *Element // 在队列头部插⼊元素
	//func (l *List) PushFrontList(other *List) // 在队列头部插⼊接上新队列
	//func (l *List) Remove(e *Element) interface{} // 删除某个元素

}
