package consumer

import (
	"github.com/bitly/go-nsq"
)

type queue struct {
	handler func(*Message)
	*nsq.Reader
}

func (q *queue) HandleMessage(message *nsq.Message, responseChannel chan *nsq.FinishedMessage) {
	go q.handler(&Message{responseChannel, message})
}
