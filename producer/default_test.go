package producer

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	topicDefault = "producerTestTopic"
	destNsqdTCPAddrDefault = "127.0.0.1:4150"
)

func TestPublish(t *testing.T) {
	Convey("Given a json message to publish", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddrDefault)			
			var message = []byte{0x18}
			err := Publish(topicDefault, message)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestPublishAsync(t *testing.T) {
	Convey("Given a json message to publish asynchronously", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddrDefault)
			var message = []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}			
			err := PublishAsync(topicDefault, message, nil)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestMultiPublish(t *testing.T) {
	Convey("Given a multiple message to publish", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddrDefault)
			var message1 = []byte{0x18}
			var message = [][]byte{message1}
			err := MultiPublish(topicDefault, message)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestMultiPublishAsync(t *testing.T) {
	Convey("Given a multiple message to publish asynchrnously", t, func() {
		Convey("It should not produce any error", func() {
			Connect(destNsqdTCPAddrDefault)
			var message1 = []byte{0x18}
			var message = [][]byte{message1}
			err := MultiPublishAsync(topicDefault, message, nil)
			So(err, ShouldEqual, nil)
		})
	})
}

