package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	b, err := ioutil.ReadFile("./ttt.html")
	if err != nil {
		panic(err)
	}
	ParseCity(b)
	//for _, v := range result.Requests {
	//	fmt.Println(v.Url)
	//}
}
