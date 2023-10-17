package utils

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("SuperSecretKey"))

func GetToken(r *http.Request) string {
	session, _ := Store.Get(r, "session")
	token, ok := session.Values["token"].(string)
	if ok {
		return token
	}
	return ""
}
