package main

import (
	"crawler/single/distribute/rpcSupport"
	"crawler/single/distribute/worker"
	"flag"
	"fmt"
	"log"
)
var port = flag.Int("port",0,"the port for me to listen on")


func main(){
	flag.Parse()
	if *port == 0{
		fmt.Println("must specify port")
	}
	log.Fatal(rpcSupport.ServeRpc(
		fmt.Sprintf("%d",*port),
		&worker.CrawlService{}))
 }
