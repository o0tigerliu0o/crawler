package engine

import (
	"log"
)

type SimpleEngine struct {
}

// 根据传入的请求，进行实际的解析及分析
func (e SimpleEngine) Run(seeds ...Request) {

	// 请求队列
	var requests []Request

	// 将请求的种子放入队列
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		// 从队列获取请求
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if nil != err {
			log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
			continue
		}

		// 对解析后结果中url再次进行转码、解析
		// "parseResult.Requests..." 为切片赋值简写方式
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v\n", item)
		}

	}
}
