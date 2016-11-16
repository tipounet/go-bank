package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"github.com/tipounet/go-bank/dao"
	"github.com/tipounet/go-bank/model"
	"github.com/tipounet/go-bank/service"
)

var (
	userService service.UserService
	jwt         service.JWTService
)

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
		user, err := userService.Search(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest, w)
		} else {
			if user.UserID == 0 {
				errorResponse(&HTTPerror{Code: http.StatusNotFound, Message: "Unknown User for ID " + stringID}, http.StatusNotFound, w)
			} else {
				writeHTTPJSONResponse(w, user)
			}
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
					w.WriteHeader(http.StatusNoContent)
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
				w.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

//UserAuthenticate : authentification de l'utilisateur, un utilisateur est en paylaod de la requête
func UserAuthenticate(w http.ResponseWriter, r *http.Request) {
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
				if isEmptyString(user.Pwd) {
					log.Printf("Pwd vide : %v", user.Pwd)
					errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Information de connexion (utilisateur ou / ou mot de passe) manquante(s)"}, http.StatusBadRequest, w)
				} else {
					var aerr error
					var retour model.User
					if !isEmptyString(user.Email) {
						log.Printf("Recherche par email : %v\n", user.Email)
						retour, aerr = userService.UserAuthenticateByEMail(user.Email, user.Pwd)
					} else if !isEmptyString(user.Pseudo) {
						log.Printf("Recherche par pseudo : %v %v\n", user.Email, user.Pwd)
						retour, aerr = userService.UserAuthenticate(user.Pseudo, user.Pwd)
					} else {
						log.Printf("Fail y a ni mail ni pseudo %v\n", user)
						aerr = &HTTPerror{Code: http.StatusBadRequest, Message: "Information de connexion (utilisateur ou / ou mot de passe) manquante(s)"}
					}
					if aerr != nil {
						code := http.StatusBadRequest
						errorResponse(aerr, code, w)
					} else {
						if retour.UserID > 0 {
							// suppression du mot de passe de l'objet que l'on renvoit au client.
							retour.Pwd = ""
							addJWTtokenToResponse(retour, w)
							writeHTTPJSONResponse(w, retour)
						} else {
							// FIXME : code http for bad credential, StatusUnauthorized ?
							aerr = &HTTPerror{Code: http.StatusUnauthorized, Message: "Erreur d'authentification, utilisateur inconnu ou mot de passe erroné"}
							errorResponse(aerr, http.StatusUnauthorized, w)
						}
					}
				}
			}
		}
	}
}

// UserLogout : traitement de la requête http DELTE sur /user/logout. Il s'agit de la déconnexion de l'utilisateur
func UserLogout(w http.ResponseWriter, r *http.Request) {
	// en cas de sauvegarde de l'utilisateur connecté en base il faut le supprimer
	// suppression du cookie jwt
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Path:    "/",
		Expires: time.Now().Add(-20 * time.Minute),
	})
	w.Header().Set("Authorization", "")
}

// cette fonction ne fonctionne pas, comment tester correctement qu'une chaine de caractère est vide ????
func isEmptyString(s string) (retour bool) {
	retour = true
	// TODO : voir le fonctionnement du trim en go !
	if s != "" {
		retour = false
	}
	return
}
