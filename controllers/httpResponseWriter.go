package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tipounet/go-bank/configuration"
	"github.com/tipounet/go-bank/model"
)

// writeHTTPJSONResponse : écrit la réponse sous forme de json
func writeHTTPJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST,  PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type,  Authorization, Lang, Accept")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.WriteHeader(http.StatusOK)

	conf := configuration.GetConfiguration()

	if conf.Prettyprint == true {
		b, _ := json.MarshalIndent(data, "", "    ")
		w.Write(b)
	} else {
		json.NewEncoder(w).Encode(data)
	}
}

// addJWTtokenToResponse ajout du jeton en cookie et en entête 'Authorization'
// func addJWTtokenToResponse(email string, w http.ResponseWriter) {
func addJWTtokenToResponse(user model.User, w http.ResponseWriter) {
	token := jwt.GenerateToken(user)
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(20 * time.Minute),
		Secure:   false, // permet d'avoir le cookie qu'en version securisé en général si url = https
		HttpOnly: true,  // permet de restreindre l'accès au cookie. si true javascript n'y a pas accès (si implementé coté serveur)
	})
	w.Header().Set("Authorization", token)
	// w.Header().Set("Authorization-jwt", token)
}
