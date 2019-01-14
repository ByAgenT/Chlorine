package auth

import (
	"fmt"
	"net/http"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const (
	// AuthCallback used by Spotify OAuth for complete authorization flow and receive token.
	AuthCallback = "http://localhost:8080/authcomplete"
)

var (
	scopes    = []string{"streaming", spotify.ScopeUserReadBirthdate, spotify.ScopeUserReadEmail, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistReadPrivate}
	ch        = make(chan *spotify.Client)
	CSRFState = "7w7%^3tgrhku^@fbhe!"
)

// GetSpotifyAuthenticator create new instance of Spotify Authenticator
func GetSpotifyAuthenticator() spotify.Authenticator {
	return spotify.NewAuthenticator(AuthCallback, scopes...)
}

// ProcessReceivedToken gets OAuth Token from the callback request.
func ProcessReceivedToken(r *http.Request) (*oauth2.Token, error) {
	authenticator := GetSpotifyAuthenticator()
	token, err := authenticator.Token(CSRFState, r)
	if err != nil {
		return nil, err
	}
	if st := r.FormValue("state"); st != CSRFState {
		return nil, fmt.Errorf("Possible CSRF Detected: state mismatch")
	}
	return token, nil
}
