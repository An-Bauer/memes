package server

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func initServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{key}", handleRoot)
	mux.HandleFunc("GET /img/{key}", handleImage)

	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("handle root %s", r.PathValue("key"))
	path := "E:/InProgress/memes/web/index.html"

	http.ServeFile(w, r, path)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	//check first
	fmt.Printf("handle Image %s", r.PathValue("key"))
	path := filepath.Join("E:/InProgress/memes/images", r.PathValue("key")+".png")

	http.ServeFile(w, r, path)

}
