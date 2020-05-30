package engine

type Parser interface {
	Parse(content []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

// 用于存放需要解析的url及用于解析url的函数
type Request struct {
	Url    string
	Parser Parser
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

// nil parser
type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// 公共解析函数
type ParserFunc func(
	contents []byte, url string) ParseResult

type FuncParser struct {
	parser ParserFunc // 函数
	name   string     // 函数名
}

func (f *FuncParser) Parse(content []byte, url string) ParseResult {
	return f.parser(content, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
