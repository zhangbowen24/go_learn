package bdd

import (
	. "github.com/smartystreets/goconvey/convey" //直接使用当前包的命名空间
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Give 2 even number", t, func() {
		a := 2
		b := 4

		Convey("When add the two numbers", func() {
			c := a + b

			Convey("Then  the result is still even", func() {
				So(c%2, ShouldEqual, 0)
			})
		})
	})
}
