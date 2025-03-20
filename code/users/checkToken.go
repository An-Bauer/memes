package users

import (
	"errors"
	"fmt"
	"log"
	"memes/code/db"
	"net/http"
)

func CheckToken(r *http.Request) (bool, string, error) {
	usernameCookie, err := r.Cookie("username")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			fmt.Println("username cookie not found", err)
		default:
			log.Println(err)
		}
		return false, "", nil
	}

	tokenCookie, err := r.Cookie("token")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			fmt.Println("hash cookie not found", err)
		default:
			log.Println(err)
		}
		return false, "", nil
	}

	fmt.Printf("username:%s, token:%s\n", usernameCookie.Value, tokenCookie.Value)

	token, err := db.GetToken(usernameCookie.Value)
	if err != nil {
		return false, "", err
	}

	if token != tokenCookie.Value {
		fmt.Println("wrong token", token, tokenCookie.Value)
		return false, "", nil
	}

	return true, usernameCookie.Value, nil
}
