package parser

import (
	"crawler/distributed/config"
	"crawler/engine"
	"regexp"
	"strings"
)

var (
	profileRe = regexp.MustCompile(`<th><a href="([^"]*[^"])" target="_blank">([^>]*[^<])</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte, url string) engine.ParseResult {

	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		// 转手机页面，跳过电脑版页面的爬虫拦截
		url := strings.Replace(string(m[1]), "album", "m", 1)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			// 通过闭包传入账号名
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			// 通过闭包传入账号名
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(content []byte, url string) engine.ParseResult {
	return parseProfile(content, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser", p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
