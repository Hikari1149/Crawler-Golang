package main

import (
	"crawler/single/distribute/config"
	"crawler/single/distribute/persist"
	"crawler/single/distribute/rpcSupport"
	"fmt"
	"github.com/olivere/elastic"
)

func main() {
	client,err:=elastic.NewClient(
		elastic.SetSniff(false))
	if err!=nil{
		panic(err)
	}
	rpcSupport.ServeRpc(fmt.Sprintf(":%d",config.ItemSaverPort),&persist.ItemSaverService{
		Client:client,
		Index:config.ElasticIndex,

	})

}
