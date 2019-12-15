package main

import (
	"crawler/single/engine"
	"crawler/single/scheduler"
	"crawler/single/zhenai/parser"
)

func main() {
	e:=engine.ConcurrentEngine{
		Scheduler:  &scheduler.QueuedScheduler{},
		WorkerCount:10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}