package parser

import (
	"crawler/single/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
	)
func ParseCity (contents []byte,url string) engine.ParseResult{
	matches:=profileRe.FindAllSubmatch(contents,-1)
	result:=engine.ParseResult{}
	for _,m:=range matches {
		name:=string(m[2])
		//result.Items = append(result.Items, "User: "+ name)
		result.Requests = append(result.Requests,engine.Request{
			Url:        string(m[1]),
			ParserFunc: func (contents []byte,url string) engine.ParseResult{
				return ParseProfile(contents,name,string(m[1]))//闭包
			},
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents,-1) //next page and other filter
	for _,m:=range matches{
		result.Requests = append(result.Requests,engine.Request{
			Url:       string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}