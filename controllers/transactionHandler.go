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

var transactionService service.TransactionService

func init() {
	if transactionService.Dao == nil {
		dao := dao.TransactionDao{
			DB: dao.GetDbConnexion(),
		}
		transactionService = service.TransactionService{
			Dao: &dao,
		}
	}
}

// TransactionResource : Ressoure au sens http d'un service rest
type TransactionResource struct{}

// RegisterTo : Permet l'enregistrement des Ressoures pour la version dans le container http global
func (t TransactionResource) RegisterTo() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/transaction").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("").To(t.GetAllTransaction).Filter(jwtFilter))
	ws.Route(ws.GET("/{id}").To(t.SearchTransactionByID).Filter(jwtFilter))
	ws.Route(ws.GET("/account/{id}").To(t.SearchTransactionByAccountID).Filter(jwtFilter))
	ws.Route(ws.POST("").To(t.CreateTransaction).Filter(jwtFilter))
	ws.Route(ws.PUT("").To(t.UpdateTransaction).Filter(jwtFilter))
	ws.Route(ws.DELETE("/{id}").To(t.DeleteTransactionID).Filter(jwtFilter))
	return ws
}

// GetAllTransaction : service qui retourne la liste complète des comptes
func (t TransactionResource) GetAllTransaction(request *restful.Request, response *restful.Response) {
	if transaction, e := transactionService.Read(); e != nil {
		response.WriteError(http.StatusBadRequest, e)
	} else {
		response.WriteEntity(transaction)
	}
}

//SearchTransactionByID :tous est dans le nom
func (t TransactionResource) SearchTransactionByID(request *restful.Request, response *restful.Response) {
	stringID := request.PathParameter("id")
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre id obligatoire")
	} else {
		transaction, err := transactionService.SearchByID(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Printf("Erreur SQL %v \n", err)
			// Le errorResponse est dupliqué parce si j'essai httpCode := http.StatusNotFound
			// Le compilateur couine soit parce qu'il veux pas passer d'un int à un int64 soit l'inverse suivant le type que je donne a httpCode ...
			if err.Error() == "record not found" {
				response.WriteErrorString(http.StatusBadRequest, "Unknown transaction for ID "+stringID)
			} else {
				response.WriteError(http.StatusBadRequest, err)
			}
		} else {
			if transaction.Transactionid == 0 {
				response.WriteErrorString(http.StatusNotFound, "Unknown transaction for ID "+stringID)
			} else {
				response.WriteEntity(transaction)
			}
		}
	}
}

//SearchTransactionByAccountID : Retourne a liste des transactions d'un compte spécifique.
func (t TransactionResource) SearchTransactionByAccountID(request *restful.Request, response *restful.Response) {
	stringID := request.PathParameter("id")
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre id obligatoire")
	} else {
		transactions, err := transactionService.SearchByAccount(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Printf("Erreur SQL %v \n", err)
			// Le errorResponse est dupliqué parce si j'essai httpCode := http.StatusNotFound
			// Le compilateur couine soit parce qu'il veux pas passer d'un int à un int64 soit l'inverse suivant le type que je donne a httpCode ...
			if err.Error() == "record not found" {
				response.WriteErrorString(http.StatusBadRequest, "Unknown transaction for ID "+stringID)
			} else {
				response.WriteError(http.StatusBadRequest, err)
			}
		} else {
			if len(transactions) == 0 {
				response.WriteErrorString(http.StatusBadRequest, "No transaction for account ID "+stringID)
			} else {
				response.WriteEntity(transactions)
			}
		}
	}
}

// CreateTransaction : Réponse sur requete POST a /user avec l'utilisateur en JSON dans le body
func (t TransactionResource) CreateTransaction(request *restful.Request, response *restful.Response) {
	transaction := new(model.Transaction)
	err := request.ReadEntity(&transaction)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if err := transactionService.Create(transaction); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteEntity(transaction)
		}
	}
}

// UpdateTransaction : Mise a jour d'une transaction
func (t TransactionResource) UpdateTransaction(request *restful.Request, response *restful.Response) {
	strID := request.PathParameter("id")
	log.Printf("id dans l'url : %v", strID)
	idOriginal, errConv := strconv.Atoi(strID)
	if errConv != nil {
		msg := "Erreur de conversion\n" + errConv.Error()
		response.WriteErrorString(http.StatusInternalServerError, msg)
	} else {
		transaction := new(model.Transaction)
		err := request.ReadEntity(&transaction)
		if err != nil {
			response.WriteError(http.StatusBadRequest, err)
		} else {
			log.Printf("la transaction a mettre a jour : %v", transaction)
			if err := transactionService.Update(transaction, int64(idOriginal)); err != nil {
				response.WriteError(http.StatusInternalServerError, err)
			} else {
				response.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

// DeleteTransactionID : reponse http à la demande de suppression d'un utilisateur a partir de son ID
func (t TransactionResource) DeleteTransactionID(request *restful.Request, response *restful.Response) {
	strID := request.PathParameter("id")
	if strID == "" {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre ID obligatoire non vide")
	} else {
		ID, errConv := strconv.Atoi(strID)
		if errConv != nil {
			response.WriteError(http.StatusBadRequest, errConv)
		} else {
			if err := transactionService.Delete(&model.Transaction{Transactionid: int64(ID)}); err != nil {
				msg := "Suppresion du compte d'id `" + string(ID) + "` impossible. \n" + err.Error()
				response.WriteErrorString(http.StatusInternalServerError, msg)
			} else {
				response.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

// TODO : ajouter les méthodes de recherches en plus du standard (searchby bank, user etc)
