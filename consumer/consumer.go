package consumer

import (
	"github.com/bitly/go-nsq"
	"log"
)

type topicChan struct {
	topic   string
	channel string
}

type Consumer struct {
	Handlers map[topicChan]*queue
}

func NewConsumer() *Consumer {
	return &Consumer{Handlers: make(map[topicChan]*queue)}
}

// Registers topic/channel handler for messages
func (c *Consumer) Register(topic, channel string, maxInFlight int, handler func(*Message)) error {
	tch := topicChan{topic, channel}
	// Create nsq reader
	r, err := nsq.NewReader(topic, channel)
	if err != nil {
		return err
	}
	r.SetMaxInFlight(maxInFlight)

	q := &queue{handler, r}
	r.AddAsyncHandler(q)
	c.Handlers[tch] = q
	return nil
}

// Connects all readers to NSQ lookupd
func (c *Consumer) ConnectToLookupd(addr string) error {
	for _, q := range c.Handlers {
		if err := q.ConnectToLookupd(addr); err != nil {
			return err
		}
	}
	return nil
}

// Connects all readers to NSQ lookupd instances
func (c *Consumer) ConnectToLookupds(addrs []string) error {
	for _, addr := range addrs {
		if err := c.ConnectToLookupd(addr); err != nil {
			return err
		}
	}
	return nil
}

// Connects all readers to NSQ
func (c *Consumer) ConnectToNSQ(addr string) error {
	for _, q := range c.Handlers {
		if err := q.ConnectToNSQ(addr); err != nil {
			return err
		}
	}
	return nil
}

// Connects all readers to NSQ instances
func (c *Consumer) ConnectToNSQs(addrs []string) error {
	for _, addr := range addrs {
		if err := c.ConnectToNSQ(addr); err != nil {
			return err
		}
	}
	return nil
}

// Just waits
func (c *Consumer) Start(debug bool) {
	if debug {
		for i, q := range c.Handlers {
			log.Printf("Handler: topic=%s channel=%s max=%d\n", i.topic, i.channel, q.MaxInFlight())
		}
	}
	<-make(chan bool)
}
