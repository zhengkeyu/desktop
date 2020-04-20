package engine

import (
	"desktop/test_crawler/fetcher"
	"log"
	"time"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		time.Sleep(time.Second)
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher error,url: %s,error: %v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		//for _, item := range parseResult.Items {
		//	log.Printf("Got item: %v",item)
		//}
	}
}
