package consumer

import (
	"github.com/bitly/go-nsq"
)

type Queue struct {
	Handler func(*Message)
	*nsq.Reader
}

func (q *Queue) HandleMessage(message *nsq.Message, responseChannel chan *nsq.FinishedMessage) {
	go q.Handler(&Message{responseChannel, message})
}
