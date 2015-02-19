package producer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	topic           = "producerTestTopic"
	destNsqdTCPAddr = "127.0.0.1:4150"
)

func TestPublish(t *testing.T) {
	Convey("Given a json message to publish", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddr)
			var message = []byte{0x18}
			err := Publish(topic, message)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestPublishAsync(t *testing.T) {
	Convey("Given a json message to publish asynchronously", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddr)
			var message = []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
			err := PublishAsync(topic, message, nil)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestMultiPublish(t *testing.T) {
	Convey("Given a multiple message to publish", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddr)
			var message1 = []byte{0x18}
			var message = [][]byte{message1}
			err := MultiPublish(topic, message)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestMultiPublishAsync(t *testing.T) {
	Convey("Given a multiple message to publish asynchrnously", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddr)
			var message1 = []byte{0x18}
			var message = [][]byte{message1}
			err := MultiPublishAsync(topic, message, nil)
			So(err, ShouldEqual, nil)
		})
	})
}

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
