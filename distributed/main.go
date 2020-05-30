package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/web01/parser"
	"fmt"
)

func main() {
	// 初始化数据存储
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%v",config.ItemSaverPort))
	if nil != err {
		panic(err)
	}

	// 初始化页面解析器
	e := engine.ConcurrentEnigne{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
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
		ParserFunc: parser.ParseCity,
	})

}
