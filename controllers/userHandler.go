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
				log.Printf("user %v\n\n", user)
				if isEmptyString(user.Pwd) {
					log.Printf("Pwd vide : %v", user.Pwd)
					errorResponse(&HTTPerror{Code: http.StatusBadRequest, Message: "Information de connexion (utilisateur ou / ou mot de passe) manquante(s)"}, http.StatusBadRequest, w)
				} else {
					var aerr error
					var retour bool
					if !isEmptyString(user.Email) {
						log.Printf("Recherche par email : %v\n", user.Email)
						retour, aerr = userService.UserAuthenticateByEMail(user.Email, user.Pwd)

					} else if !isEmptyString(user.Pseudo) {
						log.Printf("Recherche par pseudo : %v\n", user.Email)
						retour, aerr = userService.UserAuthenticate(user.Pseudo, user.Pwd)
					} else {
						log.Printf("Fail y a ni mail ni pseudo %v\n", user)
						aerr = &HTTPerror{Code: http.StatusBadRequest, Message: "Information de connexion (utilisateur ou / ou mot de passe) manquante(s)"}
					}
					log.Printf("le retour de l'authentification %v\n", retour)
					if aerr != nil {
						errorResponse(aerr, http.StatusBadRequest, w)
					} else {
						if retour {
							u := getUser(user)
							log.Printf("L'utilisateur trouvé : %v\n", u)
							writeHTTPJSONResponse(w, u)
						}
					}
				}
			}
		}
	}
}

// retourne un utilisateur complet depuis un utilisateur avec soit id, soit email soit pseudo si non rien fail !
func getUser(user model.User) (retour model.User) {
	// voir si c'est bien la valeur par défaut et surtout ne pas filler de zéro comme id pour l'utilisateur ;)
	if user.UserID != 0 {
		retour, _ = userService.Search(user.UserID)
	} else if !isEmptyString(user.Email) {
		retour, _ = userService.GetByEmail(user.Email)
	} else if !isEmptyString(user.Pseudo) {
		retour, _ = userService.GetByPseudo(user.Pseudo)
	} else {
		// epic fail !
	}

	return
}

// cette fonction ne fonctionne pas, comment tester correctement qu'une chaine de caractère est vide ????
func isEmptyString(s string) (retour bool) {
	retour = true
	// TODO : voir le fonctionnement du trim en go !
	if len(s) != 0 {
		retour = false
	}
	return
}
