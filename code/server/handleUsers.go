package server

import (
	"fmt"
	"memes/code/db"
	"memes/code/users"
	"net/http"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {
	//todo: check username and pasword format
	username := r.PostFormValue("user")
	password := r.PostFormValue("password")

	exists, err := db.CheckUserExistance(username)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		fmt.Fprint(w, "user already exists")
		fmt.Printf("DEBUG: user already exists")
		return
	}

	err = users.RegisterUser(username, password)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = users.HandleNewToken(username, w)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("LOG: handeled register (username:%s, password:%s)\n", username, password)
	fmt.Fprint(w, "success")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("user")
	password := r.PostFormValue("password")

	exists, err := db.CheckUserExistance(username)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exists {
		fmt.Fprint(w, "user dosn't exist")
		fmt.Printf("DEBUG: user dosn't exist")
		return
	}

	correct, err := users.LoginUser(username, password)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !correct {
		fmt.Fprint(w, "wrong password")
		fmt.Printf("DEBUG: wrong password")
		return
	}

	err = users.HandleNewToken(username, w)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("LOG: handeled login (username:%s)\n", username)
	fmt.Fprint(w, "success")
}
