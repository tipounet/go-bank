// test project main.go
package main

// "strings" pour les fonctions sur les chaînes de charactères.
import (
	"log"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tipounet/go-bank/configuration"
	"github.com/tipounet/go-bank/controllers"
	"github.com/tipounet/go-bank/dao"
)

// TODO : Il faut prévoir de vérifier qu'il y ai un utilisateur en base sinon en ajouter un de base pour pouvoir ajouter les autres (utilisateur par défaut admin/4dm1n)
func main() {
	// init de la base de données pour être certain de la fermeture. Reste a voir pour que ce soit automatique à la fin de l'appli ? (equivalent de l'init mais en destroy ?)
	db := dao.GetDbConnexion()
	defer db.Close()

	log.Printf("Rest API Bank Account v%s\n", configuration.GetConfiguration().Version)

	wsContainer := initContainer()
	// Add container filter to enable CORS
	getCORSFilter(wsContainer)
	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)
	// filtre ajoutant les logs des requêtes
	wsContainer.Filter(controllers.NCSACommonLogFormatLogger())

	port := strconv.FormatInt(configuration.GetConfiguration().HTTP.Port, 10)
	log.Printf("start listening on localhost:%v\n", port)
	server := &http.Server{Addr: ":" + port, Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func panicAbord(e error) {
	if e != nil {
		panic(e)
	}
}

// initContainer : création du contener WS et ajout des routes
func initContainer() *restful.Container {
	wsContainer := restful.NewContainer()
	// Ajout utilisation gzip / deflate automatique pour les requetes / reponses
	wsContainer.EnableContentEncoding(true)
	// pretty print ou pas en fonction du fichier de conf (par défaut true du coup ce sera false avec le fichier de configuration lorsque l'info n'est pas précisée)
	restful.PrettyPrintResponses = configuration.GetConfiguration().Prettyprint

	// tableau contenant les "routeurs / ressources"
	ctrl := []controllers.ControllerRessource{
		controllers.HomePageResource{},
		controllers.VersionResource{},
		controllers.BankResource{},
		controllers.UserResource{},
		controllers.AccountResource{},
		controllers.TransactionResource{},
		controllers.TransactionTypeResource{},
	}
	// ajout des "routeurs / ressources" au conteneur global
	for _, c := range ctrl {
		wsContainer.Add(c.RegisterTo())
	}
	return wsContainer
}

// Configuration des entêtes permettant l'accès depuis un serveur autre ue celui ui sert les services : CORS
func getCORSFilter(wsContainer *restful.Container) {
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"Authorization"},
		AllowedHeaders: []string{"Content-Type", "Accept", "Origin", "Authorization", "Lang"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)
}
