package persist

import (
	"context"
	"crawler/single/engine"
	"errors"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver(index string) (chan engine.Item,error){
	client,err:=elastic.NewClient(
		//turn off sniff in docker
		elastic.SetSniff(false))
	if err!=nil{
		return nil,err
	}
	out:=make(chan engine.Item)
	go func() {
		itemCount :=0
		for {
			log.Printf("%d\n", itemCount)
			item := <-out //
			log.Printf("Got item %d: %v", itemCount, item)
			itemCount++
			err := Save(client,index,item)
			if err != nil {
				log.Printf("Item Error: %v\n", item)
			}
		}

	}()
	return out,nil
}
func Save(client *elastic.Client,index string,item engine.Item) (err error){
	if item.Type == ""{
		return errors.New("must supply type")

	}
	indexService:=client.Index().Index(index).Type(item.Type).
		BodyJson(item)
	if item.Id !=""{
		indexService.Id(item.Id)
	}
	_,err=indexService.Do(context.Background())
	if err!=nil{
		return err
	}
	return nil
}
