package main

import (
	"math"
	"fmt"
)
//接口
type res interface {
	peri()float64
	erea()float64
}
//三角形结构体
type triangle struct {
	leight,width,height float64
}
//圆形结构体
type around struct {
	radiu float64
}
func main() {
var sanjiao = triangle{5.6,7.3,4.3}
var yuan = around{6.5}
test(sanjiao)
test(yuan)
}
//求三角形周长
func (t triangle)peri()float64  {
return t.height+t.leight+t.width
}
//求三角形面积
func (t triangle)erea()float64  {
res := (t.width+t.leight+t.height)/2
return math.Sqrt(res*(res-t.height)*(res-t.leight)*(res-t.width))
}
//求圆形周长
func (t around)peri()float64  {
return t.radiu*2*math.Pi
}
//求圆形面积
func (t around)erea()float64  {
	return t.radiu*t.radiu*math.Pi
}
//判断参数属于三角形还是圆形，并调用对应的方法
func test(r res) {
	if value,ok := r.(triangle);ok{ //if...ok表达式也可用于结构体对象的判断(r是不是triangle结构体的实例化)
		fmt.Println("三角形面积：",value.erea(),"三角形周长：",value.peri())
	}
	if value,ok := r.(around);ok{
		fmt.Println("圆面积：",value.erea(),"圆周长：",value.peri())
	}
}