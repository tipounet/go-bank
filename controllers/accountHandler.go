package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

// GetAllAccount : service qui retourne la liste complète des comptes
func GetAllAccount(w http.ResponseWriter, r *http.Request) {
	accounts, _ := accountService.Read()
	writeHTTPJSONResponse(w, accounts)
}

//SearchAccountByID :tous est dans le nom
func SearchAccountByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringID := vars["id"]
	// FIXME : comment je passe d'une string à un int64 ?
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre name obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		account, err := accountService.SearchByID(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest, w)
		} else {
			if account.BankaccountID == 0 {
				errorResponse(&HTTPerror{Code: http.StatusNotFound, Message: "Unknown Account for ID " + stringID}, http.StatusNotFound, w)
			} else {
				writeHTTPJSONResponse(w, account)
			}
		}
	}
}

// CreateAccount : Réponse sur requete POST a /account avec l'utilisateur en JSON dans le body
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &account); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := accountService.Create(&account); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					writeHTTPJSONResponse(w, account)
				}
			}
		}
	}
}

// UpdateAccount : Mise a jour d'un utilisateur
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var account model.Account
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &account); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := accountService.Update(&account); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					w.WriteHeader(http.StatusNoContent)
				}
			}
		}
	}
}

// DeleteAccountID : reponse http à la demande de suppression d'un compte a partir de son ID
func DeleteAccountID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	if strID == "" {
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre ID obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		ID, errConv := strconv.Atoi(strID)
		if errConv != nil {
			msg := "Erreur de conversion\n" + errConv.Error()
			errorResponse(&HTTPerror{Code: http.StatusInternalServerError, Message: msg}, http.StatusBadRequest, w)
		} else {
			if err := accountService.Delete(&model.Account{BankaccountID: int64(ID)}); err != nil {
				msg := "Suppresion du compte d'id `" + string(ID) + "` impossible. \n" + err.Error()
				errorResponse(&HTTPerror{Code: http.StatusInternalServerError, Message: msg}, http.StatusBadRequest, w)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

// DeleteAccountByNumber : reponse http à la demande de suppression d'un compte a partir de son numéro
func DeleteAccountByNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountNumber := vars["number"]
	if accountNumber == "" {
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre accountNumber obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		if err := accountService.Delete(&model.Account{Accountnumber: accountNumber}); err != nil {
			msg := "Suppresion du compte numéro `" + accountNumber + "` impossible. \n" + err.Error()
			errorResponse(&HTTPerror{Code: http.StatusInternalServerError, Message: msg}, http.StatusBadRequest, w)
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

// TODO : ajouter les méthodes de recherches en plus du standar (searchby bank, user etc)
