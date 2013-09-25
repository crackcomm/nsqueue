package producer

import (
	"bytes"
	"encoding/json"
	"github.com/bitly/go-nsq"
)

type Producer struct {
	*nsq.Writer
}

func (p *Producer) PublishJsonAsync(topic string, v interface{}, doneChan chan *nsq.WriterTransaction, args ...interface{}) error {
	body, err := EncJson(v)
	if err != nil {
		return err
	}
	return p.PublishAsync(topic, body, doneChan, args...)
}

func (p *Producer) PublishJson(topic string, v interface{}) (int32, []byte, error) {
	body, err := EncJson(v)
	if err != nil {
		return 0, []byte{}, err
	}
	return p.Publish(topic, body)
}

func (p *Producer) Connect(addr string) {
	p.Writer = nsq.NewWriter(addr)
}

func EncJson(v interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(v)
	return buf.Bytes(), err
}
