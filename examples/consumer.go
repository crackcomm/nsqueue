package main

import (
	"github.com/crackcomm/nsqueue/consumer"
)

var (
	lookupdAddr = "127.0.0.1:4161"
)

func HandleCrawl(message *consumer.Message) {
	var crawlMessage struct {
		Host string
	}

	err := message.ReadJson(&crawlMessage)
	if err != nil {
		log.Printf("Error decoding json msg: %v\n", err)
		message.GiveUp()
		return
	}

	log.Printf("new msg: %s\n", crawlMessage.Host)
	message.Success()
}

func main() {
	consumer.Register("crawl", "queue", 2500, HandleCrawl)
	consumer.ConnectToLookupd(lookupdAddr)
	consumer.Start(true) // starts to wait
}
