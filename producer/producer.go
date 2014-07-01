package producer

import (
	"bytes"
	"encoding/json"
	"github.com/bitly/go-nsq"
)

// Producer inherets the nsq Producer object
type Producer struct {
	*nsq.Producer
}

// PublishJSONAsync - sends message to nsq  topic in json format asynchronously
func (p *Producer) PublishJSONAsync(topic string, v interface{}, doneChan chan *nsq.ProducerTransaction, args ...interface{}) error {
	body, err := EncJSON(v)
	if err != nil {
		return err
	}
	return p.PublishAsync(topic, body, doneChan, args...)
}

// PublishJSON - sends message to nsq  topic in json format
func (p *Producer) PublishJSON(topic string, v interface{}) error {
	body, err := EncJSON(v)
	if err != nil {
		return err
	}
	return p.Publish(topic, body)
}

// Connect method initialize the connection to nsq
func (p *Producer) Connect(addr string) error {
	config := nsq.NewConfig()
	p.Producer, _ = nsq.NewProducer(addr, config)
	return nil
}

// EncJSON - 
func EncJSON(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(v)
	return buf.Bytes(), err
}
