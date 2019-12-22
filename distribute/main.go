package main

import (
	"crawler/single/distribute/config"
	"crawler/single/distribute/persist/client"
	"crawler/single/engine"
	"crawler/single/scheduler"
	"crawler/single/zhenai/parser"
	"fmt"
)

func main(){
	itemChan,err:=client.ItemSaver(fmt.Sprintf(":%d",config.ItemSaverPort))
	if err!=nil{
		panic(err)
	}
	e:=engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 3,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		Parser:engine.NewFuncParser(parser.ParseCity,"ParseCity"),
	})
}
