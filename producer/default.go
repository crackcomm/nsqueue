package producer

import (
	"github.com/bitly/go-nsq"
)

// Instantiates the Producer object
var DefaultProducer *Producer

func init() {
	DefaultProducer = New()
}

// Publish - sends message to nsq  topic
func Publish(topic string, body []byte) error {
	return DefaultProducer.Publish(topic, body)
}

// PublishJSONAsync - sends message to nsq  topic in json format asynchronously
func PublishJSONAsync(topic string, v interface{}, doneChan chan *nsq.ProducerTransaction, args ...interface{}) error {
	return DefaultProducer.PublishJSONAsync(topic, v, doneChan, args...)
}

// PublishJSON - sends message to nsq  topic in json format
func PublishJSON(topic string, v interface{}) error {
	return DefaultProducer.PublishJSON(topic, v)
}

// PublishAsync - sends a message to nsq  topic asynchronously
func PublishAsync(topic string, body []byte, doneChan chan *nsq.ProducerTransaction, args ...interface{}) error {
	return DefaultProducer.PublishAsync(topic, body, doneChan, args...)
}

// MultiPublish - sends multiple message to to nsq  topic
func MultiPublish(topic string, body [][]byte) error {
	return DefaultProducer.MultiPublish(topic, body)
}

// MultiPublishAsync - sends multiple message to nsq  topic asynchronously
func MultiPublishAsync(topic string, body [][]byte, doneChan chan *nsq.ProducerTransaction, args ...interface{}) error {
	return DefaultProducer.MultiPublishAsync(topic, body, doneChan, args...)
}

// Connect method initialize the connection to nsq
func Connect(addr string) error {
	return DefaultProducer.Connect(addr)
}

// ConnectConfig method initialize the connection to nsq with config
func ConnectConfig(addr string, config *nsq.Config) error {
	return DefaultProducer.ConnectConfig(addr, config)
}
