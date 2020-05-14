package server

import (
	"chlorine/auth"
	"chlorine/music"
	"github.com/gorilla/sessions"
)

// ExternalMusicHandler contains external MusicService and authentication provider for it to retrieve music information.
type ExternalMusicHandler struct {
	auth.Session
	MusicService           music.Service
	AuthenticationProvider auth.SessionAuthentication
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
