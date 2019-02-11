package music

import (
	"chlorine/auth"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/zmb3/spotify"

	"golang.org/x/oauth2"
)

// SpotifySessionAuthentication provides authentication for the Spotify using information from the session.
type SpotifySessionAuthentication struct{}

// GetAuth returns SpotifyAuthenticator populates with OAuth token from session.
func (s SpotifySessionAuthentication) GetAuth(session *sessions.Session) (auth.Authenticator, error) {
	token, err := auth.GetTokenFromSession(session)
	if err != nil {
		return nil, err
	}
	return SpotifyAuthenticator{Token: token}, nil
}

// SpotifyAuthenticator is an implemetation of an authenticator for the Spotify.
type SpotifyAuthenticator struct {
	Token *oauth2.Token
}

// SpotifyService is a service for authenticating and work with Spotify music service.
type SpotifyService struct{}

// Authenticate provides authenticator for the Spotify and return Client for Spotify.
func (s SpotifyService) Authenticate(authenticator auth.Authenticator) (Client, error) {
	spotifyAuth := auth.GetSpotifyAuthenticator()
	oauthAuthenticator, ok := authenticator.(SpotifyAuthenticator)
	if !ok {
		return nil, fmt.Errorf("spotify: cannot process authentication")
	}
	client := spotifyAuth.NewClient(oauthAuthenticator.Token)
	spotifyClient := &SpotifyClient{client: &client}
	return spotifyClient, nil
}

// SpotifyClient is a client for interacting with Spotify music service.
type SpotifyClient struct {
	client *spotify.Client
}

// CurrentUsersPlaylists returns playlists from the Spotify API.
func (c SpotifyClient) CurrentUsersPlaylists() (*spotify.SimplePlaylistPage, error) {
	playlist, err := c.client.CurrentUsersPlaylists()
	return playlist, err
}

// PlayerDevices returns available devices for the playback.
func (c SpotifyClient) PlayerDevices() ([]spotify.PlayerDevice, error) {
	devices, err := c.client.PlayerDevices()
	return devices, err
}

// PlayerState returns current player state.
func (c SpotifyClient) PlayerState() (*spotify.PlayerState, error) {
	playerState, err := c.client.PlayerState()
	return playerState, err
}

// TransferPlayback transfer playback to another player by a ID.
func (c SpotifyClient) TransferPlayback(id spotify.ID, play bool) error {
	err := c.client.TransferPlayback(id, play)
	return err
}