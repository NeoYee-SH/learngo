package main

import (
	"fmt"

	"github.com/yihuaiyuan/learngo/retriver/really"

	"github.com/yihuaiyuan/learngo/retriver/mock"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func Session(s RetrieverPoster) {
	s.Get("http://www.baidu.com")
	s.Post("http://www.baidu.com", map[string]string{
		"username": "ogenes",
		"password": "12345678",
	})
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com", map[string]string{
		"username": "ogenes",
		"password": "12345678",
	})
}

func download(r Retriever) string {
	return r.Get("http://www.yihuaiyuan.com")
}
func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "this is fake retriever !"}
	//fmt.Println(download(r))
	fmt.Printf("%T , %v\n", r, r) //*mock.Retriever , &{this is fake retriever !}

	r = really.Retriever{UserAgent: "Mozilla/5.0"}
	//fmt.Println(download(r))
	fmt.Printf("%T , %v\n", r, r) //really.Retriever , { 0s}

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(v.Contents)
	case really.Retriever:
		fmt.Println(v.UserAgent)
	}

}
