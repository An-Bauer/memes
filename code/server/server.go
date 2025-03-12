package server

import (
	"fmt"
	"memes/code/db"
	"net/http"
)

func runServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("GET /{key}", handleKey)
	mux.HandleFunc("GET /img/{key}", handleImage)

	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle root")
	path := "E:/InProgress/memes/web/index.html"

	http.ServeFile(w, r, path)
}

func handleKey(w http.ResponseWriter, r *http.Request) {
	//check first
	fmt.Printf("handle key:%s \n", r.PathValue("key"))
	path := "E:/InProgress/memes/web/index.html"

	http.ServeFile(w, r, path)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	fmt.Printf("handle Image key:%s \n", key)

	available, err := db.GetAvailability(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path, err := safeJoin("E:/InProgress/memes/images", r.PathValue("key")+".png")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !available {
		fmt.Fprint(w, "content not available")
		return
	}

	http.ServeFile(w, r, path)
}
