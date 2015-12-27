package producer

import (
	"encoding/json"
	"log"

	"github.com/crackcomm/nsqueue/nsqlog"
	"github.com/nsqio/go-nsq"
)

// Producer inherets the nsq Producer object
type Producer struct {
	Logger   *log.Logger
	LogLevel nsq.LogLevel

	*nsq.Producer
}

// New - Creates a new Producer.
func New() *Producer {
	return &Producer{
		Logger:   nsqlog.Logger,
		LogLevel: nsqlog.LogLevel,
	}
}

// Connect - Connects prodocuer to nsq instance.
func (p *Producer) Connect(addr string) (err error) {
	return p.ConnectConfig(addr, nsq.NewConfig())
}

// ConnectConfig method initialize the connection to nsq with config.
func (p *Producer) ConnectConfig(addr string, config *nsq.Config) (err error) {
	p.Producer, err = nsq.NewProducer(addr, config)
	p.Producer.SetLogger(p.Logger, p.LogLevel)
	return
}

// PublishJSONAsync - sends message to nsq  topic in json format asynchronously
func (p *Producer) PublishJSONAsync(topic string, v interface{}, doneChan chan *nsq.ProducerTransaction, args ...interface{}) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return p.PublishAsync(topic, body, doneChan, args...)
}

// PublishJSON - sends message to nsq  topic in json format
func (p *Producer) PublishJSON(topic string, v interface{}) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return p.Publish(topic, body)
}
