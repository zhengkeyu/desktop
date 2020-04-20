package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Data []string `json:"data"`  //因为结构体字段首字母大写和底下的json字段不符，所以需要转
	Ishave bool `json:"ishave"`  //👆
	Name string `json:"name"`    //👆
	Sex string `json:"sex"`      //👆
}
type maptest struct {
	Name string
	Data []interface{}
	Ishave bool
	Sex string
}
func main() {
	var js string = `{
	"data": [
		"语文",
		"数学",
		"英语"
	],
	"ishave": true,
	"name": "zhangsan",
	"sex": "女"
}`
    //json转为结构体
	//var struc01 Student
	//err := json.Unmarshal([]byte(js),&struc01)//这里一点要传指针
	//if err != nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Println(struc01)
	//}
    //json转为map
    var ma01 = make(map[string]interface{},4)
    err := json.Unmarshal([]byte(js),&ma01)
    if err != nil{
    	fmt.Println(err)
	}else{
		fmt.Println(ma01)
	}
	//类型断言
	var res maptest
	for k,v := range ma01 {
		switch t := v.(type) {//t := v.(type)写法只能用于switch，t等于类型的值，case匹配类型
		case string:
			if k == "name"{
				res.Name = t
			}
			if k == "sex"{
				res.Sex = t
			}
		case []interface{}: //注意：json类型里的切片转化成map会变成[]interface{}
			res.Data = t
		case bool:
			res.Ishave = t
		default:
			fmt.Println("都不是")
		}
	}
    fmt.Println(res)

}
