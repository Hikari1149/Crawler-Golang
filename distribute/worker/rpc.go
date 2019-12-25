package worker

import (
	"crawler/single/engine"
)

type CrawlService struct{}




func (s *CrawlService) Process(req Request,result *ParseResult) error{
	engineReq,err:=DeserializeRequest(req)
	if err!=nil{
		return err
	}
	engineResult,err:=engine.Worker(engineReq)
	if err!=nil{
		return err
	}
	*result = SerializeResult(engineResult)//
	return nil
}
