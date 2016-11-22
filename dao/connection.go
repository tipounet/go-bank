package dao

import (
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/tipounet/go-bank/configuration"
)

var db *gorm.DB

// Initialisation de la connexion vers PG avec gorm
func init() {
	conf := configuration.GetConfiguration()
	if db == nil {
		var err error
		cnx := "host=" + conf.Pg.Host
		cnx = cnx + " port=" + strconv.FormatInt(conf.Pg.Port, 10)
		cnx = cnx + " user=" + conf.Pg.User
		cnx = cnx + " dbname=" + conf.Pg.Schema
		cnx = cnx + " sslmode=disable"
		cnx = cnx + " password=" + conf.Pg.Password

		db, err = gorm.Open("postgres", cnx)
		if err != nil {
			panic(err)
		}
		// configuration de la base
		// supprime le fait que gorm cherche une table qui correspond au nom de la struc avec un s au bout
		// db.SingularTable(true)
		// est ce que ça marche en général ça ?
		// db.Set("gorm:save_associations", false)
		db.LogMode(true)
	}
}

// GetDbConnexion : connection à la base de données
func GetDbConnexion() *gorm.DB {
	return db
}
