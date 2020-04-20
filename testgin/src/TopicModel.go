package src

import (
	"github.com/go-playground/validator/v10"
)

type Topic struct {
	TopicId    int `json:"topic_id"`
	TopicTitle string
}
type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page"`
	Age      int    `form:"age" binding:"required,fun"` //使用自定义验证器函数Fun
	//binding标签代表query必须传这字段，否则c.BindQuery(&qs)会有错
	PageSize int `json:"pagesize" form:"pagesize" binding:"required"`
}

//验证器实例,也可以自定义验证器(需实现一个函数，并传函数名到struct tag)
var Fun validator.Func = func(fl validator.FieldLevel) bool { //返回false为没错误，验证通过
	age, ok := fl.Field().Interface().(int) //传的是int
	if ok {
		if age<=0||age>=150{//age不能小于1，不能大于149
			return false
		}
		return true
	}
	return true
}

//验证器文档,https://godoc.org/github.com/go-playground/validator
type YanZheng struct {
	Name       string `binding:"min=4,max=20"`         //字节,最小4，最大20
	SecondName string `binding:"require,nefield=Name"` //必须传,并且不能和Name字段值一样
	Ip         string `binding:"ipv4"`                 //必须是ipv4
	Age        int    `binding:"omitempty,gt=5"`       //可以不传,传的话必须大于5,反之写上lt=5传的就必须小于5(切片和map将适用与长度)
}
type DiGui struct {
	YZ []YanZheng `binding:"dive"` //dive将递归验证切片YZ里的所有YanZheng对象
}
