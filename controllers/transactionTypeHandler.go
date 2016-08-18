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

// GetAllTransactionType : service qui retourne la liste complète des comptes
func GetAllTransactionType(w http.ResponseWriter, r *http.Request) {
	tts, _ := transactionTypeService.Read()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tts)
}

//SearchTransactionTypeByID :tous est dans le nom
func SearchTransactionTypeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringID := vars["id"]
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre name obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		types, err := transactionTypeService.SearchByID(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest, w)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(types)
		}
	}
}

// CreateTransactionType : Réponse sur requete POST a /typeTransaction
func CreateTransactionType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var ttype model.TransactionType
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &ttype); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := transactionTypeService.Create(&ttype); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(ttype)
				}
			}
		}
	}
}

// UpdateTransactionType : Mise a jour d'un type de transaction
func UpdateTransactionType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var tt model.TransactionType
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &tt); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := transactionTypeService.Update(&tt); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}
		}
	}
}

// DeleteTransactionTypeID : reponse http à la demande de suppression d'un type de transaction
func DeleteTransactionTypeID(w http.ResponseWriter, r *http.Request) {
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
			if err := transactionTypeService.Delete(&model.TransactionType{ID: int64(ID)}); err != nil {
				msg := "Suppresion du type de transaction `" + string(ID) + "` impossible. \n" + err.Error()
				errorResponse(&HTTPerror{Code: http.StatusInternalServerError, Message: msg}, http.StatusBadRequest, w)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

// TODO : ajouter les méthodes de recherches en plus du standard (searchby bank, user etc)
