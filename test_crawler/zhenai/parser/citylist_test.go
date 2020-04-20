package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

//单元测试
func TestParseCityList(t *testing.T) {
	//测试直接用的本地html文件，以防其他原因影响
	body, err := ioutil.ReadFile("./citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(body)
	if len(result.Requests) != 470 || len(result.Items) != 470 {
		fmt.Println("have error")
	} else {
		fmt.Println("no error")
	}
}
