package users

import (
	"memes/code/db"
	"memes/code/encode"
	"net/http"
)

func SetToken(username string, w http.ResponseWriter) error {
	token, err := encode.RandomeString(20)
	if err != nil {
		return err
	}

	err = db.UpdateToken(username, token)
	if err != nil {
		return err
	}

	cookieUsernam := http.Cookie{
		Name:  "username",
		Value: username,

		//for security
		//Domain:   "127.0.0.1:3000",
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   30 * 24 * 3600,
		//Secure:   true,
	}
	http.SetCookie(w, &cookieUsernam)

	cookieToken := http.Cookie{
		Name:  "token",
		Value: token,

		//for security
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   30 * 24 * 3600,
		//Secure:   true,
	}
	http.SetCookie(w, &cookieToken)

	return nil
}
