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

var accountService service.AccountService

func init() {
	if accountService.Dao == nil {
		dao := dao.AccountDao{
			DB: dao.GetDbConnexion(),
		}
		accountService = service.AccountService{
			Dao: &dao,
		}
	}
}

// AccountResource : Ressoure au sens http d'un service rest
type AccountResource struct{}

// RegisterTo : Permet l'enregistrement des Ressoures pour la version dans le container http global
func (a AccountResource) RegisterTo() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/account").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("").To(a.GetAllAccount).Filter(jwtFilter))
	ws.Route(ws.GET("/{id}").To(a.SearchAccountByID).Filter(jwtFilter))
	ws.Route(ws.GET("/user/{id}").To(a.SearchAccountByUserID).Filter(jwtFilter))
	ws.Route(ws.GET("/bank/{id}").To(a.SearchByBank).Filter(jwtFilter))
	ws.Route(ws.POST("").To(a.CreateAccount).Filter(jwtFilter))
	ws.Route(ws.PUT("").To(a.UpdateAccount).Filter(jwtFilter))
	ws.Route(ws.DELETE("/{id}").To(a.DeleteAccountID).Filter(jwtFilter))
	ws.Route(ws.DELETE("/number/{id}").To(a.DeleteAccountByNumber).Filter(jwtFilter))
	return ws
}

// GetAllAccount : service qui retourne la liste complète des comptes
func (a AccountResource) GetAllAccount(request *restful.Request, response *restful.Response) {
	accounts, _ := accountService.Read()
	response.WriteEntity(accounts)
}

//SearchAccountByID :tous est dans le nom
func (a AccountResource) SearchAccountByID(request *restful.Request, response *restful.Response) {
	stringID := request.PathParameter("id")
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre id obligatoire")
	} else {
		account, err := accountService.SearchByID(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			response.WriteError(http.StatusBadRequest, err)
		} else {
			if account.BankaccountID == 0 {
				response.WriteErrorString(http.StatusNotFound, "Unknown Account for ID "+stringID)
			} else {
				response.WriteEntity(account)
			}
		}
	}
}

// SearchAccountByUserID : search all account for a UserID
func (a AccountResource) SearchAccountByUserID(request *restful.Request, response *restful.Response) {
	stringID := request.PathParameter("id")
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre id obligatoire")
	} else {
		account, err := accountService.SearchByUserID(int64(ID))
		if err != nil {
			log.Println("Erreur sur le select SQL ", err)
			response.WriteError(http.StatusBadRequest, err)
		} else {
			if len(account) == 0 {
				response.WriteErrorString(http.StatusNotFound, "no Account for UserID "+stringID)
			} else {
				response.WriteEntity(account)
			}
		}
	}
}

// SearchByBank liste des comptes d'une banque
func (a AccountResource) SearchByBank(request *restful.Request, response *restful.Response) {
	stringID := request.PathParameter("id")
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre id obligatoire")
	} else {
		account, err := accountService.SearchByBank(int64(ID))
		if err != nil {
			log.Println("Erreur sur le select SQL ", err)
			response.WriteError(http.StatusBadRequest, err)
		} else {
			if len(account) == 0 {
				response.WriteErrorString(http.StatusNotFound, "no Account for bankid "+stringID)
			} else {
				response.WriteEntity(account)
			}
		}
	}
}

// CreateAccount : Réponse sur requete POST a /account avec l'utilisateur en JSON dans le body
func (a AccountResource) CreateAccount(request *restful.Request, response *restful.Response) {
	account := new(model.Account)
	err := request.ReadEntity(&account)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		log.Printf("Compte à créer : %v\n", account)
		if err := accountService.Create(account); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteEntity(account)
		}
	}
}

// UpdateAccount : Mise a jour d'un utilisateur
func (a AccountResource) UpdateAccount(request *restful.Request, response *restful.Response) {
	account := new(model.Account)
	err := request.ReadEntity(&account)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if err := accountService.Update(account); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteHeader(http.StatusNoContent)
		}
	}
}

// DeleteAccountID : reponse http à la demande de suppression d'un compte a partir de son ID
func (a AccountResource) DeleteAccountID(request *restful.Request, response *restful.Response) {
	strID := request.PathParameter("id")
	if strID == "" {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre ID obligatoire non vide")
	} else {
		ID, errConv := strconv.Atoi(strID)
		if errConv != nil {
			msg := "Erreur de conversion\n" + errConv.Error()
			response.WriteErrorString(http.StatusInternalServerError, msg)
		} else {
			if err := accountService.Delete(&model.Account{BankaccountID: int64(ID)}); err != nil {
				msg := "Suppresion du compte d'id `" + string(ID) + "` impossible. \n" + err.Error()
				response.WriteErrorString(http.StatusInternalServerError, msg)
			} else {
				response.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

// DeleteAccountByNumber : reponse http à la demande de suppression d'un compte a partir de son numéro
func (a AccountResource) DeleteAccountByNumber(request *restful.Request, response *restful.Response) {
	accountNumber := request.PathParameter("number")
	if accountNumber == "" {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre accountNumber obligatoire non vide")
	} else {
		if err := accountService.Delete(&model.Account{Accountnumber: accountNumber}); err != nil {
			msg := "Suppresion du compte numéro `" + accountNumber + "` impossible. \n" + err.Error()
			response.WriteErrorString(http.StatusInternalServerError, msg)
		} else {
			response.WriteHeader(http.StatusNoContent)
		}
	}
}
