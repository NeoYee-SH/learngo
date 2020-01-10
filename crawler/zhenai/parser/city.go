package parser

import (
	"regexp"

	"github.com/yihuaiyuan/learngo/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)"[^>]+>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td> `

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		gender := string(m[3])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, gender, name)
			},
		})
		result.Items = append(result.Items, name)
	}
	return result
}
