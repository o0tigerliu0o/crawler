package client

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"crawler/engine"
	"fmt"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%v", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		// 序列化请求
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		// 调用rpc服务
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		// 反序列话后返回结果
		return worker.DeserializeResult(sResult), nil
	}, nil

}
