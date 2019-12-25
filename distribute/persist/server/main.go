package main

import (
	"crawler/single/distribute/config"
	"crawler/single/distribute/persist"
	"crawler/single/distribute/rpcSupport"
	"flag"
	"fmt"
	"github.com/olivere/elastic"
)
var port = flag.Int("port",0,"the port for me to listen on")
func main() {
	flag.Parse()
	client,err:=elastic.NewClient(
		elastic.SetSniff(false))
	if err!=nil{
		panic(err)
	}
	rpcSupport.ServeRpc(fmt.Sprintf("%d",*port),&persist.ItemSaverService{
		Client:client,
		Index:config.ElasticIndex,

	})

}
