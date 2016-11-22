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

var transactionTypeService service.TransactionTypeService

func init() {
	if transactionTypeService.Dao == nil {
		dao := dao.TransactionTypeDao{
			DB: dao.GetDbConnexion(),
		}
		transactionTypeService = service.TransactionTypeService{
			Dao: &dao,
		}
	}
}

// TransactionTypeResource : Ressoure au sens http d'un service rest
type TransactionTypeResource struct{}

// RegisterTo : Permet l'enregistrement des Ressoures pour la version dans le container http global
func (t TransactionTypeResource) RegisterTo() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/transactionType").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("").To(t.GetAllTransactionType).Filter(jwtFilter))
	ws.Route(ws.GET("/{id}").To(t.SearchTransactionTypeByID).Filter(jwtFilter))
	ws.Route(ws.POST("").To(t.CreateTransactionType).Filter(jwtFilter))
	ws.Route(ws.PUT("").To(t.UpdateTransactionType).Filter(jwtFilter))
	ws.Route(ws.DELETE("/{id}").To(t.DeleteTransactionTypeID).Filter(jwtFilter))
	return ws
}

// GetAllTransactionType : service qui retourne la liste complète des comptes
func (t TransactionTypeResource) GetAllTransactionType(request *restful.Request, response *restful.Response) {
	if tts, e := transactionTypeService.Read(); e != nil {
		response.WriteError(http.StatusBadRequest, e)
	} else {
		response.WriteEntity(tts)
	}
}

//SearchTransactionTypeByID :tous est dans le nom
func (t TransactionTypeResource) SearchTransactionTypeByID(request *restful.Request, response *restful.Response) {
	stringID := request.PathParameter("id")
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre name obligatoire non vide")
	} else {
		tt, err := transactionTypeService.SearchByID(int64(ID))
		if err != nil {
			log.Println("Erreur SQL ", err)
			response.WriteError(http.StatusBadRequest, err)
		} else {
			if tt.TransactionTypeID == 0 {
				response.WriteErrorString(http.StatusBadRequest, "Unknown Transaction type for ID "+stringID)
			} else {
				response.WriteEntity(tt)
			}
		}
	}
}

// CreateTransactionType : Réponse sur requete POST a /typeTransaction
func (t TransactionTypeResource) CreateTransactionType(request *restful.Request, response *restful.Response) {
	ttype := new(model.TransactionType)
	err := request.ReadEntity(&ttype)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if err := transactionTypeService.Create(ttype); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteEntity(ttype)
		}
	}
}

// UpdateTransactionType : Mise a jour d'un type de transaction
func (t TransactionTypeResource) UpdateTransactionType(request *restful.Request, response *restful.Response) {
	ttype := new(model.TransactionType)
	err := request.ReadEntity(&ttype)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if err := transactionTypeService.Update(ttype); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteHeader(http.StatusNoContent)
		}
	}
}

// DeleteTransactionTypeID : reponse http à la demande de suppression d'un type de transaction
func (t TransactionTypeResource) DeleteTransactionTypeID(request *restful.Request, response *restful.Response) {
	strID := request.PathParameter("id")
	if strID == "" {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre ID obligatoire non vide")
	} else {
		ID, errConv := strconv.Atoi(strID)
		if errConv != nil {
			msg := "Erreur de conversion\n" + errConv.Error()
			response.WriteErrorString(http.StatusBadRequest, msg)
		} else {
			if err := transactionTypeService.Delete(&model.TransactionType{TransactionTypeID: int64(ID)}); err != nil {
				msg := "Suppresion du type de transaction `" + string(ID) + "` impossible. \n" + err.Error()
				response.WriteErrorString(http.StatusInternalServerError, msg)
			} else {
				response.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

// TODO : ajouter les méthodes de recherches en plus du standard (searchby bank, user etc)
