package consumer

import (
	"github.com/bitly/go-nsq"
)

type queue struct {
	fnc Handler
	*nsq.Consumer
}

func (q *queue) HandleMessage(message *nsq.Message) error {
	q.fnc(&Message{message})
	return nil
}
