package server

import (
	"fmt"
	"memes/code/db"
	"net/http"
	"path/filepath"
)

func runServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("GET /{key}", handleMeme)
	mux.HandleFunc("GET /img/{key}", handleImage)
	mux.HandleFunc("GET /favicon.ico", handleIcon)

	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle root")

	fmt.Fprint(w, "main page")
}

func handleMeme(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	fmt.Printf("handle meme (key:%s)\n", key)

	path := "E:/InProgress/memes/web/index.html"

	available, err := db.GetAvailability(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if !available {
		fmt.Fprint(w, "site not available")
		return
	}

	http.ServeFile(w, r, path)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	fmt.Printf("handle Image (key:%s) \n", key)

	base := "E:/InProgress/memes/images"
	path := filepath.Join(base, r.PathValue("key")+".png")
	if err := checkPath(base, path); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	available, err := db.GetAvailability(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if !available {
		fmt.Fprint(w, "content not available")
		fmt.Printf("tried to get unavailable image key:%s \n", key)
		return
	}

	http.ServeFile(w, r, path)
}

func handleIcon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle icon")
	w.WriteHeader(http.StatusBadRequest)
}
