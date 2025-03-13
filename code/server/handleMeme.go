package server

import (
	"fmt"
	"memes/code/db"
	"memes/code/imgflip"
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
	case 0:
		fmt.Printf("LOG: serving upload (key:%s)\n", key)
		http.ServeFile(w, r, "E:/InProgress/memes/web/upload.html")
	case 1:
		fmt.Printf("LOG: serving meme (key:%s)\n", key)
		http.ServeFile(w, r, "E:/InProgress/memes/web/meme.html")
	case 2:
		fmt.Printf("LOG: serving blocked (key:%s)\n", key)
		http.ServeFile(w, r, "E:/InProgress/memes/web/blocked.html")
	default:
		fmt.Printf("ERROR: got invalid status while serving meme (key:%s, status:%d)\n", key, status)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func handleMemePost(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	url := r.PostFormValue("url")

	status, err := db.GetStatus(key)
	if err != nil {
		fmt.Printf("SUS: invalid status request while posting meme (key:%s, url:%s, error:%v)\n", key, url, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch status {
	case 0:
		imgflipKey := url

		if err := imgflip.DownloadImag(imgflipKey); err != nil {
			fmt.Printf("ERROR: downloading failed (key:%s, url:%s, error:%v)\n", key, url, err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		if err := db.UpdateStatus(key, 1); err != nil {
			fmt.Printf("ERROR: updating meme status after upload failed (key:%s, url:%s, error:%v)\n", key, url, err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Printf("LOG: handeling upload (key:%s, url:%s)\n", key, url)
		http.ServeFile(w, r, "E:/InProgress/memes/web/uploadSuccessful.html")
	case 1:
		fmt.Printf("SUS: someone tried to post to baked key (key:%s, url:%s)\n", key, url)
		w.WriteHeader(http.StatusBadRequest)
	case 2:
		fmt.Printf("SUS: someone tried to post to blocked key (key:%s, url:%s)\n", key, url)
		w.WriteHeader(http.StatusBadRequest)
	default:
		fmt.Printf("ERROR: got invalid status while posting (key:%s, url:%s,status:%d)\n", key, url, status)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
