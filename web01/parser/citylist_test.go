package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 我们希望单测能跑通后，以后就一直可以跑通，不受其他的限制，所以将网页存于本地
	contents, err := ioutil.ReadFile("cityList_content.html")

	if nil != err {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)
	result := ParseCityList(contents,"http://www.zhenai.com/zhenghun")
	const resultSize = 516
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	// 对url进行判断
	// 数量判断
	if resultSize != len(result.Requests) {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	// url内容判断
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s", i, url, result.Requests[i].Url)
		}
	}

	// 对iterm进行判断
	// 数量判断
	if resultSize != len(result.Items) {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

}
