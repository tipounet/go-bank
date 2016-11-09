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
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre id obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		transaction, err := transactionService.SearchByID(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Printf("Erreur SQL %v \n", err)
			// Le errorResponse est dupliqué parce si j'essai httpCode := http.StatusNotFound
			// Le compilateur couine soit parce qu'il veux pas passer d'un int à un int64 soit l'inverse suivant le type que je donne a httpCode ...
			if err.Error() == "record not found" {
				errorResponse(&HTTPerror{Code: http.StatusNotFound, Message: "Unknown transaction for ID " + stringID}, http.StatusNotFound, w)
			} else {
				errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest, w)
			}
		} else {
			if transaction.Transactionid == 0 {
				errorResponse(&HTTPerror{Code: http.StatusNotFound, Message: "Unknown transaction for ID " + stringID}, http.StatusNotFound, w)
			} else {
				writeHTTPJSONResponse(w, transaction)
			}
		}
	}
}

//SearchTransactionByAccountID : Retourne a liste des transactions d'un compte spécifique.
func SearchTransactionByAccountID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringID := vars["id"]
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre id obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		transactions, err := transactionService.SearchByAccount(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Printf("Erreur SQL %v \n", err)
			// Le errorResponse est dupliqué parce si j'essai httpCode := http.StatusNotFound
			// Le compilateur couine soit parce qu'il veux pas passer d'un int à un int64 soit l'inverse suivant le type que je donne a httpCode ...
			if err.Error() == "record not found" {
				errorResponse(&HTTPerror{Code: http.StatusNotFound, Message: "Unknown transaction for ID " + stringID}, http.StatusNotFound, w)
			} else {
				errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest, w)
			}
		} else {
			if len(transactions) == 0 {
				errorResponse(&HTTPerror{Code: http.StatusNotFound, Message: "No transaction for account ID " + stringID}, http.StatusNotFound, w)
			} else {
				writeHTTPJSONResponse(w, transactions)
			}
		}
	}
}

// CreateTransaction : Réponse sur requete POST a /user avec l'utilisateur en JSON dans le body
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction model.Transaction
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	log.Printf("\n\nle JSON : %v\n\n", string(body))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &transaction); err != nil {
				errorResponse(err, 422, w)
			} else {
				log.Printf("on essai d'insérer ça %T %v \n\n", transaction, transaction)
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
	vars := mux.Vars(r)
	strID := vars["id"] // l'id est passé dans l'url on fait une mise à complete donc potentiellement l'id aussi
	log.Printf("id dans l'url : %v", strID)
	idOriginal, errConv := strconv.Atoi(strID)
	if errConv != nil {
		msg := "Erreur de conversion\n" + errConv.Error()
		errorResponse(&HTTPerror{Code: http.StatusInternalServerError, Message: msg}, http.StatusBadRequest, w)
	} else {
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
					log.Printf("la transaction a mettre a jour : %v", transaction)
					if err := transactionService.Update(&transaction, int64(idOriginal)); err != nil {
						errorResponse(err, http.StatusInternalServerError, w)
					} else {
						w.WriteHeader(http.StatusNoContent)
					}
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
				w.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

// TODO : ajouter les méthodes de recherches en plus du standard (searchby bank, user etc)
