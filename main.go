// test project main.go
package main

// "strings" pour les fonctions sur les chaînes de charactères.
import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tipounet/go-bank/configuration"
	"github.com/tipounet/go-bank/controllers"
	"github.com/tipounet/go-bank/dao"
)

func main() {
	// init de la base de données pour être certain de la fermeture. Reste a voir pour que ce soit automatique à la fin de l'appli ? (equivalent de l'init mais en destroy ?)
	db := dao.GetDbConnexion()
	defer db.Close()
	fmt.Printf("Rest API Bank Account v%s\n", configuration.GetConfiguration().Version)
	port := strconv.FormatInt(configuration.GetConfiguration().HTTP.Port, 10)
	fmt.Println("Please visite http://localhost:" + port)

	myRouter := controllers.NewRouter()

	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func panicAbord(e error) {
	if e != nil {
		panic(e)
	}
}
