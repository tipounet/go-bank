package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	restful "github.com/emicklei/go-restful"

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

//UserResource gestion des ressources http pour les utilisateurs
type UserResource struct{}

// RegisterTo : enregistrement des ws
func (u UserResource) RegisterTo() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/user").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("").To(u.GetAllUser).Filter(jwtFilter))
	ws.Route(ws.GET("/{id}").To(u.SearchUserByID).Filter(jwtFilter))
	ws.Route(ws.POST("").To(u.CreateUser).Filter(jwtFilter))
	ws.Route(ws.POST("/authenticate").To(u.UserAuthenticate))
	ws.Route(ws.PUT("").To(u.UpdateUser).Filter(jwtFilter))
	ws.Route(ws.DELETE("/{id}").To(u.DeleteUserID).Filter(jwtFilter))
	ws.Route(ws.DELETE("/logout").To(u.DeleteUserID).Filter(jwtFilter))
	return ws
}

// GetAllUser : service qui retourne la liste complète des utilisateurs
func (u UserResource) GetAllUser(request *restful.Request, response *restful.Response) {
	response.AddHeader("Access-Control-Allow-Origin", "*")
	if users, e := userService.Read(); e != nil {
		response.WriteError(http.StatusBadRequest, e)
	} else {
		response.WriteEntity(users)
	}
}

//SearchUserByID :tous est dans le nom
func (u UserResource) SearchUserByID(request *restful.Request, response *restful.Response) {
	stringID := request.PathParameter("id")
	ID, e := strconv.Atoi(stringID)
	if e != nil {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre id obligatoire")
	} else {
		user, err := userService.Search(int64(ID))
		if err != nil {
			// FIXME meilleur Message
			log.Println("Erreur sur le select SQL ", err)
			response.WriteError(http.StatusBadRequest, err)
		} else {
			if user.UserID == 0 {
				response.WriteErrorString(http.StatusBadRequest, "Unknown User for ID ")
			} else {
				response.WriteEntity(user)
			}
		}
	}
}

// CreateUser : Réponse sur requete POST a /user avec l'utilisateur en JSON dans le body
func (u UserResource) CreateUser(request *restful.Request, response *restful.Response) {
	user := new(model.User)
	err := request.ReadEntity(&user)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if err := userService.Create(user); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteEntity(user)
		}
	}
}

// UpdateUser : Mise a jour d'un utilisateur
func (u UserResource) UpdateUser(request *restful.Request, response *restful.Response) {
	user := new(model.User)
	err := request.ReadEntity(&user)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if err := userService.Update(user); err != nil {
			response.WriteError(http.StatusInternalServerError, err)
		} else {
			response.WriteHeader(http.StatusNoContent)
		}
	}
}

// DeleteUserID : reponse http à la demande de suppression d'un utilisateur a partir de son ID
func (u UserResource) DeleteUserID(request *restful.Request, response *restful.Response) {
	strID := request.PathParameter("id")
	if strID == "" {
		response.WriteErrorString(http.StatusBadRequest, "Paramètre ID obligatoire non vide")
	} else {
		ID, errConv := strconv.Atoi(strID)
		if errConv != nil {
			response.WriteError(http.StatusBadRequest, errConv)
		} else {
			if err := userService.Delete(&model.User{UserID: int64(ID)}); err != nil {
				msg := "Suppresion du user d'id `" + string(ID) + "` impossible. \n" + err.Error()
				response.WriteErrorString(http.StatusInternalServerError, msg)
			} else {
				response.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

//UserAuthenticate : authentification de l'utilisateur, un utilisateur est en paylaod de la requête
func (u UserResource) UserAuthenticate(request *restful.Request, response *restful.Response) {
	user := new(model.User)
	err := request.ReadEntity(&user)
	if err != nil {
		response.WriteError(http.StatusBadRequest, err)
	} else {
		if isEmptyString(user.Pwd) {
			response.WriteErrorString(http.StatusBadRequest, "Information de connexion (utilisateur ou / ou mot de passe) manquante(s)")
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
				aerr = fmt.Errorf("Information de connexion (utilisateur ou / ou mot de passe) manquante(s)")
			}
			if aerr != nil {
				response.WriteError(http.StatusBadRequest, aerr)
			} else {
				if retour.UserID > 0 {
					// suppression du mot de passe de l'objet que l'on renvoit au client.
					retour.Pwd = ""
					addJWTtokenToResponse(retour, response)
					response.WriteEntity(retour)
				} else {
					response.WriteErrorString(http.StatusUnauthorized, "Erreur d'authentification, utilisateur inconnu ou mot de passe erroné")
				}
			}
		}
	}
}

// UserLogout : traitement de la requête http DELTE sur /user/logout. Il s'agit de la déconnexion de l'utilisateur
func (u UserResource) UserLogout(request *restful.Request, response *restful.Response) {
	// en cas de sauvegarde de l'utilisateur connecté en base il faut le supprimer
	// suppression du cookie jwt
	http.SetCookie(response.ResponseWriter, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Path:    "/",
		Expires: time.Now().Add(-20 * time.Minute),
	})
	response.AddHeader("Authorization", "")
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
