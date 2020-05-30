package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const HOST = ":9000"
	go rpcsupport.ServRpc(HOST, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(HOST)

	if nil != err {
		panic(err)
	}

	req := worker.Request{
		Url: "http://m.zhenai.com/u/1020164114",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "Bien",
		},
	}

	var result worker.ParseResult
	err = client.Call(
		config.CrawlServiceRpc, req, &result)

	if nil != err {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
