package producer

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	topicDefault = "producerTestTopic"

	destNsqdTCPAddrDefault = "127.0.0.1:4150"
)

func TestPublishJSONAsyncDefault(t *testing.T) {
	Convey("Given a topic and a message to publish asynchronously", t, func() {
		Convey("It should not produce any error", nil )
		//	err := PublishJSONAsync()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestPublishJSONDefault(t *testing.T) {
	Convey("Given topic to publish a json message", t, func() {
		Convey("It should not produce any error", nil )
		//	err := PublishJSON()
		//	So(err, ShouldEqual, nil)
		//})
	})
}


func TestConnectDefault(t *testing.T) {
	Convey("Given nsqd address to connect to", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Connect()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestPublishDefault(t *testing.T) {
	Convey("Given a json message to publish", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Publish()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestPublishAsyncDefault(t *testing.T) {
	Convey("Given a json message to publish asynchronously", t, func() {
		Convey("It should not produce any error", nil )
		//	err := PublishAsync()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestMultiPublishDefault(t *testing.T) {
	Convey("Given a multiple message to publish", t, func() {
		Convey("It should not produce any error", nil )
		//	err := MultiPublish()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestMultiPublishAsyncDefault(t *testing.T) {
	Convey("Given a multiple message to publish asynchrnously", t, func() {
		Convey("It should not produce any error", nil )
		//	err := MultiPublishAsync()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

