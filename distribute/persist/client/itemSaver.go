package client

import (
	"crawler/single/distribute/config"
	"crawler/single/distribute/rpcSupport"
	"crawler/single/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item,error){
	client,err:=rpcSupport.NewClient(host)
	if err!=nil{
		return nil,err
	}
	out:=make(chan engine.Item)
	go func() {
		itemCount:=0
		for{
			item:=<-out
			itemCount++
			//call rpc to save item
			result:=""
			err:=client.Call(config.ItemSaverRpc,item,&result)
			if err!=nil{
				log.Printf("Item Saver error: %v\n",err)
			}

		}

	}()
	return out,nil
}
