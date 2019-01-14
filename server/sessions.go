package server

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var sessionStore = createStore()

func createStore() *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	store.Options.HttpOnly = true
	store.Options.Path = "/"

	return store
}

func initSession(r *http.Request) *sessions.Session {
	session, err := sessionStore.Get(r, "chlorine_session")
	if err != nil {
		panic(err)
	}
	return session
}
