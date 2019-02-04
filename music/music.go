package music

import (
	"chlorine/auth"

	"github.com/zmb3/spotify"
)

// Service is an inteface for connecting with music services in a unified way.
type Service interface {
	Authenticate(authenticator auth.Authenticator) (Client, error)
}

// Client is an interface to work with music services.
type Client interface {
	CurrentUsersPlaylists() (*spotify.SimplePlaylistPage, error)
	PlayerDevices() ([]spotify.PlayerDevice, error)
	PlayerState() (*spotify.PlayerState, error)
	TransferPlayback(spotify.ID, bool) error
}
