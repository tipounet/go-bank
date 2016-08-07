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

var bankService service.BankService

func initBankService() {
	if bankService.Dao == nil {
		dao := dao.BankDao{
			DB: dao.DbConnect(),
		}
		bankService = service.BankService{
			Dao: &dao,
		}
	}
}

// GetAllBanK : service qui retourne la liste complète des banques
func GetAllBanK(w http.ResponseWriter, r *http.Request) {
	initBankService()
	banks, _ := bankService.Get()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(banks)
}

//SearchBankByID :tous est dans le nom
func SearchBankByID(w http.ResponseWriter, r *http.Request) {
	initBankService()
	vars := mux.Vars(r)
	stringID := vars["bankid"]
	// FIXME : comment je passe d'une string à un int64 ?
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre name obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		banks, err := bankService.Search(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest, w)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(banks)
		}
	}
}

//SearchBankByName :tous est dans le nom
func SearchBankByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	initBankService()
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(HTTPerror{
			Code:    http.StatusBadRequest,
			Message: "Paramètre name obligatoire non vide",
		})
	} else {
		banks, e := bankService.SearchPartialName(name)
		if e != nil {
			errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre name obligatoire non vide"}, http.StatusBadRequest, w)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(banks)
	}
}

// CreateBank : Réponse sur requete POST a /bank avec la bank en JSON dans le body
func CreateBank(w http.ResponseWriter, r *http.Request) {
	initBankService()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var bank model.Bank
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &bank); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := bankService.Create(&bank); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(bank)
				}
			}
		}
	}
}

// UpdateBank : Mise a jour d'une banque
func UpdateBank(w http.ResponseWriter, r *http.Request) {
	initBankService()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var bank model.Bank
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &bank); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := bankService.Update(&bank); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}
		}
	}
}

// DeleteBankID : reponse http à la demande de suppression d'une banque a partir de son ID
func DeleteBankID(w http.ResponseWriter, r *http.Request) {
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
			if err := bankService.Delete(&model.Bank{Bankid: int64(ID)}); err != nil {
				msg := "Suppresion de la banque d'id `" + string(ID) + "` impossible. \n" + err.Error()
				errorResponse(&HTTPerror{Code: http.StatusInternalServerError, Message: msg}, http.StatusBadRequest, w)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}
