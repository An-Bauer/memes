package server

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles(
	"E:/InProgress/memes/web/meme.html",
))

func RunServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)               // homepage
	mux.HandleFunc("GET /{key}", handleMemeGet)   // upload page | meme page | blocked page
	mux.HandleFunc("POST /{key}", handleMemePost) // upload imgflip url -> meme page | error

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("E:/InProgress/memes/web"))))
	mux.HandleFunc("/static/output.css", handleCSS)
	mux.HandleFunc("GET /img/{key}", handleImage)  // serve image | errror
	mux.HandleFunc("GET /favicon.ico", handleIcon) // icon

	http.ListenAndServe("192.168.178.53:8", mux)
}

//func handleMeme(w http.ResponseWriter, r *http.Request) {
//key := r.PathValue("key")
//fmt.Printf("handle meme (key:%s)\n", key)

//path := "E:/InProgress/memes/web/index.html"

//available, err := db.GetAvailability(key)
//if err != nil {
//w.WriteHeader(http.StatusBadRequest)
//fmt.Println(err)
//return
//}
//if !available {
//fmt.Fprint(w, "site not available")
//return
//}

//http.ServeFile(w, r, path)
//}

//func handleUpload(w http.ResponseWriter, r *http.Request) {
//fmt.Println(r.PostFormValue("url"))
//}

//func handlePostUpload(w http.ResponseWriter, r *http.Request) {
//fmt.Fprint(w, "danke!")
//}
