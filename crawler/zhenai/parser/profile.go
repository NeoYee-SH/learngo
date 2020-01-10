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
var xinzuoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^8b1eac0c]*座)[^>]*</div`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^8b1eac0c]+房)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^8b1eac0c]+车)</div>`)
var jiguanRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>籍贯:([^>]+)</div>`)
var educationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:[^<]+</div>(<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>){2}`)
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:[^<]+</div>(<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>)`)

func ParseProfile(contents []byte, gender string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender

	profile.Marriage = extractString(contents, marriageRe, 1)
	age, err := strconv.Atoi(extractString(contents, ageRe, 1))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe, 1))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe, 1))
	if err == nil {
		profile.Weight = weight
	}

	profile.Xinzuo = extractString(contents, xinzuoRe, 1)
	profile.Income = extractString(contents, incomeRe, 1)
	profile.Car = extractString(contents, carRe, 1)
	profile.House = extractString(contents, houseRe, 1)
	profile.Jiguan = extractString(contents, jiguanRe, 1)

	education := extractString(contents, educationRe, 2)
	if education == "" {
		profile.Education = extractString(contents, occupationRe, 2)
	} else {
		profile.Education = education
		profile.Occupation = extractString(contents, occupationRe, 2)
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp, n int) string {
	match := re.FindSubmatch(contents)
	if len(match) > n {
		return string(match[n])
	}

	return ""
}
