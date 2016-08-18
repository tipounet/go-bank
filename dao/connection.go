package dao

import (
	"strconv"

	"github.com/tipounet/go-bank/configuration"
	pg "gopkg.in/pg.v4"
)

var db *pg.DB

func init() {
	conf := configuration.GetConfiguration()
	if db == nil {
		db = pg.Connect(&pg.Options{
			User:     conf.Pg.User,
			Password: conf.Pg.Password,
			Addr:     conf.Pg.Host + ":" + strconv.FormatInt(conf.Pg.Port, 10),
			Database: conf.Pg.Schema,
		})
	}
}

// GetDbConnexion : connection à la base de données
func GetDbConnexion() *pg.DB {
	return db
}
