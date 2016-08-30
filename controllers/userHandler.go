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

var userService service.UserService

func init() {
	if userService.Dao == nil {
		dao := dao.UserDao{
			DB: dao.GetDbConnexion(),
		}
		userService = service.UserService{
			Dao: &dao,
		}
	}
}

// GetAllUser : service qui retourne la liste complète des utilisateurs
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	if users, e := userService.Read(); e != nil {
		errorResponse(e, http.StatusBadRequest, w)
	} else {
		writeHTTPJSONResponse(w, users)
	}
}

//SearchUserByID :tous est dans le nom
func SearchUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringID := vars["id"]
	// FIXME : comment je passe d'une string à un int64 ?
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Paramètre id obligatoire non vide"}, http.StatusBadRequest, w)
	} else {
		users, err := userService.Search(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest, w)
		} else {
			writeHTTPJSONResponse(w, users)
		}
	}
}

// CreateUser : Réponse sur requete POST a /user avec l'utilisateur en JSON dans le body
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &user); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := userService.Create(&user); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					writeHTTPJSONResponse(w, user)
				}
			}
		}
	}
}

// UpdateUser : Mise a jour d'un utilisateur
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		errorResponse(err, http.StatusBadRequest, w)
	} else {
		if err := r.Body.Close(); err != nil {
			errorResponse(err, http.StatusBadRequest, w)
		} else {
			if err := json.Unmarshal(body, &user); err != nil {
				errorResponse(err, 422, w)
			} else {
				if err := userService.Update(&user); err != nil {
					errorResponse(err, http.StatusInternalServerError, w)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}
		}
	}
}

// DeleteUserID : reponse http à la demande de suppression d'un utilisateur a partir de son ID
func DeleteUserID(w http.ResponseWriter, r *http.Request) {
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
			if err := userService.Delete(&model.User{UserID: int64(ID)}); err != nil {
				msg := "Suppresion du user d'id `" + string(ID) + "` impossible. \n" + err.Error()
				errorResponse(&HTTPerror{Code: http.StatusInternalServerError, Message: msg}, http.StatusBadRequest, w)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}
