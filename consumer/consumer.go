package consumer

import (
	"log"

	"github.com/bitly/go-nsq"
	"github.com/crackcomm/nsqueue/nsqlog"
)

type topicChan struct {
	topic   string
	channel string
}

// Consumer - NSQ messages consumer.
type Consumer struct {
	Logger   *log.Logger
	LogLevel *nsq.LogLevel

	handlers map[topicChan]*queue
}

// New - Creates a new consumer structure
func New() *Consumer {
	return &Consumer{
		handlers: make(map[topicChan]*queue),
	}
}

// Register - Registers topic/channel handler for messages
// This function creates a new nsq.Reader
func (c *Consumer) Register(topic, channel string, maxInFlight int, handler Handler) error {
	tch := topicChan{topic, channel}

	config := nsq.NewConfig()
	config.Set("verbose", true)
	config.Set("max_in_flight", maxInFlight)

	r, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return err
	}

	r.SetLogger(c.logger(), c.loglevel())

	q := &queue{handler, r}
	r.AddConcurrentHandlers(q, maxInFlight)
	c.handlers[tch] = q
	return nil
}

// ConnectLookupd - Connects all readers to NSQ lookupd
func (c *Consumer) ConnectLookupd(addr string) error {
	for _, q := range c.handlers {
		if err := q.ConnectToNSQLookupd(addr); err != nil {
			return err
		}
	}
	return nil
}

// ConnectLookupdList - Connects all readers to NSQ lookupd instances
func (c *Consumer) ConnectLookupdList(addrs []string) error {
	for _, addr := range addrs {
		if err := c.ConnectLookupd(addr); err != nil {
			return err
		}
	}
	return nil
}

// Connect - Connects all readers to NSQ
func (c *Consumer) Connect(addr string) error {
	for _, q := range c.handlers {
		if err := q.ConnectToNSQD(addr); err != nil {
			return err
		}
	}
	return nil
}

// ConnectList - Connects all readers to NSQ instances
func (c *Consumer) ConnectList(addrs []string) error {
	for _, addr := range addrs {
		if err := c.Connect(addr); err != nil {
			return err
		}
	}
	return nil
}

// Start - Just waits
func (c *Consumer) Start(debug bool) error {
	if debug {
		for i := range c.handlers {
			log.Printf("Handler: topic=%s channel=%s\n", i.topic, i.channel)
		}
	}

	<-make(chan bool)

	return nil
}
func (c *Consumer) logger() *log.Logger {
	if c.Logger == nil {
		return nsqlog.Logger
	}
	return c.Logger
}

func (c *Consumer) loglevel() nsq.LogLevel {
	if c.LogLevel == nil {
		return nsqlog.LogLevel
	}
	return *c.LogLevel
}
