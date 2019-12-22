package main

import (
	"crawler/single/frontend/controller"
	"net/http"
)

func main(){
	http.Handle("/",http.FileServer(
		http.Dir("single/frontend/view")))
	http.Handle("/search",controller.CreateSearchResultHandler(
		"single/frontend/view/template.html"))
	err:=http.ListenAndServe(":8888",nil)
	if err!=nil{
		panic(err)
	}


}

