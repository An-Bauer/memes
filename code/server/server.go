package server

import (
	"fmt"
	"memes/code/db"
	"net/http"
	"path/filepath"
)

func RunServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("GET /{key}", handleMemeGet)
	mux.HandleFunc("POST /{key}", handleMemePost)
	mux.HandleFunc("GET /favicon.ico", handleIcon)

	//mux.HandleFunc("GET /img/{key}", handleImage)
	//mux.HandleFunc("POST /upload", handleUpload)
	//mux.HandleFunc("GET /upload", handlePostUpload)

	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("handle root (%s)\n", r.URL.Path)

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

func handleUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.PostFormValue("url"))
}

func handlePostUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "danke!")
}
