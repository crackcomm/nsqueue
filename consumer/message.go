package consumer

import (
	"bytes"
	"encoding/json"
	"github.com/bitly/go-nsq"
)

// Message - Inherent nsq
type Message struct {
	*nsq.Message
}

// ReadJSON -
func (m *Message) ReadJSON(v interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(m.Body))
	return dec.Decode(v)
}

// GiveUp - Finish message with success state because message never will be possible to process
func (m *Message) GiveUp() error {
	m.Finish(true)
	return nil
}

// Success - Finish message as successfully proccessed
func (m *Message) Success() error {
	m.Finish(true)
	return nil
}

// Fail - Mark message as failed to process
func (m *Message) Fail() error {
	m.Finish(false)
	return nil
}

// Finish - Finish processing message
func (m *Message) Finish(success bool) error {
	if success {
		m.Message.Finish()
	} else {
		m.Message.Requeue(-1)
	}
	return nil
}
