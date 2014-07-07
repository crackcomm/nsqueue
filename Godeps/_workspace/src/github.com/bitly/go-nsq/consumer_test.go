package nsq

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/bitly/go-simplejson"
)

type MyTestHandler struct {
	t                *testing.T
	q                *Consumer
	messagesSent     int
	messagesReceived int
	messagesFailed   int
}

var nullLogger = log.New(ioutil.Discard, "", log.LstdFlags)

func (h *MyTestHandler) LogFailedMessage(message *Message) {
	h.messagesFailed++
	h.q.Stop()
}

func (h *MyTestHandler) HandleMessage(message *Message) error {
	if string(message.Body) == "TOBEFAILED" {
		h.messagesReceived++
		return errors.New("fail this message")
	}

	data, err := simplejson.NewJson(message.Body)
	if err != nil {
		return err
	}

	msg, _ := data.Get("msg").String()
	if msg != "single" && msg != "double" {
		h.t.Error("message 'action' was not correct: ", msg, data)
	}
	h.messagesReceived++
	return nil
}

func SendMessage(t *testing.T, port int, topic string, method string, body []byte) {
	httpclient := &http.Client{}
	endpoint := fmt.Sprintf("http://127.0.0.1:%d/%s?topic=%s", port, method, topic)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	resp, err := httpclient.Do(req)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	resp.Body.Close()
}

func TestConsumer(t *testing.T) {
	consumerTest(t, false, false, false)
}

func TestConsumerTLS(t *testing.T) {
	consumerTest(t, false, false, true)
}

func TestConsumerDeflate(t *testing.T) {
	consumerTest(t, true, false, false)
}

func TestConsumerSnappy(t *testing.T) {
	consumerTest(t, false, true, false)
}

func TestConsumerTLSDeflate(t *testing.T) {
	consumerTest(t, true, false, true)
}

func TestConsumerTLSSnappy(t *testing.T) {
	consumerTest(t, false, true, true)
}

func consumerTest(t *testing.T, deflate bool, snappy bool, tlsv1 bool) {

	topicName := "rdr_test"
	if deflate {
		topicName = topicName + "_deflate"
	} else if snappy {
		topicName = topicName + "_snappy"
	}
	if tlsv1 {
		topicName = topicName + "_tls"
	}
	topicName = topicName + strconv.Itoa(int(time.Now().Unix()))

	config := NewConfig()
	// so that the test can simulate reaching max requeues and a call to LogFailedMessage
	config.DefaultRequeueDelay = 0
	// so that the test wont timeout from backing off
	config.MaxBackoffDuration = time.Millisecond * 50
	config.Deflate = deflate
	config.Snappy = snappy
	config.TlsV1 = tlsv1
	if tlsv1 {
		config.TlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	q, _ := NewConsumer(topicName, "ch", config)
	q.SetLogger(nullLogger, LogLevelInfo)

	h := &MyTestHandler{
		t: t,
		q: q,
	}
	q.SetHandler(h)

	SendMessage(t, 4151, topicName, "put", []byte(`{"msg":"single"}`))
	SendMessage(t, 4151, topicName, "mput", []byte("{\"msg\":\"double\"}\n{\"msg\":\"double\"}"))
	SendMessage(t, 4151, topicName, "put", []byte("TOBEFAILED"))
	h.messagesSent = 4

	addr := "127.0.0.1:4150"
	err := q.ConnectToNSQD(addr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = q.ConnectToNSQD(addr)
	if err == nil {
		t.Fatalf("should not be able to connect to the same NSQ twice")
	}

	<-q.StopChan

	if h.messagesReceived != 8 || h.messagesSent != 4 {
		t.Fatalf("end of test. should have handled a diff number of messages (got %d, sent %d)", h.messagesReceived, h.messagesSent)
	}
	if h.messagesFailed != 1 {
		t.Fatal("failed message not done")
	}
}
