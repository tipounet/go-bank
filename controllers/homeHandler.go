package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/shurcooL/github_flavored_markdown"
)

// HomePage : page de garde
func HomePage(w http.ResponseWriter, r *http.Request) {
	initUserService()

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	data, err := ioutil.ReadFile("README.md")
	if err != nil {
		panic(err)
	} else {
		// b := bytes.NewBufferString(data)
		// b.Bytes()
		// TODO format ?
		// output := blackfriday.MarkdownCommon(data)
		output := github_flavored_markdown.Markdown(data)
		w.Write(bytes.NewBufferString("<html><head><meta charset=\"utf-8\"><link href=\"/assets/gfm.css\" media=\"all\" rel=\"stylesheet\" type=\"text/css\" /><link href=\"//cdnjs.cloudflare.com/ajax/libs/octicons/2.1.2/octicons.css\" media=\"all\" rel=\"stylesheet\" type=\"text/css\" /></head><body><article class=\"markdown-body entry-content\" style=\"padding: 30px;\">").Bytes())
		w.Write(output)
		w.Write(bytes.NewBufferString("</article></body></html>").Bytes())
	}
}

// Gfmstyle : refourgue la css Gfmstyle
// func Gfmstyle(w http.ResponseWriter, r *http.Request) {
// 	gfmstyle.Assets
// }
