package parser

import (
	"regexp"
	"strconv"

	"github.com/yihuaiyuan/learngo/crawler/engine"
	"github.com/yihuaiyuan/learngo/crawler/model"
)

var profileRe = regexp.MustCompile(`个人资料</div> <div class="m-content-box" data-v-8b1eac0c><div class="purple-btns" data-v-8b1eac0c><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div><div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div><div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div></div> <div class="pink-btns" data-v-8b1eac0c>`)

func ParseProfile(contents []byte, gender string, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender

	match := profileRe.FindSubmatch(contents)
	if match != nil {
		profile.Marriage = string(match[1])
		if age, err := strconv.Atoi(string(match[2])); err == nil {
			profile.Age = age
		}
		profile.Xinzuo = string(match[3])
		if height, err := strconv.Atoi(string(match[4])); err == nil {
			profile.Height = height
		}
		if weight, err := strconv.Atoi(string(match[5])); err == nil {
			profile.Weight = weight
		}
		profile.Hukou = string(match[6])
		profile.Income = string(match[7])
		profile.Occupation = string(match[8])
		profile.Education = string(match[9])

	}
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
