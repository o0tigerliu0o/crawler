package main

import (
	"crawler/distributed/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T){
	const HOST=":1234"

	// 启动服务
	go func() {
		err := serveRpc(HOST,"test1")
		if nil != err{
			panic(err)
		}
	}()

	// 让服务先起起来
	time.Sleep(time.Second)

	// 启动服务链接
	client,err := rpcsupport.NewClient(HOST)
	if nil != err{
		panic(err)
	}

	// 调用rpc函数将数据存进去
	item := engine.Item{
		Url:  "http://m.zhenai.com/u/1020164114",
		Type: "zhenai",
		Id:   "1020164114",
		Payload: model.Profile{
			Name:       "Bien",
			Gender:     "男士",
			Age:        35,
			Height:     165,
			Weight:     "一般",
			Income:     "20001-50000元",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "专业顾问",
			WorkDest:   "北京朝阳区",
			Hokou:      "北京",
			Xinzuo:     "魔羯",
			House:      "打算婚后购",
			Car:        "已买",
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save",item,&result)
	if nil != err || result != "ok"{
		t.Errorf("result=[%v] err=[%v]",result,err)
	}
}