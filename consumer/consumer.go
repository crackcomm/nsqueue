package consumer

import (
	"github.com/bitly/go-nsq"
	"log"
)

type topicChan struct {
	topic   string
	channel string
}

// Function that handles incoming message.
type Handler func(*Message)

// NSQ messages consumer.
type Consumer struct {
	handlers map[topicChan]*queue
}

// Creates a new consumer structure
func NewConsumer() *Consumer {
	return &Consumer{
		handlers: make(map[topicChan]*queue),
	}
}

// Registers topic/channel handler for messages
// This function creates a new nsq.Reader
func (c *Consumer) Register(topic, channel string, maxInFlight int, handler Handler) error {
	tch := topicChan{topic, channel}
	// Create nsq reader
	r, err := nsq.NewReader(topic, channel)
	if err != nil {
		return err
	}
	r.SetMaxInFlight(maxInFlight)

	q := &queue{handler, r}
	r.AddAsyncHandler(q)
	c.handlers[tch] = q
	return nil
}

// Connects all readers to NSQ lookupd
func (c *Consumer) ConnectLookupd(addr string) error {
	for _, q := range c.handlers {
		if err := q.ConnectToLookupd(addr); err != nil {
			return err
		}
	}
	return nil
}

// Connects all readers to NSQ lookupd instances
func (c *Consumer) ConnectLookupdList(addrs []string) error {
	for _, addr := range addrs {
		if err := c.ConnectLookupd(addr); err != nil {
			return err
		}
	}
	return nil
}

// Connects all readers to NSQ
func (c *Consumer) Connect(addr string) error {
	for _, q := range c.handlers {
		if err := q.ConnectToNSQ(addr); err != nil {
			return err
		}
	}
	return nil
}

// Connects all readers to NSQ instances
func (c *Consumer) ConnectList(addrs []string) error {
	for _, addr := range addrs {
		if err := c.Connect(addr); err != nil {
			return err
		}
	}
	return nil
}

// Just waits
func (c *Consumer) Start(debug bool) {
	if debug {
		for i, q := range c.handlers {
			log.Printf("Handler: topic=%s channel=%s max=%d\n", i.topic, i.channel, q.MaxInFlight())
		}
	}

	<-make(chan bool)
}
