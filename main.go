package main

import (
	"crawler/distributed/config"
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/web01/parser"
	//"io/ioutil"
)

func main() {
	// 初始化数据存储
	itemChan, err := persist.ItemSaver("dating_profile_new")
	if nil != err {
		panic(err)
	}

	// 初始化页面解析器
	e := engine.ConcurrentEnigne{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	/*// 解析页面
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: ngine.NewFuncParser(parser.ParseCityList,config.ParseCityList)
	})*/

	// 解析页面
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/jinan",
		//Url:        "http://www.zhenai.com/zhenghun/beijing",
		Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
	})

}
