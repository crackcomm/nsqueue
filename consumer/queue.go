package consumer

import (
	"github.com/nsqio/go-nsq"
)

type queue struct {
	fnc Handler
	*nsq.Consumer
}

func (q *queue) HandleMessage(message *nsq.Message) error {
	q.fnc(&Message{message})
	return nil
}
