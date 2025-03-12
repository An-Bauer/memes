package server

import (
	"fmt"
	"net/http"
)

func handleMemeGet(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	fmt.Printf("get meme %s\n", key)

	http.ServeFile(w, r, "E:/InProgress/memes/web/upload.html")
}

func handleMemePost(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	fmt.Println(r.PostFormValue("url"))
	fmt.Fprintf(w, "thanks for submitting %s ", key)
}
