package main

import (
	"strings"
	"fmt"
)

func main() {
	var str string = "huangjiaxunlun"
	var stri string = "a,b,c,d"
	//判断字符串中是否包含指定的字符串，返回bool值
	boolres := strings.Contains(str,"l")
	fmt.Println(boolres)
	//将[]string中的每个字符串用指定的字符串拼接起来形成新的字符串
	str1 := strings.Join([]string{"a","b","c"},"")
	fmt.Println(str1)
    //给定字符串，返回在str字符串中的索引号
	index := strings.Index(str,"hu")
	fmt.Println(index)
	//把一个字符串重复给定的次数后返回
	str2 := strings.Repeat("ha",3)
	fmt.Println(str2)
	//把str字符串里的"un"替换为"oo",1表示替换几个，填-1表示全部替换
	str3 := strings.Replace(str,"un","oo",1)
	fmt.Println(str3)
	//把字符串头尾的"!"去掉	·
	var str22 string = "!!3!hello!!" // => 3!hello
	str4 := strings.Trim(str22,"!")
	fmt.Println("============")
	fmt.Println(str4)
	//把stri字符串按所给的字符串分割成[]string,如果分割的字符串不存在则返回1长度的[]string,如果是"",则返回每个字符的数组
	arraystr := strings.Split(stri,",")
	fmt.Println(len(arraystr))
    //把字符串按空格分割返回[]string
	var str11 string = "1 2 3 4 5"
	arraystr1 := strings.Fields(str11)
	fmt.Println(arraystr1,str11)
	//把字符串变成大写
	var str5 string = "abc"
	res := strings.ToUpper(str5)
	fmt.Println(res)
	//替换字符串
	r := strings.NewReplacer("a","b","c","d")//设定一个替换规则=>(把"a"换成"b",把"c"换成"d")
	fmt.Println(r.Replace("abcba")) //执行替换，放入替换的字符串，生成新字符串

}