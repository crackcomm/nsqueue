package main

import (
	"github.com/crackcomm/nsqueue/producer"
)

var (
	nsqdAddr = "127.0.0.1:4150"
)

func main() {
	producer.Connect(nsqdAddr)

	var crawlJob struct {
		host string
	}
	crawlJob.host = "google.com"

	body, _ := producer.EncJson(crawlJob)
	for i := 0; i < 10000; i++ {
		producer.PublishAsync("crawl", body, nil)
		// producer.PublishJsonAsync("crawl", crawlJob, nil)
	}
}
