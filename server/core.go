package server

import (
	"akovalyov/chlorine/auth"
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/zmb3/spotify"
)

// SessionedHandler implements Handler interface and compose Session fields and methods
type SessionedHandler struct {
	http.Handler
	auth.Session
}

// StartChlorineServer starts Chlorine to listen to HTTP connections on the given port.
func StartChlorineServer(port string) {
	handler := GetApplicationHandler()
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}

// InitSpotifyClientFromSession doing client initialization from session storage.
func InitSpotifyClientFromSession(s *sessions.Session) (*spotify.Client, error) {
	authenticator := auth.GetSpotifyAuthenticator()
	token, err := auth.GetTokenFromSession(s)
	if err != nil {
		return nil, err
	}
	client := authenticator.NewClient(token)
	return &client, nil
}

func init() {
	gob.Register(&time.Time{})
	gob.Register(&time.Location{})
}
