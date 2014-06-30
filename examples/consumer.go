package main

import (
	"flag"
	"fmt"
	"github.com/zvelo/nsqueue/consumer"
	"time"
)

var (
	nsqdAddr    = flag.String("nsqd", "127.0.0.1:4150", "nsqd http address")
	maxInFlight = flag.Int("max-in-flight", 30, "Maximum amount of messages in flight to consume")
)

func HandleTest(msg *consumer.Message) {
	t := &time.Time{}
	t.UnmarshalBinary(msg.Body)
	fmt.Printf("Consume latency: %s\n", time.Since(*t))
	msg.Success()
}

func main() {
	flag.Parse()
	consumer.Register("latency-test", "consume", *maxInFlight, HandleTest)
	consumer.Connect(*nsqdAddr)
	consumer.Start(true)
}
