package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	sessionStore = createStore()
)

const (
	// DefaultSessionName is a default name for the cookie which store session.
	DefaultSessionName = "chlorine_session"
)

// Session is a structure for creating handlers with session.
type Session struct {
	session *sessions.Session
}

// SessionAuthenticaton is an interface for authenticating music service via session.
type SessionAuthenticaton interface {
	GetAuth(*sessions.Session) (Authenticator, error)
}

// GetSession return session instance.
func (s *Session) GetSession() (*sessions.Session, error) {
	if s.session == nil {
		return nil, fmt.Errorf("auth: session: session is not initialized")
	}
	return s.session, nil
}

// InitSession method initialize sesstion objects within the handler
func (s *Session) InitSession(r *http.Request) *sessions.Session {
	s.session = InitSession(r)
	return s.session
}

func createStore() *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(secretKey))
	store.Options.HttpOnly = true
	store.Options.Path = "/"
	store.Options.MaxAge = 0

	return store
}

// InitSession creates new session in the store and return it.
func InitSession(r *http.Request) *sessions.Session {
	session, err := sessionStore.Get(r, DefaultSessionName)
	if err != nil {
		log.Fatalf("auth: session: %s", err.Error())
	}
	return session
}

// SessionMiddleware is a middleware for initializing session between server and client.
func SessionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}
