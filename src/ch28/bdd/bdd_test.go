package bdd_test

import "testing"

import "github.com/smartystreets/goconvey/convey"

func TestBdd(t *testing.T) {
	convey.Convey("Given 2 even numbers", t, func() {
		a := 2
		b := 4
		convey.Convey("When add the two numbers", func() {
			c := a + b
			convey.Convey("Then the result is still even", func() {
				convey.So(c%2, convey.ShouldEqual, 0)
			})
		})
	})
}
