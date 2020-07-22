package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test.html")

	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const resultsize = 470
	if len(result.Requests) != resultsize {
		t.Errorf("result should have %d " + "request; but had %d", resultsize,len(result.Requests))
	}
	if len(result.Items) != resultsize {
		t.Errorf("result should have %d " + "request; but had %d", resultsize,len(result.Items))
	}

}
