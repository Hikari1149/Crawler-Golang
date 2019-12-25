package engine

import "log"

type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
	RequestProcessor Processor
}
type Processor func(Request) (ParseResult,error)
type Scheduler interface{
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}
type ReadyNotifier interface{
	WorkerReady(chan Request)
}

var visitedUrl = make(map[string]bool)

// url duplicate
func isDuplicate(url string) bool{
	if visitedUrl[url]{
		return true
	}
	visitedUrl[url] =true
	return false
}


func (e *ConcurrentEngine) Run(seeds ...Request){
	out:=make(chan ParseResult)
	e.Scheduler.Run()
	for i:=0;i<e.WorkerCount;i++{
		e.createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _,r:=range seeds{
		if isDuplicate(r.Url){
			continue
		}
		e.Scheduler.Submit(r) //request to chanel
	}

	for {
		result:= <-out
		for _,item:=range result.Items{
			go func() {
				e.ItemChan <- item
			}() //send item to save
		}

		for _,request:=range result.Requests{
			if isDuplicate(request.Url){
				continue
			}
			e.Scheduler.Submit(request) //
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request ,out chan ParseResult,ready ReadyNotifier) {
	go func(){
		for{
			ready.WorkerReady(in) //chan is ready
			request:=<-in
			//result,err:=Worker(request) //replace with rpc
			result,err:=e.RequestProcessor(request)
			if err!=nil{
				log.Printf("process err: %v\n",err)
				continue
			}
			out <-result
		}
	}()

}

