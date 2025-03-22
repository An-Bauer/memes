package users

import (
	"crypto/rand"
	"memes/code/db"
	"net/http"
)

func HandleNewToken(username string, w http.ResponseWriter) error {
	token := rand.Text()

	err := db.UpdateToken(username, token)
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
