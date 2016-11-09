package controllers

import (
	"go/build"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shurcooL/github_flavored_markdown/gfmstyle"
)

// NewRouter : créé un router a partir des routes que l'on met dans le tableau route le fichier routes.go . cela devrait permettre de mettre les routes en conf ? ou pas
// TODO : avoir un truc plus générique que les deux boucles for !!!
// les subrouteur permettent d'avoir un routeur lié à un préfix et les routes s'ajoute dessus
// exemple : https://gist.githubusercontent.com/danesparza/eb3a63ab55a7cd33923e/raw/f3e0be8a7cdb8779a3b109618b9f9c73523978fd/main.go
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	addRoute(router, getRouteWithAuth(), []func(http.Handler) http.Handler{jwtHandler})
	addRoute(router, getRouteWithoutAuth(), []func(http.Handler) http.Handler{})

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(gfmstyle.Assets)))
	// FIXME : howto expose static file (js / html / css / etc...) for webapp ?
	// => dl du code source de gogs pour voir comment c'est fait leur bazard !
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(???)))
	getStaticPath()
	return router
}

// addRoute : ajoutes les routes au routeur avecs les handlers
func addRoute(router *mux.Router, routes []Route, handlers []func(http.Handler) http.Handler) {
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		// handlers = append(handlers, Logger(handler, route.Name))
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(use(handler, handlers))

		router.
			Methods(http.MethodOptions).
			Path(route.Pattern).
			Name(route.Name).
			Handler(optionsHandler(handler))
	}

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

// from https://gist.github.com/elithrar/21cb76b8e29398722532
func use(handler http.Handler, middleware []func(http.Handler) http.Handler) http.Handler {
	for _, m := range middleware {
		handler = m(handler)
	}
	return handler
}
