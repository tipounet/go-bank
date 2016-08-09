package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown/gfmstyle"
)

// NewRouter : créé un router a partir des routes que l'on met dans le tableau route en bas de ce fichier. cela va permttre de mettre les routes en conf ? ou pas
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(gfmstyle.Assets)))
	// http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(gfmstyle.Assets)))
	return router
}
