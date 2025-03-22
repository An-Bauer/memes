package server

import (
	"fmt"
	"memes/code/db"
	"net/http"
)

func handleMemeGet(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	status, err := db.GetStatus(key)
	if err != nil {
		fmt.Printf("SUS: invalid status request while getting meme (key:%s,error:%v)\n", key, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch status {
	case 0: // blank key
		fmt.Printf("LOG: serving upload (key:%s)\n", key)
		http.ServeFile(w, r, "E:/InProgress/memes/web/upload.html")
	case 1: // baked key
		fmt.Printf("LOG: serving meme (key:%s)\n", key)
		serveTemplate(w, "meme3", key)
	case 2: // blocked key
		fmt.Printf("LOG: serving blocked (key:%s)\n", key)
		http.ServeFile(w, r, "E:/InProgress/memes/web/blocked.html")
	default:
		fmt.Printf("ERROR: got invalid status while serving meme (key:%s, status:%d)\n", key, status)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
