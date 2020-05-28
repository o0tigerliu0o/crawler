package engine


// 公共解析函数
type ParserFunc func(content []byte,url string) ParseResult

// 用于存放需要解析的url及用于解析url的函数
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

// 用于存放解析结果
type ParseResult struct {
	Requests []Request
	Items    []Item
}

// 单个结果的结构
type Item struct {
	Url     string      // 结果对应的网页地址
	Type    string      // 结果的类型
	Id      string      // 结果在网站中的Id 用于去重
	Payload interface{} // 结果内容
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
