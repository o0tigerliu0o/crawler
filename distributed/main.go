package main

import (
	"crawler/distributed/config"
	itemSaver "crawler/distributed/persist/client"
	"crawler/distributed/rpcsupport"
	worker "crawler/distributed/worker/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/web01/parser"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_host", "", "worker host(comma separated)")
)

func main() {
	flag.Parse()

	// 初始化数据存储
	itemChan, err := itemSaver.ItemSaver(fmt.Sprintf("%v", *itemSaverHost))
	if nil != err {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	// 初始化页面解析器
	e := engine.ConcurrentEnigne{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	/*// 解析页面
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/

	// 解析页面
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/jinan",
		//Url:        "http://www.zhenai.com/zhenghun/beijing",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})

}

func createClientPool(hosts []string) chan *rpc.Client {
	clients := make([]*rpc.Client, 0, len(hosts))
	// 创建连接池
	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)
		if nil == err {
			clients = append(clients, client)
			log.Printf("Connected to %v", host)
		} else {
			log.Printf("Error connecting to %v. err=[%v]", host, err)
		}
	}

	out := make(chan *rpc.Client)

	// 将生成的worker链接放入链接池
	go func() {
		for {
			// 顺序循环放入
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
