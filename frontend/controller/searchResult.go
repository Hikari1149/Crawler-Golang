package controller

import (
	"context"
	"crawler/single/engine"
	"crawler/single/frontend/model"
	"crawler/single/frontend/view"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct{
	view view.SearchResultView
	client *elastic.Client
}
func CreateSearchResultHandler(
	template string) SearchResultHandler{
	client,err:=elastic.NewClient(
		elastic.SetSniff(false))
	if err!=nil{
		panic(err)
	}
	return SearchResultHandler{
		view:view.CreateSearchResultView(template),
		client:client,
	}

}


//
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter,req *http.Request){
	q:=strings.TrimSpace(req.FormValue("q"))
	from,err:=strconv.Atoi(req.FormValue("from"))
	if err!=nil{
		from =0
	}
	//fmt.Fprintf(w,"q=%s, from=%d",q,from)

	var page model.SearchResult
	page,err=h.getSearchResult(q,from)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}

	err=h.view.Render(w,page)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
}
// get data from elasticSearch
func (h SearchResultHandler) getSearchResult(q string,from int) (model.SearchResult,error){
	var result model.SearchResult
	resp,err:=h.client.Search("data_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())
	if err!=nil{
		return result,err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	return result,nil
}