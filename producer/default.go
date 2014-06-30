package producer

import (
	"github.com/bitly/go-nsq"
)

// Instantiates the Producer object
var DefaultProducer = new(Producer)

func Publish(topic string, body []byte) error {
	return DefaultProducer.Publish(topic, body)
}

func PublishJsonAsync(topic string, v interface{}, doneChan chan *nsq.ProducerTransaction, args ...interface{}) error {
	return DefaultProducer.PublishJsonAsync(topic, v, doneChan, args...)
}

func PublishJson(topic string, v interface{}) error {
	return DefaultProducer.PublishJson(topic, v)
}

// PublishAsync sends a message to nsq  topic
func PublishAsync(topic string, body []byte, doneChan chan *nsq.ProducerTransaction, args ...interface{}) error {
	return DefaultProducer.PublishAsync(topic, body, doneChan, args...)
}

func MultiPublish(topic string, body [][]byte) error {
	return DefaultProducer.MultiPublish(topic, body)
}

func MultiPublishAsync(topic string, body [][]byte, doneChan chan *nsq.ProducerTransaction, args ...interface{}) error {
	return DefaultProducer.MultiPublishAsync(topic, body, doneChan, args...)
}

// Connect method initialize the connection to nsq 
func Connect(addr string) {
	DefaultProducer.Connect(addr)
}
