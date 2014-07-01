package producer

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	topic = "producerTestTopic"

	destNsqdTCPAddr = "127.0.0.1:4150"
)

func TestPublishJSONAsync(t *testing.T) {
	Convey("Given a topic and a message to publish asynchronously", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddr)
			var message interface{} = "testMessage"
			err := PublishJSONAsync(topic, message, nil)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestPublishJSON(t *testing.T) {
	Convey("Given topic to publish a json message", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddr)
			var message interface{} = "testMessage"
			err := PublishJSON(topic, message)
			So(err, ShouldEqual, nil)
		})
	})
}


func TestConnect(t *testing.T) {
	Convey("Given nsqd address to connect to", t, func() {
		Convey("It should not produce any error", func() {
			err := Connect(destNsqdTCPAddr)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestEncJSON(t *testing.T) {
	Convey("Given a json message to encode", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddr)
			var message interface{} = "testMessage"
			_, err := EncJSON(message)
			So(err, ShouldEqual, nil)
		})
	})
}


