package controllers

import (
	"go/build"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown/gfmstyle"
)

// NewRouter : créé un router a partir des routes que l'on met dans le tableau route le fichier routes.go . cela devrzit permettre de mettre les routes en conf ? ou pas
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range getRoute() {
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
	// FIXME : howto expose static file (js / html / css / etc...) for webapp ?
	// => dl du code source de gogs pour voir comment c'est fait leur bazard !
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(???)))
	getStaticPath()
	return router
}

func getStaticPath() string {
	p, err := build.Import("github.com/tipounet/go-bank", "", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("type de p : %T", p)
	h := http.Dir(p.Dir)
	log.Printf("Type de h : %T", h)
	return ""
}
