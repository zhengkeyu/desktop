package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   string `json:"name"`
	Fun    func()
	Funtwo func(int, string) (int, string)
}

func TestPrint() {
	fmt.Println("这是一个测试函数")
}
func TestValuefun(num int, str string) (int, string) {
	return num + 1, str + str
}
func (this Person) Test(num1, num2 int) int {
	return num1 + num2
}
func (this Person) ATest(num1, num2 int) int {
	return num1 * num2
}
func main() {
	//获取Type对象
	rtype := reflect.TypeOf(Person{})
	//获取Value对象,如果要获取结构体的信息或者修改结构体的值必须传入指针
	rval := reflect.ValueOf(new(Person))
	//获取结构体的信息或者修改结构体的值都必须用Elem()方法先获取到指针
	//获取结构体字段数
	count := rval.Elem().NumField()
	fmt.Printf("这个结构体有%d个字段", count)
	//获取结构体的字段
	name := rval.Elem().FieldByName("Name")
	//获取到第二个字段
	fun := rval.Elem().Field(1)
	funtwo := rval.Elem().FieldByName("Funtwo")
	//赋予或修改结构体字段,参数必须先转成Value对象或者直接调用对应数据类型的方法
	name.SetString("张三")
	fmt.Println(name)
	fun.Set(reflect.ValueOf(TestPrint))
	fmt.Println(fun)
	funtwo.Set(reflect.ValueOf(TestValuefun))
	fmt.Println(funtwo)
	//调用字段里的函数
	//转成interface{}并断言，然后调用
	fun.Interface().(func())()
	i, s := funtwo.Interface().(func(int, string) (int, string))(9, "hello")
	fmt.Printf("这是结构体字段函数的返回值%v,%v\n", i, s)
	//调用结构体的方法
	var data []reflect.Value
	data = append(data, reflect.ValueOf(10))
	data = append(data, reflect.ValueOf(40))
	//0代表调用的第一个方法，方法是按照方法名的首字母的阿斯克码表大小排序
	//方法返回的返回值会存在[]reflect.Value里
	res := rval.Method(0).Call(data)
	fmt.Printf("这是方法调用的返回值%v\n", res[0])
	//也可以直接知道方法名
	var data1 []reflect.Value
	data1 = append(data1, reflect.ValueOf(10))
	data1 = append(data1, reflect.ValueOf(20))
	res1 := rval.MethodByName("Test").Call(data1)
	fmt.Printf("这是方法调用的返回值%v\n", res1[0])
	//获取结构体标签,也可以用FieldByName()方法获取，但要注意此方法返回2个参数
	res_js := rtype.Field(0).Tag.Get("json")
	fmt.Println(res_js)
}

//反射实现遍历结构体
func PrintStruct(d interface{}, fun func(k string, v reflect.Value)) {
	t, v := reflect.TypeOf(d), reflect.ValueOf(d)
	if t.Kind() != reflect.Struct { //如果不是结构体不处理
		return
	} else if t.Kind() == reflect.Ptr { //如果是指针
		t, v = t.Elem(), v.Elem() //取出指针的值
	}
	for i := 0; i < v.NumField(); i++ {
		k, v := t.Field(i), v.Field(i)
		fun(k.Name, v) //
	}
}
