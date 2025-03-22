package server

import (
	"fmt"
	"memes/code/db"
	"memes/code/users"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("LOG: served homepage (url:%s)\n", r.URL.Path)
	http.ServeFile(w, r, "E:/InProgress/memes/web/home.html")
}

func handleIcon(w http.ResponseWriter, r *http.Request) {
	fmt.Print("LOG: served icon\n")
}

func handleStatic(w http.ResponseWriter, r *http.Request) {
	valid, username, err := users.CheckToken(r)
	fmt.Println(valid, username, err)

	file := r.PathValue("file")
	fmt.Printf("LOG: handeling static (file:%s)\n", file)

	secureFileServer(w, r, "E:/InProgress/memes/web", file)
}

//func handleCSS(w http.ResponseWriter, r *http.Request) {
//fmt.Println("LOG: served css")
//http.ServeFile(w, r, "E:/InProgress/memes/web/output.css")
//}

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
		fmt.Printf("SUS: someone tried to access blank image (key:%s)\n", key)
		fmt.Fprint(w, "meme not yet uploaded")
	case 1:
		fmt.Printf("LOG: served image (key:%s)\n", key)

		base := "E:/InProgress/memes/images"
		secureFileServer(w, r, base, key+".jpg")
	case 2:
		fmt.Printf("SUS: someone tried to access blocked image (key:%s)\n", key)
		w.WriteHeader(http.StatusBadRequest)
	default:
		fmt.Printf("ERROR: got invalid status while serving image (key:%s,status:%d)\n", key, status)
		w.WriteHeader(http.StatusBadRequest)
	}

}
