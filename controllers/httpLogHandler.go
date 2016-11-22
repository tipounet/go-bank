package controllers

import (
	"log"
	"os"
	"strings"
	"time"

	restful "github.com/emicklei/go-restful"
)

// logger : permet de définir son propore logger, par exemple pour écrire dans un fichier (en changeant le stdout je suppose)
// type *log.Logger
var logger = log.New(os.Stdout, "", 0)

// NCSACommonLogFormatLogger : formater de la log de façon standard from
func NCSACommonLogFormatLogger() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		// voir ce que l'on peux faire de cela
		var username = "-"
		if req.Request.URL.User != nil {
			if name := req.Request.URL.User.Username(); name != "" {
				username = name
			}
		}
		chain.ProcessFilter(req, resp)
		// TODO : utiliser le log par défaut plutôt qu'un implementation  qui ne me servira pas ?
		logger.Printf("%s - %s [%s] \"%s %s %s\" %d %d",
			strings.Split(req.Request.RemoteAddr, ":")[0],
			username,
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			req.Request.Method,
			req.Request.URL.RequestURI(),
			req.Request.Proto,
			resp.StatusCode(),
			resp.ContentLength(),
		)
	}
}
