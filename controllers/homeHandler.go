package controllers

import (
	"bytes"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	restful "github.com/emicklei/go-restful"
	"github.com/shurcooL/github_flavored_markdown"
)

// HomePageResource : Ressoure au sens http d'un service rest
type HomePageResource struct{}

// RegisterTo : Permet l'enregistrement des Ressoures pour la version dans le container http global
func (h HomePageResource) RegisterTo() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/").
		Consumes("*/*").
		Produces("Content-type: text/html; charset=UTF-8")

	ws.Route(ws.GET("").To(h.HomePage))
	// route pour servir la css quiva avec
	ws.Route(ws.GET("/assets/{subpath:*}").To(staticFromPathParam))
	// TODO : get octicons.css for serve in app, for unconnected usage
	return ws
}

// staticFromPathParam : récupère les fichiers "static" depuis leurs chemins
func staticFromPathParam(req *restful.Request, resp *restful.Response) {
	rootdir := importPathToDir("github.com/shurcooL/github_flavored_markdown/gfmstyle/_data")
	actual := path.Join(rootdir, req.PathParameter("subpath"))
	fmt.Printf("serving %s ... (from %s)\n", actual, req.PathParameter("subpath"))
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actual)
}

// From github.com/shurcooL/github_flavored_markdown/gfmstyle/asserts.go
func importPathToDir(importPath string) string {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}
	return p.Dir
}

// HomePage : page de garde
func (h HomePageResource) HomePage(request *restful.Request, response *restful.Response) {
	// response.AddHeader("Content-Type", "text/html; charset=UTF-8")
	data, err := ioutil.ReadFile("README.md")
	if err != nil {
		panic(err)
	} else {
		output := github_flavored_markdown.Markdown(data)
		response.Write(bytes.NewBufferString("<html><head><meta charset=\"utf-8\"><link href=\"/assets/gfm.css\" media=\"all\" rel=\"stylesheet\" type=\"text/css\" /><link href=\"//cdnjs.cloudflare.com/ajax/libs/octicons/2.1.2/octicons.css\" media=\"all\" rel=\"stylesheet\" type=\"text/css\" /></head><body><article class=\"markdown-body entry-content\" style=\"padding: 30px;\">").Bytes())
		response.Write(output)
		response.Write(bytes.NewBufferString("</article></body></html>").Bytes())
	}
}
