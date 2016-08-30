package controllers

import (
	"net/http"

	"github.com/tipounet/go-bank/service"
)

var versionService service.VersionService

func init() {
	versionService = service.VersionService{}
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	writeHTTPJSONResponse(w, versionService.Get())
}
