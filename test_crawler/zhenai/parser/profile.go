package parser

import (
	"desktop/test_crawler/engine"
	"desktop/test_crawler/model"
	"fmt"
	"regexp"
	"strconv"
)

//预先初始化正则
var (
	ageRe      = regexp.MustCompile(`<td><span class="label">年龄: </span>([\d]+)岁</td>`)
	marriageRe = regexp.MustCompile(`<td><span class="label">婚况: </span>([^<]+)</td>`)
)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{Name: name}                     //名字已经取到了
	age, err := strconv.Atoi(extractString(contents, ageRe)) //年龄
	if err == nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents, marriageRe) //婚姻情况
	//...昵称、身高、体重等
	result := engine.ParseResult{
		Items: []interface{}{profile},
		//Requests字段可能再产生新的Request
	}
	fmt.Println("用户信息: ",result)
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	//fmt.Println("content: ",string(contents))
	match := re.FindSubmatch(contents)
	//fmt.Println("len: ",len(match))
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
