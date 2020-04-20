package main

import "fmt"

//接口的运用
type Class interface {
	run(string) //只写方法名和参数类型
	eat()
}
//接口的继承
type Class1 interface {
	Class
}
//接口的实现类
type Person struct {
	Name string
	Age int
}

func (p Person)run(str string){
	fmt.Printf("%s在跑步\n",str)
}
func (p Person)eat(){
	fmt.Printf("%s在吃饭\n",p.Name)
}

func main() {
	//方法类型复制给变量，再调用
	//隐藏方法值
	var per1 = Person{"lisi",20}
	pfun1 := per1.eat
	pfun1()
	fmt.Printf("%T\n",pfun1)
	//方法表达式
	var per = Person{"zhangsan",15}
	pfun := (Person).run
	pfun(per,"zzz")
	fmt.Printf("%T\n",pfun)
}