package persist

import (
	"context"
	"crawler/single/engine"
	"crawler/single/model"
	"encoding/json"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
	profile:= engine.Item{
		Url:     "https://album.zhenai.com/u/1875184795",
		Id:      "123",
		Type:    "zhenai",
		Payload: model.Profile{
			Name:    "test",
			Income:  "123",
			Address: "test 666",
		},
	}
	client,err:=elastic.NewClient(
		elastic.SetSniff(false))
	if err!=nil{
		panic(err)
	}
	const index = "dating_test"
	err =save(client,index,profile)
	if err!=nil {
		panic(err)
	}
	resp,err:=client.Get().Index(index).Type("zhenai").Id(profile.Id).Do(context.Background())
	if err!=nil{
		panic(err)
	}
	t.Logf("%s",resp.Source)
	var actual model.Profile
	err = json.Unmarshal([]byte(resp.Source),&actual)
	if err!=nil{
		panic(err)
	}

}
