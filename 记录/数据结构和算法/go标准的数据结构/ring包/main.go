package main

import (
	"container/ring"
	"fmt"
)

var count = 5

func main() {
	ring := ring.New(count)
	fmt.Println(ring)
	//ring的方法
	//func New(n int) *Ring // 初始化环
	//func (r *Ring) Do(f func(interface{})) // 循环环进⾏操作
	//func (r *Ring) Len() int // 环⻓度
	//func (r *Ring) Link(s *Ring) *Ring // 连接两个环
	//func (r *Ring) Move(n int) *Ring // 指针从当前元素开始向后移动或者向前（n可以为负数）
	//func (r *Ring) Next() *Ring // 当前元素的下个元素
	//func (r *Ring) Prev() *Ring // 当前元素的上个元素
	//func (r *Ring) Unlink(n int) *Ring // 从当前元素开始， 删除n个元素
}
