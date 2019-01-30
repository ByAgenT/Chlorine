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

func (s SpotifySessionAuthentication) GetAuth(session *sessions.Session) (Authenticator, error) {
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

// SpotifyClient is a client for interacting with Spotify music service.
type SpotifyClient struct {
	client *spotify.Client
}

func (s SpotifyService) Authenticate(authenticator Authenticator) (Client, error) {
	spotifyAuth := auth.GetSpotifyAuthenticator()
	oauthAuthenticator, ok := authenticator.(SpotifyAuthenticator)
	if !ok {
		return nil, fmt.Errorf("spotify: cannot process authentication")
	}
	client := spotifyAuth.NewClient(oauthAuthenticator.Token)
	spotifyClient := &SpotifyClient{client: &client}
	return spotifyClient, nil
}

func (c SpotifyClient) CurrentUsersPlaylists() (*spotify.SimplePlaylistPage, error) {
	playlist, err := c.client.CurrentUsersPlaylists()
	return playlist, err
}
