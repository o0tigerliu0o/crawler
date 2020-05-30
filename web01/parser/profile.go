package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

// [\d]+ : 数字匹配
var ageRe = regexp.MustCompile(`<div class="tag" data-v-10352ec0>([\d]+)岁<`)

// [^"] : 去到"之前的内容
var MarrigeRe = regexp.MustCompile(`"marriageString":"([^"]+)"`)

var GenderRe = regexp.MustCompile(`"genderString":"([^"]+)"`)
var HeightRe = regexp.MustCompile(`data-v-3e01facc>([\d]+)cm</div>`)
var OccupationRe = regexp.MustCompile(`[千|万]","([^"]+)","`)
var IncomeRe = regexp.MustCompile(`"salaryString":"([^"]+)"`)
var EducationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
var WeightRe = regexp.MustCompile(`data-v-3e01facc>体型:([^<]+)<`)
var WorkDestRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>工作地:([^:]*[^<])</div>`)
var HokouRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>籍贯:([^:][^<])<`)
var XinzuoRe = regexp.MustCompile(`<div class="tag" data-v-3e01facc>([^>]*[^座])座`)
var HouseRe = regexp.MustCompile(`data-v-3e01facc>([^>]*[^房])房`)
var CarRe = regexp.MustCompile(`data-v-3e01facc>([^>]*[^车])车<`)
var idUrlRe = regexp.MustCompile(`http://m.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	// 年龄
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if nil == err {
		profile.Age = age
	}

	// 身高
	height, err := strconv.Atoi(extractString(contents, HeightRe))
	if nil == err {
		profile.Height = height
	}

	// 性别
	profile.Gender = extractString(contents, GenderRe)

	// 职业
	profile.Occupation = extractString(contents, OccupationRe)

	// 收入
	profile.Income = extractString(contents, IncomeRe)

	// 学历
	profile.Education = extractString(contents, EducationRe)

	// 体型
	profile.Weight = extractString(contents, WeightRe)

	// 工作地
	profile.WorkDest = extractString(contents, WorkDestRe)

	// 籍贯
	profile.Hokou = extractString(contents, HokouRe)

	// 星座
	profile.Xinzuo = extractString(contents, XinzuoRe)

	// 房
	profile.House = extractString(contents, HouseRe)

	// 车
	profile.Car = extractString(contents, CarRe)

	// 婚姻
	profile.Marriage = extractString(contents, MarrigeRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			engine.Item{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if 2 <= len(match) {
		return string(match[1])
	} else {
		return ""
	}
}
