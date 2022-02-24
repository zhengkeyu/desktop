package main

import (
	"fmt"
	"regexp"
)

//func main() {
//	var str = "abc4c56avxaaccac8cav"
//	//定义一个规则
//	rule := regexp.MustCompile(`a.c`)//匹配"a?c"的字符串
//	//result := rule.FindAllString(str,-1)与底下的函数一样，但是返回的是[]string
//	result := rule.FindAllStringSubmatch(str,-1)//所有匹配"a?c"的字符串都会匹配并放到[][]string中返回，-1表示匹配所有符合的
//	fmt.Println(result)
//}


func main() {//
	//定义一个规则
	//rule := regexp.MustCompile(`<(?U)([a-z]?) class(.+)>(.+)</([a-z]?)>`) //匹配"a?c"的字符串
	rule := regexp.MustCompile(`(?U)<.+class="(date-p|clearfix|clearfix zhiding)".+>`)
	//result := rule.FindAllString(str,-1)//与底下的函数一样，但是返回的是[]string
	result := rule.FindAllStringSubmatch(str, -1) //所有匹配"a?c"的字符串都会匹配并放到[][]string中返回，-1表示匹配所有符合的
	for _, v := range result {
		fmt.Println(v)
	}
}
