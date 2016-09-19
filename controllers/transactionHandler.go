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

// GetAllTransaction : service qui retourne la liste complète des comptes
func GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	if transaction, e := transactionService.Read(); e != nil {
		errorResponse(e, http.StatusBadRequest, w)
	} else {
		writeHTTPJSONResponse(w, transaction)
	}
}

//SearchTransactionByID :tous est dans le nom
func SearchTransactionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringID := vars["id"]
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre name obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		transactions, err := transactionService.SearchByID(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest, w)
		} else {
			writeHTTPJSONResponse(w, transactions)
		}
	}
}

// CreateTransaction : Réponse sur requete POST a /user avec l'utilisateur en JSON dans le body
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &transaction); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := transactionService.Create(&transaction); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					writeHTTPJSONResponse(w, transaction)
				}
			}
		}
	}
}

// UpdateTransaction : Mise a jour d'une transaction
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &transaction); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := transactionService.Update(&transaction); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}
		}
	}
}

// DeleteTransactionID : reponse http à la demande de suppression d'un utilisateur a partir de son ID
func DeleteTransactionID(w http.ResponseWriter, r *http.Request) {
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
			if err := transactionService.Delete(&model.Transaction{Transactionid: int64(ID)}); err != nil {
				msg := "Suppresion du compte d'id `" + string(ID) + "` impossible. \n" + err.Error()
				errorResponse(&HTTPerror{Code: http.StatusInternalServerError, Message: msg}, http.StatusBadRequest, w)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

// TODO : ajouter les méthodes de recherches en plus du standard (searchby bank, user etc)
