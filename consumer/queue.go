package consumer

import (
	"github.com/bitly/go-nsq"
)

type queue struct {
	fnc Handler
	*nsq.Reader
}

func (q *queue) HandleMessage(message *nsq.Message, responseChannel chan *nsq.FinishedMessage) {
	go q.fnc(&Message{responseChannel, message})
}
