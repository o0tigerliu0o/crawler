package client

import (
	"crawler/distributed/config"
	"crawler/distributed/worker"
	"crawler/engine"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		// 序列化请求
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		// 调用rpc服务
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		// 反序列话后返回结果
		return worker.DeserializeResult(sResult), nil
	}

}
