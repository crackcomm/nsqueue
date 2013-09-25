package consumer

import (
	"bytes"
	"encoding/json"
	"github.com/bitly/go-nsq"
)

type Message struct {
	responseChannel chan *nsq.FinishedMessage
	*nsq.Message
}

func (m *Message) ReadJson(v interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(m.Body))
	return dec.Decode(v)
}

// Finish message with success state because message never will be possible to process
func (m *Message) GiveUp() {
	m.Finish(true)
}

// Finish message as successfully proccessed
func (m *Message) Success() {
	m.Finish(true)
}

// Mark message as failed to process
func (m *Message) Fail() {
	m.Finish(false)
}

// Finish processing message
func (m *Message) Finish(success bool) {
	m.responseChannel <- &nsq.FinishedMessage{m.Message.Id, 0, success}
}
