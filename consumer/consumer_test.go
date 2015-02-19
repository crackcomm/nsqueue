package consumer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	topic       = "consumerTestTopic"
	channel     = "consumerTestChannel"
	maxInFlight = 3

	lookupdHTTPAddr  = "127.0.0.1:4161"
	lookupdHTTPAddrs = []string{"127.0.0.2:4161"}

	destNsqdTCPAddr  = "127.0.0.1:4150"
	destNsqdTCPAddrs = []string{"127.0.0.2:4150"}

	debug = true
)

func mgsHandle(msg *Message) {

}

func TestRegister(t *testing.T) {
	Convey("Given topic, channel, maxInflight and message handler method", t, func() {
		Convey("It should not produce any error", func() {
			err := Register(topic, channel, maxInFlight, mgsHandle)
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Given wrong topic, channel", t, func() {
		Convey("It should produce an error", func() {
			err := Register("", "", maxInFlight, mgsHandle)
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestConnectLookupd(t *testing.T) {
	Convey("Given lookupd address", t, func() {
		Convey("It should not produce any error", func() {
			err := ConnectLookupd(lookupdHTTPAddr)
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Given wrong lookupd address", t, func() {
		Convey("It should produce an error", func() {
			err := ConnectLookupd("127.0.0.1")
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestConnectLookupdList(t *testing.T) {
	Convey("Given list of lookupd address", t, func() {
		Convey("It should not produce any error", func() {
			err := ConnectLookupdList(lookupdHTTPAddrs)
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Given list of wrong lookupd address", t, func() {
		Convey("It should not produce an error", func() {
			err := ConnectLookupdList([]string{"127.0.0.2"})
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestConnect(t *testing.T) {
	Convey("Given nsqd address", t, func() {
		Convey("It should not produce any error", func() {
			err := Connect(destNsqdTCPAddr)
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Given wrong nsqd address", t, func() {
		Convey("It should produce an error", func() {
			err := Connect("127.0.0.1")
			So(err, ShouldNotEqual, nil)
		})
	})
}

func TestConnectList(t *testing.T) {
	Convey("Given list of nsqd address", t, func() {
		Convey("It should not produce any error", func() {
			err := ConnectList(destNsqdTCPAddrs)
			So(err, ShouldEqual, nil)
		})
	})

	Convey("Given list of wrong nsqd address", t, func() {
		Convey("It should produce an error", func() {
			err := ConnectList(destNsqdTCPAddrs)
			So(err, ShouldNotEqual, nil)
		})
	})
}
