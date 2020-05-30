package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(rpcsupport.ServRpc(fmt.Sprintf(":%v", config.WorkerPort0),
		worker.CrawlService{}))
}
