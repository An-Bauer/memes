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
		return
	}

	err = users.HandleNewToken(username, w)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Printf("LOG: handeled register (username:%s, password:%s)\n", username, password)
	fmt.Fprint(w, "jo") // important
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("user")
	password := r.PostFormValue("password")

	exists, err := db.CheckUserExistance(username)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}
	if !exists {
		fmt.Fprint(w, "user doesn't exist")
		fmt.Printf("DEBUG: user doesn't exists")
		return
	}

	err = users.LoginUser(username, password)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	err = users.HandleNewToken(username, w)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	fmt.Printf("LOG: handeled login (username:%s)\n", username)
	fmt.Fprint(w, "test") // important
}
