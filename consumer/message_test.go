package consumer

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

var (

)

func TestReadJSON(t *testing.T) {
	Convey("Given a JSON message to parse ", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Message.ReadJSON()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestGiveUp(t *testing.T) {
	Convey("Given a message to giveup", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Message.GiveUp()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestSuccess(t *testing.T) {
	Convey("Given a message tht succeed", t, func() {
		Convey("It should not produce any error", nil )
		//	err := Message.Success()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestFail(t *testing.T) {
	Convey("Given a message that failed", t, func() {
		Convey("It should not produce any error",  nil )
		//	err := Message.Fail()
		//	So(err, ShouldEqual, nil)
		//})
	})
}

func TestFinish(t *testing.T) {
	Convey("Given a message that finish", t, func() {
		Convey("It should not produce any error",  nil )
		//	err := Message.Finish(true)
		//	So(err, ShouldEqual, nil)
		//})
	})
}

