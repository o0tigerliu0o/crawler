/* Copyright 2018 Inc. All Rights Reserved. */

/* File : worker */
/*
modification history
--------------------
2018-05-22 09:38 , by o0TigerLiu0o, create
*/
/*
DESCRIPTION
*/
package engine

import (
	"crawler/fetcher"
	"log"
	"strings"
)

func Worker(r Request) (parseResult ParseResult, err error) {
	log.Printf("Fetching: %s", r.Url)
	var body []byte

	if strings.Contains(r.Url, "http://u.zhenai.com/u/") {
		// 将请求中的url抓下来并进行转码
		body, err = fetcher.FetchProfile(r.Url)
		if nil != err {
			return ParseResult{}, err
		}
	} else {
		// 将请求中的url抓下来并进行转码
		body, err = fetcher.Fetch(r.Url)
		if nil != err {
			return ParseResult{}, err
		}
	}
	// 解析转码后的内容获得自己需要的部分
	return r.Parser.Parse(body, r.Url), err
}
