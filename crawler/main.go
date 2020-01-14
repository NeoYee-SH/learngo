package main

import (
	"github.com/yihuaiyuan/learngo/crawler/engine"
	"github.com/yihuaiyuan/learngo/crawler/scheduler"
	"github.com/yihuaiyuan/learngo/crawler/zhenai/parser"
)

func main() {

	//r := engine.SimpleEngine{}
	r := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
	}
	/*r.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})*/
	r.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun/bangbu",
		ParseFunc: parser.ParseCity,
	})
}
