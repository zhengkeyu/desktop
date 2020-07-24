package main

type YanZ struct {
	Name      string `binding:"required,min=1,max=5"`
	SmallName string `binding:"nefield=Name"`   //不能等于Name字段
	Ip        string `binding:"ipv4"`           //必须是ipv4
	Age       int    `binding:"omitempty,gt=5"` //可以不传,传的话必须大于5,反之写上lt=5传的就必须小于5
}

//required //不能不传或传零值
//min //最小长度,对字符串,slice,map //最小值,对数值型
//max //与上同理
//len //与上同理
//gt  //大于,与上同理
//lt  //小于,与上同理

type YanZAll struct {
	YZ []YanZ `binding:"dive"` //dive将递归验证切片YZ里的所有YanZ对象
}

//验证器文档
//https://godoc.org/github.com/go-playground/validator
func main() {

}
