package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tipounet/go-bank/model"
)

func mainTest() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=bankAccountApp dbname=bankAccountApp sslmode=disable password=bankAccountApp")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	var u model.User
	log.Printf("Type de db : %T\n\n", db)
	db.First(&u, "email=?", "moogli@phpjungle.info")
	log.Printf("utilisateur touvré : %v\n", u)
}
