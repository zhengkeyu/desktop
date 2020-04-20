package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{} //拉取的有效数据
}

func NilParser([]byte)ParseResult{
	return ParseResult{}
}
