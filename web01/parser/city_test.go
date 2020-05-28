package parser

import (
	"crawler/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	// 我们希望单测能跑通后，以后就一直可以跑通，不受其他的限制，所以将网页存于本地
	contents, err := ioutil.ReadFile("city_content.html")

	if nil != err {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)
	result := ParseCity(contents,"http://www.zhenai.com/zhenghun/aba")
	const resultSize = 19
	expectedUrls := []string{
		"http://album.zhenai.com/u/1022090658",
		"http://album.zhenai.com/u/1536113680",
		"http://album.zhenai.com/u/1561457551",
	}

	expectedUsers := []string{
		"梦幻公主", "迟来的问候", "寻幸福在身边",
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
	// 内容判断
	for i, user := range expectedUsers {
		if result.Items[i].Payload.(model.Profile).Name != user {
			t.Errorf("expected user #%d: %s;but was %s", i, user, result.Items[i].Payload.(model.Profile).Name)
		}
	}

	for i, url := range result.Requests {
		fmt.Println(result.Items[i], "  ", ParseProfile(contents, result.Items[i].Payload.(model.Profile).Name, url.Url))
	}
}
