package parser

import (
	"imooc.com/ccmouse/u2pppw/crawler/engine"
	"regexp"
	"imooc.com/ccmouse/u2pppw/crawler/config"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte, _ string) engine.ParseResult {

	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	//limit := 10
	for _, m := range matches {
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, config.ParseCity),
			})
		//limit--
		//if limit <= 0 {
		//	break
		//}
		//fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	//fmt.Println("Matches found:", len(matches))
	return result
}
