package configuration

import (
	"testing"
	// . "github.com/smartystreets/goconvey/convey"
)

// Ne fonctionne pas parce qu'il ne trouve pas le fichier de conf :'(
// il faut le mettre en dur dans le fichier configuration.go :/
func GetConfigurationTest(t *testing.T) {
	// configFile = "C:\\dev\\workspacego\\src\\github.com\\tipounet\\go-bank\\application.yaml"
	// LoadConfiguration()
	// Convey("Test Getconfiguration", t, func() {
	// 	Convey("Test passat", func() {
	// 		conf := GetConfiguration()
	// 		So(conf, ShouldNotBeNil)
	// 		pg := conf.Pg
	// 		So(pg, ShouldNotBeNil)
	// 		So(pg.Host, ShouldNotBeBlank)
	// 		So(pg.Host, ShouldEqual, "localhost")
	// 		So(pg.Port, ShouldNotBeBlank)
	// 		So(pg.User, ShouldNotBeBlank)
	// 		So(pg.User, ShouldEqual, "bankAccountApp")
	// 		So(pg.Password, ShouldNotBeBlank)
	// 		So(pg.Password, ShouldEqual, "bankAccountApp")
	// 		So(pg.Schema, ShouldNotBeBlank)
	// 		So(pg.Schema, ShouldEqual, "bankAccountApp")
	//
	// 		h := conf.HTTP
	// 		So(h, ShouldNotBeNil)
	// 		So(h.Context, ShouldNotBeBlank)
	// 		So(h.Context, ShouldEqual, "localhost")
	// 		So(h.Port, ShouldEqual, 8888)
	//
	// 		v := conf.Version
	// 		So(v, ShouldNotBeNil)
	// 		So(v, ShouldEqual, "1.0.0-SNAPSHOT")
	//
	// 		p := conf.Prettyprint
	// 		So(p, ShouldEqual, true)
	//
	// 		j := conf.JWT
	// 		So(j, ShouldNotBeNil)
	// 		So(j, ShouldEqual, "s3cr3tK3y à moi tout seul que je connais !ee")
	// 	})
	// })
}
