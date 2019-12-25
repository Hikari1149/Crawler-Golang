package main

import (
	itemsaver "crawler/single/distribute/persist/client"
	"crawler/single/distribute/rpcSupport"
	worker "crawler/single/distribute/worker/client"
	"crawler/single/engine"
	"crawler/single/scheduler"
	"crawler/single/zhenai/parser"
	"flag"
	"log"
	"net/rpc"
	"strings"
)
var (
	itemSaverHost = flag.String(
		"itemSaver_host","","itemSaver host")
	workerHosts = flag.String(
		"worker_hosts","","worker hosts (comma separated)")
	)
func main(){
	flag.Parse()
	itemChan,err:=itemsaver.ItemSaver(*itemSaverHost)
	if err!=nil{
		panic(err)
	}
	pool:=createClientPool(strings.Split(*workerHosts,","))
	processor,err:=worker.CreateProcessor(pool)
	e:=engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 3,
		ItemChan:    itemChan,
		RequestProcessor:processor,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		Parser:engine.NewFuncParser(parser.ParseCity,"ParseCity"),
	})
}
func createClientPool(hosts []string) chan *rpc.Client{
	var clients []*rpc.Client
	for _,h:=range hosts{
		client,err:=rpcSupport.NewClient(h)
		if err==nil{
			clients = append(clients,client)
			log.Printf("Connect to %s\n",h)
		}else{
			log.Printf("error connect to %s: %v\n",h,err)
		}
   	}
   	out:=make(chan *rpc.Client)
   	go func() {
   		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
   	return out
}


