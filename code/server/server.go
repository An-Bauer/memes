package server

import (
	"errors"
	"fmt"
	"log"
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

	mux.HandleFunc("GET /img/{key}", handleImage)  // serve image | errror
	mux.HandleFunc("GET /favicon.ico", handleIcon) // icon

	mux.HandleFunc("GET /web/{file}", handleStatic)
	mux.HandleFunc("POST /api/register", handleRegister)
	mux.HandleFunc("POST /api/login", handleLogin)
	//mux.HandleFunc("GET /api/cookieSet", handleSetCookie)
	//mux.HandleFunc("GET /api/cookieGet", handleGetCookie)

	//http.ListenAndServe("192.168.178.53:8", mux)
	err := http.ListenAndServe("127.0.0.1:3000", mux)
	if err != nil {
		fmt.Printf("ERROR: server error (error:%v)\n", err)
	}
}

func handleSetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "test",
		Value: "bla",

		//for security
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		//MaxAge:   30 * 24 * 3600,
		//Secure:   true,
	}

	http.SetCookie(w, &cookie)

	// Write a HTTP response as normal.
	fmt.Fprint(w, "test")
}

func handleGetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("test")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	fmt.Println("cookie: ", cookie.Value)
}
