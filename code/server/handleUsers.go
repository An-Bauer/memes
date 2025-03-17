package server

import (
	"fmt"
	"net/http"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("user")
	passwort := r.PostFormValue("password")

	fmt.Println("register", username, passwort)

}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("user")
	passwort := r.PostFormValue("password")

	fmt.Println("login", username, passwort)

}
