package view

import (
	"testing"
	"html/template"
	"crawler/model"
	common "crawler/frontend/model"
	"os"
	"crawler/engine"
)

func TestTemplate(t *testing.T) {

	template := template.Must(template.ParseFiles("template.html"))

	page := common.SearchResult{}

	if err := template.Execute(os.Stdout, page); err != nil {
		panic(err)
	}
}

func TestTemplate2(t *testing.T) {

	template := template.Must(template.ParseFiles("template.html"))

	page := common.SearchResult{}

	page.Hits = 123
	page.Start = 0
	item := engine.Item {
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

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}


	file, _ := os.Create("template_test.html")

	if err := template.Execute(file, page); err != nil {
		panic(err)
	}
}


func TestSearchResultView(t *testing.T) {

	s := CreateSearchResultView("template.html")

	page := common.SearchResult{}

	page.Hits = 123
	page.Start = 0
	item := engine.Item {
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

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	file, _ := os.Create("template_test.html")

	if err := s.Render(file, page); err != nil {
		panic(err)
	}
}
