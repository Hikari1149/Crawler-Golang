package parser

import (
	"crawler/single/engine"
	"crawler/single/model"
	"regexp"
)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var addressRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)
func ParseProfile(contents []byte,name string) engine.ParseResult{
	profile:=model.Profile{}
	profile.Name=name
	profile.Income=extractToString(contents,incomeRe)
	profile.Address=extractToString(contents,addressRe)
	result:=engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extractToString(contents []byte,re *regexp.Regexp) string{
		match:=re.FindSubmatch(contents)
		if len(match)>=2{
			return string(match[1])
		}else{
			return ""
		}
}