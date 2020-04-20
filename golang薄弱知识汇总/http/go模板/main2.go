package main

import "fmt"

type Student struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	//服务器返回一个模板页面
	//http.HandleFunc("/index", func(res http.ResponseWriter, req *http.Request) {
	//	tem, err := template.ParseFiles("./http/static/html/index.html")
	//	if err != nil {
	//		fmt.Fprintf(res,"parsefile error: %v",err)
	//	}
	//	var obj = Student{"zheng",27,"男"}
	//	err = tem.Execute(res,&obj)
	//	if err != nil {
	//		fmt.Fprintf(res,"execute error: %v",err)
	//	}
	//})
	//http.ListenAndServe(":7000",nil)
	//var ch = make(chan map[int]int)
	//	//var m = make(map[int]int)
	//	//for i:=1;i<=200;i++{
	//	//	go Test(i,ch)
	//	//}
	//	//for val := range ch{
	//	//	for k,v := range val{
	//	//		m[k] = v
	//	//	}
	//	//}
	//	//fmt.Println(m)
	var arr = []int{1, 2, 3, 4, 5, 6}
	res := DeleteArrayValue(arr, 5)
	fmt.Println(res)
}

func DeleteArrayValue(arr []int, index int) []int {
	var res = []int{}
	res = append(res, arr[:index]...)
	res = append(res, arr[index+1:]...)
	return res
}
