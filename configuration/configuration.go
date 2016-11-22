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
	JWT         string
}

var (
	configuration Configuration
	configFile    = "application.yaml"
)

func init() {
	LoadConfiguration()
}

// LoadConfiguration : chargement de la configuration.
// Ne pas faire dans l'init parce que sinon pas de TU parces qu'il trouve pas le fichier de conf
func LoadConfiguration() {
	log.Println("Chargement de la configuration ...")
	source, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &configuration)
	if err != nil {
		panic(err)
	}
	log.Println("Configuration chargée !")
}

// GetConfiguration : retourne la configuration de l'application indiquée dans le fichier application.yaml
func GetConfiguration() Configuration {
	return configuration
}

// ReloadConfiguration : rechargement de la configuration de l'application
func ReloadConfiguration() {
	log.Println("Rechargement de la configuration demandé")
	LoadConfiguration()
}
