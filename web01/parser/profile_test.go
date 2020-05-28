package parser

import (
	"crawler/engine"
	"crawler/fetcher"
	"crawler/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	// 我们希望单测能跑通后，以后就一直可以跑通，不受其他的限制，所以将网页存于本地
	contents, err := ioutil.ReadFile("profile_content.html")

	if nil != err {

		panic(err)
	}

	//fmt.Printf("%s\n", contents)
	result := ParseProfile(contents, "伊米阳光", "http://m.zhenai.com/u/1020164114")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element ; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://m.zhenai.com/u/1020164114",
		Type: "zhenai",
		Id:   "1020164114",
		Payload: model.Profile{
			Name:       "Bien",
			Gender:     "男士",
			Age:        35,
			Height:     165,
			Weight:     "一般",
			Income:     "20001-50000元",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "专业顾问",
			WorkDest:   "北京朝阳区",
			Hokou:      "北京",
			Xinzuo:     "魔羯",
			House:      "打算婚后购",
			Car:        "已买",
		},
	}

	// 内容判断
	if actual != expected {
		t.Errorf("expected : %v;\n but was %v", expected, actual)
	}

}

func TestGetProfile(t *testing.T) {
	content, err := fetcher.FetchProfile("http://m.zhenai.com/u/1814582139")
	// 内容判断
	if nil != err {
		t.Errorf("expected : nil;\n but was %v \n content=[%v]", err, content)
	}
	a := ParseProfile(content, "test", "http://m.zhenai.com/u/1020164114")
	fmt.Println(a)
}
