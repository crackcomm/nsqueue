package consumer

import (
	"encoding/json"

	"github.com/bitly/go-nsq"
)

// Message - Inherent nsq
type Message struct {
	*nsq.Message
}

// GiveUp - Finish message with success state because message never will be possible to process
func (m *Message) GiveUp() {
	m.Finish(true)
}

// Success - Finish message as successfully proccessed
func (m *Message) Success() {
	m.Finish(true)
}

// Fail - Mark message as failed to process
func (m *Message) Fail() {
	m.Finish(false)
}

// Finish - Finish processing message
func (m *Message) Finish(success bool) {
	if success {
		m.Message.Finish()
	} else {
		m.Message.Requeue(-1)
	}
}

// ReadJSON - Unmarshals JSON message body to interface.
func (m *Message) ReadJSON(v interface{}) error {
	return json.Unmarshal(m.Body, v)
}
