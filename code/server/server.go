package server

import (
	"fmt"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles(
	"E:/InProgress/memes/web/meme.html",
))

func RunServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)               // homepage
	mux.HandleFunc("GET /{key}", handleMemeGet)   // upload page | meme page | blocked page
	mux.HandleFunc("POST /{key}", handleMemePost) // upload imgflip url -> meme page | error

	mux.HandleFunc("GET /img/{key}", handleImage)  // serve image | errror
	mux.HandleFunc("GET /favicon.ico", handleIcon) // icon

	mux.HandleFunc("GET /web/{file}", handleStatic)
	mux.HandleFunc("POST /api/register", handleRegister)
	mux.HandleFunc("POST /api/login", handleLogin)
	//http.ListenAndServe("192.168.178.53:8", mux)
	http.ListenAndServe("127.0.0.1:3000", mux)
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	fmt.Printf("LOG: handeling static (file:%s, )\n", file)

	secureFileServer(w, r, "E:/InProgress/memes/web", file)
}
