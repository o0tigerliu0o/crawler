package worker

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/web01/parser"
	"errors"
	"fmt"
	"log"
)

// {"ParseCityList",nil},{"ProfileParser",username}
type SerializedParser struct {
	Name string // 函数名
	Args interface{}
}

// engine.Request中有interface类型所以无法网络传输
// 所以我们需要定义自己的类型
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(
	r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {

	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}

	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if nil != err {
			log.Printf("error deserializing "+
				"request:%v", err)
			continue
		}

		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList,
			config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity,
			config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if username, ok := p.Args.(string); ok {
			return parser.NewProfileParser(
				username), nil
		} else {
			return nil, fmt.Errorf("invalid "+
				"arg. arg=[%v]", p.Args)
		}
	default:
		return nil, errors.New(fmt.Sprintf("unknow parser name. name=[%v]", p.Name))
	}
}
