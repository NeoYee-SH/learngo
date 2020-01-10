package main

import (
	"github.com/yihuaiyuan/learngo/crawler/engine"
	"github.com/yihuaiyuan/learngo/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
