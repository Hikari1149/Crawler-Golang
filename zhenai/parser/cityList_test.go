package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents,err:=ioutil.ReadFile(
		"cityList_test_data.html")
	if err!=nil{
		panic(err)
	}
	result:=ParseCityList(contents)
	const resultSize = 470
	if len(result.Requests)!=resultSize{
		t.Errorf("result should have %d but has %d",resultSize,len(result.Requests))
	}
}
