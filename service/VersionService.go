package service

import (
	"runtime"

	"github.com/tipounet/go-bank/configuration"
	"github.com/tipounet/go-bank/model"
)

// VersionService : le service métier qui récupère la version de l'application
type VersionService struct {
}

// Get : retourne la version de l'application ainsi que la version de GO utilisée
func (service VersionService) Get() model.Version {
	return model.Version{
		Application: configuration.GetConfiguration().Version,
		GoVersion: model.Goversion{
			Version: runtime.Version(),
			OS:      runtime.GOOS,
			Arch:    runtime.GOARCH,
		},
	}
}
