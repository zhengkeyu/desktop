package main

import (
	"fmt"
	"os"
)

//队列结构
type Queue struct {
	MaxMen int
	Head int
	tail int
	QueueArr []string
}
func main() {
	//新队列
	obj := NewQueue(5)
	//循环接受指令，根据键盘输入参数走对应方法
	for  {
		fmt.Println("输入push进队")
		fmt.Println("输入pull离队")
		fmt.Println("输入print遍历队伍")
		fmt.Println("输入exit退出程序")
		flag := ""
		data := ""
		fmt.Scan(&flag)
		switch flag {
		case "push":
			fmt.Println("输入name:")
			fmt.Scan(&data)
			obj.Push(data)
		case "pull":
			obj.Pull()
		case "print":
			obj.PrintArr()
		case "exit":
			fmt.Println("退出程序")
			os.Exit(1)
		default:
			fmt.Println("参数错误")
		}
	}
}
//创建新队伍
func NewQueue(num int)*Queue{
	return &Queue{
		MaxMen:num,
		Head:-1,
		tail:-1,
		QueueArr:make([]string,num),
	}
}
//进队
func (this *Queue)Push(name string){
	if this.tail == len(this.QueueArr)-1&&this.Head==0||this.tail-this.Head == -1{
		fmt.Println("人满了")
		return
	}

	if this.Head == -1&&this.tail==-1{
		this.QueueArr[0]=name
		this.Head = 0
		this.tail = 0
	}else if this.tail==len(this.QueueArr)-1{
		this.QueueArr[0] = name
		this.tail = 0
	}else{
		this.QueueArr[this.tail+1] = name
		this.tail++
	}
}
//离队
func (this *Queue)Pull(){
	//如果this.Head == -1&&this.tail==-1说明队伍没人
	if this.tail == -1&&this.Head==-1{
		fmt.Println("队伍没人")
		return
	}
	if this.Head == 0&&this.tail ==0{
		fmt.Printf("%s离队",this.QueueArr[this.tail])
		//如果队伍没人了必须把this.Head = -1，this.tail = -1重置为-1
		this.Head = -1
		this.tail = -1
	}else if this.tail > this.Head{
		fmt.Printf("%s离队",this.QueueArr[this.tail])
		this.tail --
	}else{
		fmt.Printf("%s离队",this.QueueArr[this.tail])
		this.tail++
	}
}
//遍历队伍
func (this *Queue)PrintArr(){
	if this.tail==-1&&this.Head==-1{
		fmt.Println("队伍没人")
		return
	}
	if this.tail >= this.Head{
		for i:=this.Head;i<=this.tail;i++  {
			fmt.Println(this.QueueArr[i])
		}
	}else{
		for i:=this.Head;i<len(this.QueueArr);i++  {
			fmt.Println(this.QueueArr[i])
		}
		for i:=0;i<=this.tail;i++  {
			fmt.Println(this.QueueArr[i])
		}
	}
}
