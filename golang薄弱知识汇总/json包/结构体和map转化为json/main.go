package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`  //如果需要改变生成的json字段名字，可用这种方法改动
	Data []string `json:"-"`   //此字段不输出
	Ishave bool `json:",string"` //将此字段的值以string输出
	Sex string
}
func main() {
	//利用结构体生成json格式数据
	var stu =  Student{"zhangsan",[]string{"aaa","数学","英语"},true,"男"}
	//js,_ := json.Marshal(stu)
	js,_ := json.MarshalIndent(stu,"","	")//" " => tab,格式化编码，与上面的函数相比，较明了，方便阅读
	fmt.Println(string(js)) //返回的是[]byte,转成string后打印输出

	//利用map生成json格式数据
	var ma01 = make(map[string]interface{},4)
	//"name":"zhangsan","data":[]string{"语文","数学","英语"},"ishave":true,"sex":"女"
	ma01["name"] = "zhangsan"
	ma01["data"] = []string{"语文","数学","英语"}
	ma01["ishave"] = true
	ma01["sex"] = "女"
	//js01,_ := json.Marshal(ma01)
	js01,_ := json.MarshalIndent(ma01,"","	")
	fmt.Println(string(js01))

}