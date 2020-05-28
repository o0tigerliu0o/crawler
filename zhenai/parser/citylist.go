package parser

import (
	"crawler/engine"
	"regexp"
	"strings"
)

// 从网页中获得城市名及url的正则表达式
const CITY_LIST_RE = `"linkContent":"([^"]*[^"])","linkURL":"(http:\\u002F\\u002Fwww.zhenai.com\\u002Fzhenghun\\u002F[a-z0-9]+)"`

func ParseCityList(contents []byte,url string) engine.ParseResult {

	// 拿到未分割的城市列表
	re := regexp.MustCompile(CITY_LIST_RE)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        strings.Replace(string(m[2]), "\\u002F", "/", -1),
			ParserFunc: ParseCity,
		})

	}

	return result
}
