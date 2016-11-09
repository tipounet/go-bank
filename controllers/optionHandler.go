package controllers

import "net/http"

type optionInfo struct {
	Request string
}

// optionsHandler : réponse générique aux requêtes de type option, un jourune description en JSON pour décrire le service
func optionsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Allow", "GET, POST, PUT, DELETE, OPTIONS")
			writeHTTPJSONResponse(w, optionInfo{
				"osef je sais pas ce qu'il y a dans cette putain de request, plus tard on fournit une réponse sur tout ce que l'on fait",
			})
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
func optionsRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET, POST, PUT, DELETE, OPTIONS")
	writeHTTPJSONResponse(w, optionInfo{
		"osef",
	})
}
