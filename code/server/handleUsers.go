package server

import (
	"fmt"
	"memes/code/db"
	"memes/code/users"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func handleRegister(w http.ResponseWriter, r *http.Request) {
	//todo: check username and pasword format
	username := r.PostFormValue("user")
	password := r.PostFormValue("password")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("ERROR: hashing password failed (username:%s, password:%s, error: %v)\n", username, password, err)
		return
	}

	err = db.AddUser(username, hash)
	if err != nil {
		fmt.Printf("ERROR: adding user failed (username:%s, password:%s, error: %v)\n", username, password, err)
		return
	}

	err = users.SetToken(username, w)
	if err != nil {
		fmt.Printf("ERROR: setting token failed at register (username:%s, password:%s, error: %v)\n", username, password, err)
		return
	}

	fmt.Printf("LOG: handeled register (username:%s, password:%s)\n", username, password)

	fmt.Fprint(w, "jo") // important
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("user")
	password := r.PostFormValue("password")

	hash, err := db.GetHash(username)
	if err != nil {
		fmt.Printf("ERROR: getting hash failed (username:%s, password:%s, error: %v)\n", username, password, err)
		return
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		fmt.Printf("ERROR: validating hash failed (username:%s, password:%s, error: %v)\n", username, password, err)
		return
	}

	err = users.SetToken(username, w)
	if err != nil {
		fmt.Printf("ERROR: setting token failed at login (username:%s, password:%s, error: %v)\n", username, password, err)
		return
	}

	fmt.Printf("LOG: handeled login (username:%s, password:%s)\n", username, password)
	fmt.Fprint(w, "test") // important
}
