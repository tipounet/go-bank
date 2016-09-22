package controllers

import (
	"log"
	"net/http"
)

// Middleware (just a http.Handler)
func jwtHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recup du token, si existe
		c, e := r.Cookie("jwt")
		if e != nil || c.Value == "" {
			// a priori le cookie n'existe pas donc forbiden
			log.Printf("Pas de cookie jwt (%v)", e)
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			// FIXME : valider le token !!!
			// do stuff
			log.Printf("On a un token jwt %s", c.Value)
			h.ServeHTTP(w, r)
		}
		// do stuff after ?
	})
}

// TODO :la même chose, pour le header ?
func getJwtTokenInCookie(r *http.Request) {

}
