package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

// Session is a structure for creating handlers with session.
type Session struct {
	session *sessions.Session
}

// GetSession return session instance.
func (s *Session) GetSession() *sessions.Session {
	// TODO: make validation if session is initialized
	return s.session
}

// InitSession method initialize sesstion objects within the handler
func (s *Session) InitSession(r *http.Request) {
	s.session = InitSession(r)
}

var sessionStore = createStore()

func createStore() *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(os.Getenv("SECRET_KEY")))
	store.Options.HttpOnly = true
	store.Options.Path = "/"

	return store
}

// InitSession creates new session in the store and return it.
func InitSession(r *http.Request) *sessions.Session {
	session, err := sessionStore.Get(r, "chlorine_session")
	if err != nil {
		panic(err)
	}
	return session
}
