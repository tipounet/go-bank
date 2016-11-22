package controllers

import (
	"log"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"

	"github.com/tipounet/go-bank/dao"
	"github.com/tipounet/go-bank/model"
	"github.com/tipounet/go-bank/service"
)

var bankService service.BankService

func init() {
	if bankService.Dao == nil {
		dao := dao.BankDao{
			DB: dao.GetDbConnexion(),
		}
		bankService = service.BankService{
			Dao: &dao,
		}
	}
}

// BankResource : Ressoure au sens http d'un service rest
type BankResource struct{}

// RegisterTo : Permet l'enregistrement des Ressoures pour la version dans le container http global
func (b BankResource) RegisterTo() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/bank").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("").To(b.GetAllBanK).Filter(jwtFilter))
	ws.Route(ws.GET("/{id}").To(b.SearchBankByID).Filter(jwtFilter))
	ws.Route(ws.GET("/name/{name}").To(b.SearchBankByName).Filter(jwtFilter))
	ws.Route(ws.POST("").To(b.CreateBank).Filter(jwtFilter))
	ws.Route(ws.PUT("").To(b.UpdateBank).Filter(jwtFilter))
	ws.Route(ws.DELETE("/{id}").To(b.DeleteBankID).Filter(jwtFilter))
	return ws
}

// GetAllBanK : service qui retourne la liste complète des banques
func (b BankResource) GetAllBanK(request *restful.Request, response *restful.Response) {
	if banks, e := bankService.Get(); e != nil {
		response.WriteError(http.StatusBadRequest, e)
	} else {
		response.WriteEntity(banks)
	}
}

//SearchBankByID :tous est dans le nom
func (b BankResource) SearchBankByID(request *restful.Request, response *restful.Response) {
	stringID := request.PathParameter("id")
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre id obligatoire")
	} else {
		bank, err := bankService.Search(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			response.WriteError(http.StatusBadRequest, err)
		} else {
			if bank.Name == "" {
				response.WriteErrorString(http.StatusBadRequest, "Unknown bank for ID ")
			} else {
				response.WriteEntity(bank)
			}
		}
	}
}

//SearchBankByName :tous est dans le nom
func (b BankResource) SearchBankByName(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	if name == "" {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre name obligatoire non vide")
	} else {
		banks, e := bankService.SearchPartialName(name)
		if e != nil {
			response.WriteError(http.StatusBadRequest, e)
		} else {
			response.WriteEntity(banks)
		}
	}
}

// CreateBank : Réponse sur requete POST a /bank avec la bank en JSON dans le body
func (b BankResource) CreateBank(request *restful.Request, response *restful.Response) {
	bank := new(model.Bank)
	err := request.ReadEntity(&bank)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if err := bankService.Create(bank); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteEntity(bank)
		}
	}
}

// UpdateBank : Mise a jour d'une banque
func (b BankResource) UpdateBank(request *restful.Request, response *restful.Response) {
	bank := new(model.Bank)
	err := request.ReadEntity(&bank)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if err := bankService.Update(bank); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteHeader(http.StatusNoContent)
		}
	}
}

// DeleteBankID : reponse http à la demande de suppression d'une banque a partir de son ID
func (b BankResource) DeleteBankID(request *restful.Request, response *restful.Response) {
	strID := request.PathParameter("id")
	if strID == "" {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre ID obligatoire non vide")
	} else {
		ID, errConv := strconv.Atoi(strID)
		if errConv != nil {
			response.WriteError(http.StatusBadRequest, errConv)
		} else {
			if err := bankService.Delete(&model.Bank{BankID: int64(ID)}); err != nil {
				msg := "Suppresion de la banque d'id `" + string(ID) + "` impossible. \n" + err.Error()
				response.WriteErrorString(http.StatusInternalServerError, msg)
			} else {
				response.WriteHeader(http.StatusNoContent)
			}
		}
	}
}
