package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// HTTPerror : Message d'erreur lors d'une requête
type HTTPerror struct {
	Code    int64
	Message string
}

// Implemente l'interface error et du coup on peut retourner une HTTPerror comme error
func (e *HTTPerror) Error() string {
	return e.Message
}

// errorResponse : Création d'une réponse avec objet JSON contenant l'erreur
// TODO : voir pour ajouter la stack dans la réponse ?
func errorResponse(e error, errorCode int, w http.ResponseWriter) {
	log.Print(e)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(HTTPerror{
		Code:    int64(errorCode),
		Message: e.Error(),
	})
}
