package server

import (
	"fmt"
	"memes/code/db"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("LOG: served homepage (url:%s)\n", r.URL.Path)
	http.ServeFile(w, r, "E:/InProgress/memes/web/home.html")
}

func handleIcon(w http.ResponseWriter, r *http.Request) {
	fmt.Print("LOG: served icon\n")
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	status, err := db.GetStatus(key)
	if err != nil {
		fmt.Printf("SUS: invalid Status request while serving image (key:%s,error:%v)", key, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch status {
	case 0:
		fmt.Printf("SUS: someone tried to access blank meme image (key:%s)\n", key)
		fmt.Fprint(w, "meme not yet uploaded")
	case 1:
		fmt.Printf("LOG: served image (key:%s)\n", key)

		base := "E:/InProgress/memes/images"
		secureFileServer(w, r, base, key+".png")
	case 2:
		fmt.Printf("SUS: someone tried to access blocked meme image (key:%s)\n", key)
		w.WriteHeader(http.StatusBadRequest)
	default:
		fmt.Printf("ERROR: got invalid status while serving image (key:%s,status:%d)\n", key, status)
		w.WriteHeader(http.StatusBadRequest)
	}

}
