package controllers

import restful "github.com/emicklei/go-restful"

// ControllerRessource : Interface à implementer pour être une resource au sens "restful"
type ControllerRessource interface {
	RegisterTo() *restful.WebService
}
