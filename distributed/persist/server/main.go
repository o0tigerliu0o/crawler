package main

import (
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"github.com/olivere/elastic"
	"log"
)

func main(){
	// 如果服务有问题，直接panic
	log.Fatal(serveRpc(":1234","dating_profile_new"))
}

func serveRpc(host,index string) error{
	client,err := elastic.NewClient(elastic.SetSniff(false))
	if nil != err{
		return err
	}
	// 启动存储服务
	return rpcsupport.ServRpc(host,
		&persist.ItemSaverService{
			Client:client,
			Index:index,
		})
}
