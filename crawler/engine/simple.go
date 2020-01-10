package engine

import (
	"log"

	"github.com/yihuaiyuan/learngo/crawler/fetcher"
)

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {

	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got Item %v", item)
		}
	}

}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)

	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36",
	}
	body, err := fetcher.Fetch(r.Url, headers)

	if err != nil {
		log.Printf("Fetcher url %s error: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParseFunc(body), nil
}
