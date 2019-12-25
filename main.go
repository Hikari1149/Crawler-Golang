package main

import (
	"crawler/single/engine"
	"crawler/single/persist"
	"crawler/single/scheduler"
	"crawler/single/zhenai/parser"
)

func main() {
	itemChan,err:=persist.ItemSaver("data_profile")
	if err!=nil{
		panic(err)
	}
	e:=engine.ConcurrentEngine{
		Scheduler:  &scheduler.QueuedScheduler{},
		WorkerCount:3,
		ItemChan:itemChan,
		RequestProcessor:engine.Worker,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		Parser:engine.NewFuncParser(parser.ParseCity,"ParseCity"),
	})
}