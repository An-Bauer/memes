package users

import (
	"fmt"
	"memes/code/db"
	"net/http"
)

// loggedIn, username, error
func CheckToken(r *http.Request) (bool, string, error) {
	usernameCookie, err := r.Cookie("username")
	if err != nil {
		fmt.Println("DEBUG: username not in cookie")
		return false, "", nil
	}
	username := usernameCookie.Value

	tokenCookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println("DEBUG: token not in cookie")
		return false, "", nil
	}
	token := tokenCookie.Value

	exists, err := db.CheckUserExistance(username)
	if err != nil {
		return false, "", err
	}
	if !exists {
		fmt.Printf("SUS: cookie of unknown user (username:%s)", username)
		return false, "", err
	}

	dbToken, err := db.GetToken(username)
	if err != nil {
		return false, "", fmt.Errorf("error while getting token (err:%v, username:%s)", err, username)
	}

	if dbToken != token {
		fmt.Println("DEBUG: wrong token", token, dbToken)
		return false, "", nil
	}

	return true, username, nil
}
