package parser

import (
	"regexp"
	"strconv"

	"github.com/yihuaiyuan/learngo/crawler/engine"
	"github.com/yihuaiyuan/learngo/crawler/model"
)

var marriageRe = regexp.MustCompile(`<div class="purple-btns" data-v-8b1eac0c><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^8b1eac0c]*座[^>]*)</div`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^8b1eac0c]+房)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^8b1eac0c]+车)</div>`)
var jiguanRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^>]+)</div>`)

func ParseProfile(contents []byte, gender string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender

	profile.Marriage = extractString(contents, marriageRe)
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Car = extractString(contents, carRe)
	profile.House = extractString(contents, houseRe)
	profile.Jiguan = extractString(contents, jiguanRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	}

	return ""
}
