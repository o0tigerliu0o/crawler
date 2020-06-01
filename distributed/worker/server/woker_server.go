package main

import (
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for worker to listen on")

func main() {
	flag.Parse()
	if 0 == *port {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(rpcsupport.ServRpc(fmt.Sprintf(":%v", *port),
		worker.CrawlService{}))
}
