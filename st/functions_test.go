package st

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCapitalize(t *testing.T) {
	Convey("When I #Capitalize a string", t, func() {
		result := Capitalize("hello")

		Convey("Then only the first letter is capitalized", func() {
			So(result, ShouldEqual, "Hello")
		})
	})

	Convey("I can #Capitalize a zero lengthed string", t, func() {
		So(Capitalize(""), ShouldEqual, "")
	})

	Convey("I can #Capitalize a string of length 1", t, func() {
		So(Capitalize("h"), ShouldEqual, "H")
	})
}

func TestDecapitalize(t *testing.T) {
	Convey("When I #Decapitalize a string", t, func() {
		result := Decapitalize("HELLO")

		Convey("Then only the first letter is de-capitalized", func() {
			So(result, ShouldEqual, "hELLO")
		})
	})

	Convey("I can #Decapitalize a zero lengthed string", t, func() {
		So(Decapitalize(""), ShouldEqual, "")
	})

	Convey("I can #Capitalize a string of length 1", t, func() {
		So(Decapitalize("H"), ShouldEqual, "h")
	})
}

func TestPackaged(t *testing.T) {
	Convey("#Packaged replaces dots with slashes", t, func() {
		So(Packaged("com.loyal3.foo"), ShouldEqual, "com/loyal3/foo")
	})
}

func TestRandom(t *testing.T) {
	Convey("When I #Random a string", t, func() {
		result := Random("hello")
		So(len(result), ShouldBeGreaterThan, len("hello"))
	})
}

func TestUpper(t *testing.T) {
	Convey("When I #Upper a string", t, func() {
		result := Upper("hello world")
		Convey("Then I expect all upper cases", func() {
			So(result, ShouldEqual, "HELLO WORLD")
		})
	})
}

func TestLower(t *testing.T) {
	Convey("When I #Lower a string", t, func() {
		result := Lower("Hello WORLD")
		Convey("Then I expect all lower cases", func() {
			So(result, ShouldEqual, "hello world")
		})
	})
}

func TestStart(t *testing.T) {
	Convey("When I #Start a string", t, func() {
		result := Start("hello world")
		So(result, ShouldEqual, "Hello World")
	})
}

func TestNormalize(t *testing.T) {
	Convey("When I #Normalize a string", t, func() {
		result := Normalize("The Time Has Come")
		So(result, ShouldEqual, "the-time-has-come")
	})
}

func TestCamel(t *testing.T) {
	Convey("When I #Camel a string", t, func() {
		result := Camel("hello world")
		So(result, ShouldEqual, "HelloWorld")
	})
}

func TestCamelLower(t *testing.T) {
	Convey("When I #CamelLower a string", t, func() {
		result := CamelLower("hello world")
		So(result, ShouldEqual, "helloWorld")
	})
}

func TestSnake(t *testing.T) {
	Convey("When I #Snake a string", t, func() {
		result := Snake("hello world.argle.bargle")
		So(result, ShouldEqual, "hello_world_argle_bargle")
	})
}
