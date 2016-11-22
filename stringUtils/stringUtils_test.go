// Package stringUtils : from https://github.com/elgs/gostrgen
package stringUtils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// RandGen : Génération d'une chaîne alléatoire.

func RandGenTest(t *testing.T) {
	Convey("Test génération de chaine alléatoire", t, func() {
		Convey("RandGenTest()", func() {
			saltLength := 64
			salt, err := RandGen(saltLength, All, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)
			// StringMatches ?
			saltLength = 8
			salt, err = RandGen(saltLength, Lower, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, Upper, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, Digit, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, Punct, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, LowerUpper, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, LowerDigit, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, UpperDigit, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, LowerUpperDigit, "", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, All, "[]{}<>", "")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)

			salt, err = RandGen(saltLength, All, "[]{}<>", "Ol")
			So(salt, ShouldNotBeBlank)
			So(len(salt), ShouldEqual, saltLength)
			So(err, ShouldBeNil)
		})
	})
}
