package main

import (
	"flag"
	"fmt"
	"github.com/zvelo/nsqueue/producer"
	"time"
)

var (
	amount   = flag.Int("amount", 20, "Amount of messages to produce every 100 ms")
	nsqdAddr = flag.String("nsqd", "127.0.0.1:4150", "nsqd tcp address")
)

func main() {
	flag.Parse()
	producer.Connect(*nsqdAddr)

	for _ = range time.NewTicker(100 * time.Millisecond).C {
		fmt.Println("Ping...")
		for i := 0; i < *amount; i++ {
			body, _ := time.Now().MarshalBinary()
			producer.PublishAsync("latency-test", body, nil)
		}
	}
}
