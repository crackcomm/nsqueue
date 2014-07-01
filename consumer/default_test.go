package consumer

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	topic1 = "consumerTestTopic"
	channel1 = "consumerTestChannel"
	maxInFlight1 = 1

	lookupdHTTPAddr1 = "127.0.0.1:4161"
	lookupdHTTPAddrs1 = []string{"127.0.0.1:4161"}

	destNsqdTCPAddr1 = "127.0.0.1:4150"
	destNsqdTCPAddrs1 = []string{"127.0.0.1:4150"}

	debug1 = true
)

func mgsHandle1(msg *Message) {

}

func TestRegisterDefault(t *testing.T) {
	Convey("Given topic, channel, maxInflight and message handler method", t, func() {
		Convey("It should not produce any error", func() {
			err := Register(topic1, channel1, maxInFlight1, mgsHandle1)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestConnectLookupdDefault(t *testing.T) {
	Convey("Given lookupd address", t, func() {
		Convey("It should not produce any error", func() {
			err := ConnectLookupd(lookupdHTTPAddr1)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestConnectLookupdListDefault(t *testing.T) {
	Convey("Given list of lookupd address", t, func() {
		Convey("It should not produce any error", func() {
			err := ConnectLookupdList(lookupdHTTPAddrs1)
			So(err, ShouldEqual, nil)
		})
	})
}

func TestConnectDefault(t *testing.T) {
	Convey("Given nsqd address", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Connect(destNsqdTCPAddr1)
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestConnectListDefault(t *testing.T) {
	Convey("Given list of nsqd address", t, func() {
		Convey("It should not produce any error", nil )
		//	err := ConnectList(destNsqdTCPAddrs1)
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestStartDefault(t *testing.T) {
	Convey("Given nsqd connection", t, func() {
		Convey("It should not produce any error", nil )
			//err := Start(debug1)
			//So(err, ShouldEqual, nil)
		//})
	})
}


