package producer

import (
	"github.com/bitly/go-nsq"
)

var DefaultProducer = new(Producer)

func Publish(topic string, body []byte) (int32, []byte, error) {
	return DefaultProducer.Publish(topic, body)
}

func PublishJsonAsync(topic string, v interface{}, doneChan chan *nsq.WriterTransaction, args ...interface{}) error {
	return DefaultProducer.PublishJsonAsync(topic, v, doneChan, args...)
}

func PublishJson(topic string, v interface{}) (int32, []byte, error) {
	return DefaultProducer.PublishJson(topic, v)
}

func PublishAsync(topic string, body []byte, doneChan chan *nsq.WriterTransaction, args ...interface{}) error {
	return DefaultProducer.PublishAsync(topic, body, doneChan, args...)
}

func MultiPublish(topic string, body [][]byte) (int32, []byte, error) {
	return DefaultProducer.MultiPublish(topic, body)
}

func MultiPublishAsync(topic string, body [][]byte, doneChan chan *nsq.WriterTransaction, args ...interface{}) error {
	return DefaultProducer.MultiPublishAsync(topic, body, doneChan, args...)
}

func Connect(addr string) {
	DefaultProducer.Connect(addr)
}
