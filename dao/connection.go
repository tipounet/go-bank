package dao

import (
	"strconv"

	pg "gopkg.in/pg.v4"
	"github.com/tipounet/go-bank/configuration"
)

var db *pg.DB

// DbConnect : connection à la base de données
func DbConnect() *pg.DB {
	conf := configuration.GetConfiguration()
	if db == nil {
		db = pg.Connect(&pg.Options{
			User:     conf.Pg.User,
			Password: conf.Pg.Password,
			Addr:     conf.Pg.Host + ":" + strconv.FormatInt(conf.Pg.Port, 10),
			Database: conf.Pg.Schema,
		})
	}
	return db
}
