package music

import (
	"github.com/gorilla/sessions"
	"github.com/zmb3/spotify"
)

// SessionMusicAuthenticaton is an interface for authenticating music service via session.
type SessionMusicAuthenticaton interface {
	GetAuth(*sessions.Session) (Authenticator, error)
}

// Authenticator is an object that contains information abour authentication to the music service.
type Authenticator interface{}

// Service is an inteface for connecting with music services in a unified way.
type Service interface {
	Authenticate(authenticator Authenticator) (Client, error)
}

// Client is an interface to work with music services.
type Client interface {
	CurrentUsersPlaylists() (*spotify.SimplePlaylistPage, error)
}
