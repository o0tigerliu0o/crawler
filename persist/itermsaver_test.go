/* Copyright 2018 Inc. All Rights Reserved. */

/* File : itermsaver_test.go */
/*
modification history
--------------------
2018-05-22 09:38 , by o0TigerLiu0o, create
*/
/*
DESCRIPTION
*/
package persist

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
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

	// 存
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	// 取
	resp, err := client.Get().Index(index).Type(expected.Type).
		Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", string(*resp.Source))

	var actual engine.Item
	// 转码
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	// 转actual中的profile类型
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	// 比较
	if actual != expected {
		t.Errorf("got %v; ecpected %v", actual, expected)
	}

}
