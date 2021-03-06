package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

var port = flag.Int("port", 0, "the port for itemSaver to listen on")

func main() {
	flag.Parse()
	if 0 == *port {
		fmt.Println("must specify a port")
		return
	}

	// 如果服务有问题，直接panic
	log.Fatal(serveRpc(fmt.Sprintf(":%v", *port), config.ElasticIndex))
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
