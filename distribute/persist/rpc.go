package persist

import (
	"crawler/single/engine"
	"crawler/single/persist"
	"github.com/olivere/elastic"
	"log"
)

type ItemSaverService struct{
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item,result *string) error{
	err:=persist.Save(s.Client,s.Index,item)
	log.Printf("Item %v saved.\n",item)
	if err==nil{
		*result ="ok"
	}else{
		log.Printf("error saving item: %v , eror: %v",item,err)
	}
	return err
}
