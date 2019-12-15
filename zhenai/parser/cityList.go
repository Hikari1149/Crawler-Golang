package parser

import (
	"crawler/single/engine"
	"regexp"
)
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-5e16505f>([^<]+)</a>`
func ParseCityList (contents []byte) engine.ParseResult{
	re :=regexp.MustCompile(cityListRe)
	matches:=re.FindAllSubmatch(contents,-1)
	result:=engine.ParseResult{}
	for _,m:=range matches {
		result.Items = append(result.Items, "City: "+string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
