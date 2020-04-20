package parser

import (
	"desktop/test_crawler/engine"
	"regexp"
)

const cityRe2 = `<th><a href="([^=]+)" target="_blank">([^<]*)</a></th>`
func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe2)
	matchers := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchers {
		name := string(m[2])
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}
	return result
}