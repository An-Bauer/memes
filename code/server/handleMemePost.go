package server

import (
	"fmt"
	"memes/code/db"
	"memes/code/imgflip"
	"net/http"
)

func handleMemePost(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	imgflipKey := r.PostFormValue("imgflipKey")

	status, err := db.GetStatus(key)
	if err != nil {
		fmt.Printf("SUS: invalid status request while posting meme (key:%s, imgflipKey:%s, error:%v)\n", key, imgflipKey, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch status {
	case 0: // blank key
		if err := imgflip.DownloadImag(imgflipKey, key); err != nil {
			fmt.Printf("ERROR: downloading failed (key:%s, imgflipKey:%s, error:%v)\n", key, imgflipKey, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := db.UpdateStatus(key, 1); err != nil {
			fmt.Printf("ERROR: updating meme status after upload failed (key:%s, imgflipKey:%s, error:%v)\n", key, imgflipKey, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Printf("LOG: handeling upload (key:%s, imgflipKey:%s)\n", key, imgflipKey)
		http.ServeFile(w, r, "E:/InProgress/memes/web/uploadSuccessful.html")
	case 1: // baked key
		fmt.Printf("SUS: someone tried to post to baked key (key:%s, imgflipKey:%s)\n", key, imgflipKey)
		w.WriteHeader(http.StatusBadRequest)
	case 2: // blocked key
		fmt.Printf("SUS: someone tried to post to blocked key (key:%s, imgflipKey:%s)\n", key, imgflipKey)
		w.WriteHeader(http.StatusBadRequest)
	default:
		fmt.Printf("ERROR: got invalid status while posting (key:%s, imgflipKey:%s,status:%d)\n", key, imgflipKey, status)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
