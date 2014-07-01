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
		Convey("It should not produce any error", nil )
		//	err := PublishJSONAsync()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestPublishJSON(t *testing.T) {
	Convey("Given topic to publish a json message", t, func() {
		Convey("It should not produce any error", nil )
		//	err := PublishJSON()
		//	So(err, ShouldEqual, nil)
		//})
	})
}


func TestConnect(t *testing.T) {
	Convey("Given nsqd address to connect to", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Connect()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestEncJSON(t *testing.T) {
	Convey("Given a json message to encode", t, func() {
		Convey("It should not produce any error", nil )
		//	err := EncJSON()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

