package consumer

import (
	"context"
	"encoding/json"

	"github.com/nsqio/go-nsq"
)

// Message - Inherent nsq
type Message struct {
	*nsq.Message
}

var msgkey = "nsqmsg"

// WithMessage - Returns nsq message from context.
func WithMessage(ctx context.Context, msg *Message) context.Context {
	return context.WithValue(ctx, msgkey, msg)
}

// MessageFromContext - Returns nsq message from context.
func MessageFromContext(ctx context.Context) (*Message, bool) {
	value, ok := ctx.Value(msgkey).(*Message)
	return value, ok
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
