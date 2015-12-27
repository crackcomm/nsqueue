package consumer

import (
	"log"

	"github.com/crackcomm/nsqueue/nsqlog"
	"github.com/nsqio/go-nsq"
)

type topicChan struct {
	topic   string
	channel string
}

// Consumer - NSQ messages consumer.
type Consumer struct {
	Logger   *log.Logger
	LogLevel nsq.LogLevel
	Config   *nsq.Config

	handlers map[topicChan]*queue
}

// DefaultVerbose - Default nsq "verbose" option.
var DefaultVerbose = false

// New - Creates a new consumer structure
func New() *Consumer {
	return &Consumer{
		Logger:   nsqlog.Logger,
		LogLevel: nsqlog.LogLevel,
		handlers: make(map[topicChan]*queue),
	}
}

// Register - Registers topic/channel handler for messages
// This function creates a new nsq.Reader
func (c *Consumer) Register(topic, channel string, maxInFlight int, handler Handler) error {
	tch := topicChan{topic, channel}

	var config *nsq.Config
	if c.Config == nil {
		config = nsq.NewConfig()
		config.Set("verbose", DefaultVerbose)
		config.Set("max_in_flight", maxInFlight)
	} else {
		config = c.Config
	}

	r, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return err
	}

	r.SetLogger(c.Logger, c.LogLevel)

	q := &queue{handler, r}
	r.AddConcurrentHandlers(q, maxInFlight)
	c.handlers[tch] = q
	return nil
}

// Connect - Connects all readers to NSQ
func (c *Consumer) Connect(addrs ...string) error {
	for _, q := range c.handlers {
		for _, addr := range addrs {
			if err := q.ConnectToNSQD(addr); err != nil {
				return err
			}
		}
	}
	return nil
}

// ConnectLookupd - Connects all readers to NSQ lookupd
func (c *Consumer) ConnectLookupd(addrs ...string) error {
	for _, q := range c.handlers {
		for _, addr := range addrs {
			if err := q.ConnectToNSQLookupd(addr); err != nil {
				return err
			}
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

// Stop - Gracefully closes all consumers.
func (c *Consumer) Stop() {
	for _, h := range c.handlers {
		h.Stop()
	}
}
