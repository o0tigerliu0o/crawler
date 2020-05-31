package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

func main() {
	// 如果服务有问题，直接panic
	log.Fatal(serveRpc(fmt.Sprintf(":%v", config.ItemSaverPort), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if nil != err {
		return err
	}
	// 启动存储服务
	return rpcsupport.ServRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
