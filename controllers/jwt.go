package controllers

import (
	"fmt"
	"log"
	"net/http"
)

// jwtHandler : http handler permettant de vérifier si un token jwt existe et s'il est valide.
// TODO mettre sa en session pour le cas ou ?
func jwtHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		log.Printf("le header jwt : %v\n", token)
		var err error
		if token == "" {
			log.Printf("le token n'est pas dans le header on cherche dans les cookies")
			c, e := r.Cookie("jwt")
			if e == nil && c.Value != "" {
				log.Printf("On a un token jwt %s", c.Value)
				token = c.Value
			} else {
				err = e
			}
		}
		if token != "" {
			if user, ok := jwt.ParseToken(token); ok {
				log.Printf("Email de l'utilisateur :%v", user)
				// renouveller le jeton pour pas être déco
				addJWTtokenToResponse(user, w)
				// get user by mail and put user in session ?
				h.ServeHTTP(w, r)
				return
			}
			err = fmt.Errorf("Erreur d'authentification jwt, voir dans le log en amont (expiré, token ko etc.)")
		} else {
			// pas de jeton
			err = fmt.Errorf("Erreur d'authentification pas de token jwt")
		}
		// a priori le cookie et le header n'existent pas donc forbiden
		errorResponse(err, http.StatusUnauthorized, w)
	})
}
