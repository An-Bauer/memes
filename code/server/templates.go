package server

import (
	"net/http"
	"text/template"
)

type memePage struct {
	key  string
	user int
}

var templates = template.Must(template.ParseFiles(
	"E:/InProgress/memes/web/meme.html",
))

func serveTemplate(w http.ResponseWriter, tmpl, key string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
