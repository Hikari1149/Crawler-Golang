package client

import (
	"crawler/single/distribute/config"
	"crawler/single/distribute/worker"
	"crawler/single/engine"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor,error){
/*	client,err:=rpcSupport.NewClient(
		fmt.Sprintf(":%d",config.WorkerPort0))
	if err!=nil{
		return nil,err
	}*/


	return func(request engine.Request) (engine.ParseResult,error) {
		sReq:=worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c:=<-clientChan
		err:=c.Call(config.CrawlServiceRpc+".Process",sReq,&sResult)
		if err!=nil{
			return engine.ParseResult{},err
		}
		return worker.DeserializeResult(sResult)
	},nil


}
