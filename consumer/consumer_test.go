package consumer

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	topic = "consumerTestTopic"
	channel = "consumerTestChannel"
	maxInFlight = 1

	lookupdHTTPAddr = "127.0.0.1:4161"
	lookupdHTTPAddrs = []string{"127.0.0.1:4161"}

	destNsqdTCPAddr = "127.0.0.1:4150"
	destNsqdTCPAddrs = []string{"127.0.0.1:4150"}

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
}

func TestConnectLookupd(t *testing.T) {
	Convey("Given lookupd address", t, func() {
		Convey("It should not produce any error", func() {
			err := ConnectLookupd(lookupdHTTPAddr)
			So(err, ShouldEqual, nil)
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
}

func TestConnect(t *testing.T) {
	Convey("Given nsqd address", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Connect(destNsqdTCPAddr)
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestConnectList(t *testing.T) {
	Convey("Given list of nsqd address", t, func() {
		Convey("It should not produce any error", nil )
		//	err := ConnectList(destNsqdTCPAddrs)
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestStart(t *testing.T) {
	Convey("Given nsqd connection", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Start(debug)
		//	So(err, ShouldEqual, nil)
		//})
	})
}


