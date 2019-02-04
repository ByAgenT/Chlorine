package server

import (
	"chlorine/auth"
	"chlorine/music"
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/zmb3/spotify"
)

// ExternalMusicHandler contains external MusicService and authentication provider for it to retrieve music information.
type ExternalMusicHandler struct {
	auth.Session
	MusicService           music.Service
	AuthenticationProvider auth.SessionAuthenticaton
}

// GetClient return authenticate music service and return client instance.
func (h ExternalMusicHandler) GetClient(session *sessions.Session) (music.Client, error) {
	authenticator, err := h.AuthenticationProvider.GetAuth(session)
	if err != nil {
		return nil, err
	}
	client, err := h.MusicService.Authenticate(authenticator)
	if err != nil {
		return nil, err
	}
	return client, nil
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
