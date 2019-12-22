package parser

import (
	"crawler/single/engine"
	"crawler/single/model"
	"regexp"
)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var addressRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)
var idUrlRe = regexp.MustCompile(`https://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte,name string,url string) engine.ParseResult{
	profile:=model.Profile{}
	profile.Name=name
	profile.Income=extractToString(contents,incomeRe)
	profile.Address=extractToString(contents,addressRe)
	result:=engine.ParseResult{
		Items:[]engine.Item{
			{
				Url:url,
				Type:"zhenai",
				Id:extractToString([]byte(url), idUrlRe),
				Payload:profile,
			},
		},
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
type ProfileParser struct{
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return ParseProfile(contents,p.userName,url)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return "ProfileParser",p.userName
}

func NewProfileParser(name string) *ProfileParser{
	return &ProfileParser{
		userName:name,
	}

}
