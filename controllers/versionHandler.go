package controllers

import (
	restful "github.com/emicklei/go-restful"
	"github.com/tipounet/go-bank/service"
)

var versionService service.VersionService

func init() {
	versionService = service.VersionService{}
}

func getVersion(request *restful.Request, response *restful.Response) {
	response.WriteEntity(versionService.Get())
}

// VersionResource : Ressoure au sens http d'un service rest
type VersionResource struct{}

// RegisterTo : Permet l'enregistrement des Ressoures pour la version dans le container http global
func (v VersionResource) RegisterTo() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/version").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("").To(getVersion))
	return ws
}
