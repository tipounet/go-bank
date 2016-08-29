package configuration

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// DbConfiguration : conf pour la connexion PG
type DbConfiguration struct {
	Host     string
	Port     int64
	User     string
	Password string
	Schema   string
}

// HTTPConfiguration : la configuration pour le serveur http
type HTTPConfiguration struct {
	Port    int64
	Context string
}

// Configuration : structure contenant les paramètres de configuration de l'application
type Configuration struct {
	Pg          DbConfiguration
	HTTP        HTTPConfiguration
	Version     string
	Prettyprint bool
}

var configuration Configuration

// FIXME : retourner une erreur plutôt que les deux panic ?
func init() {
	log.Printf("Chargement de la configuration")
	source, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &configuration)
	if err != nil {
		panic(err)
	}
	log.Printf("Configuration de l'application : json %v", configuration.Prettyprint)
}

// GetConfiguration : retourne la configuration de l'application indiquée dans le fichier application.yaml
func GetConfiguration() Configuration {
	return configuration
}
