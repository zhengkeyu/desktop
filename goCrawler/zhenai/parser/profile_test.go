package parser

import (
	"testing"
	"io/ioutil"
	"imooc.com/ccmouse/u2pppw/crawler/model"
	"imooc.com/ccmouse/u2pppw/crawler/engine"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)

	result := ParseProfile(contents, "http://album.zhenai.com/u/1314495053", "风中的蒲公英")
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "zhenai",
		Id:   "1314495053",
		Payload: model.Profile{
			Name:       "风中的蒲公英",
			Gender:     "女",
			Marriage:   "离异",
			Age:        41,
			Height:     158,
			Weight:     48,
			Income:     "3001-5000元",
			Education:  "中专",
			Occupation: "公务员",
			Hukou:      "四川阿坝",
			Xingzuo:    "处女座",
			House:      "已购房",
			Car:        "未购车",
		},
	}

	if len(result.Items) != 1 {
		t.Errorf("Items should have 1 item, but got %v", result.Items)
	} else {
		actual := result.Items[0]
		if actual != expected {
			t.Errorf("Expected %+v, but got %+v", expected, actual)
		}
	}
}
