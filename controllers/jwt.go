package controllers

import (
	"fmt"
	"log"

	restful "github.com/emicklei/go-restful"
	"github.com/tipounet/go-bank/model"
)

// jwtFilter : http handler permettant de vérifier si un token jwt existe et s'il est valide.
func jwtFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	token := req.Request.Header.Get("Authorization")
	var err error
	if token == "" {
		log.Printf("le token n'est pas dans le header on cherche dans les cookies")
		c, e := req.Request.Cookie("jwt")
		if e == nil && c.Value != "" {
			log.Printf("On a un token Authorizarion %s", c.Value)
			token = c.Value
		} else {
			log.Println("Pas de cookie \"jwt\" authentification impossible")
			err = e
		}
	}
	if token != "" {
		if user, ok := jwt.ParseToken(token); ok {
			log.Printf("Utilisateur pour la connexion:%v", user)
			// renouveller le jeton pour pas être déco
			addJWTtokenToResponse(user, resp)
			chain.ProcessFilter(req, resp)
			return
		}
		err = fmt.Errorf("Erreur d'authentification jwt, voir dans le log en amont (expiré, token ko etc.)")
	} else {
		// pas de jeton
		err = fmt.Errorf("Erreur d'authentification pas de token jwt")
	}
	// a priori le cookie et le header n'existent pas donc forbiden
	resp.WriteErrorString(401, "401: Not Authorized\n"+err.Error())
}

// addJWTtokenToResponse ajout du jeton en cookie et en entête 'Authorization'
func addJWTtokenToResponse(user model.User, resp *restful.Response) {
	token := jwt.GenerateToken(user)
	resp.AddHeader("Authorization", token)
}
