package engine

import (
	"crawler/single/fetcher"
	"log"
)

func worker(r Request) (ParseResult,error){
	log.Printf("Fetching %s\n",r.Url)
	body,err:=fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("Fetcher err "+"fetching url %s, error:%v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParserFunc(body,r.Url),nil
}
