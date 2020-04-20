package main

import "fmt"

func main() {
	var num = []interface{}{1,false,"abc",[]int{1,2,3},map[string]int{"a":1,"b":2}}
	var num02 interface{} = 1
	if value,ok := num02.(int);ok{//如果num是括号里面的类型，value就是num的值，ok就是true，否则ok
		fmt.Println(value)
		fmt.Println(ok)
	}else{
		fmt.Println(value)
		fmt.Println(ok)
	}
	for k,v := range num {
		switch v.(type) { //v.(type)返回v的数据类型
		case int:
			fmt.Printf("%d => %d\n",k,v)
		case string:
			fmt.Printf("%d => %s\n",k,v)
		case bool:
			fmt.Printf("%d => %t\n",k,v)
		case []int:
			fmt.Printf("%d => %v\n",k,v)
		default:
			fmt.Println("error")
		}
	}



}