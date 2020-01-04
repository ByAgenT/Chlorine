package auth

import (
	"chlorine/cl"
	"chlorine/storage"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const (
	// AuthCallback used by Spotify OAuth for complete authorization flow and receive token.
	AuthCallback = "http://localhost/authcomplete"
)

var (
	scopes = []string{"streaming",
		spotify.ScopeUserLibraryModify,
		spotify.ScopeUserReadEmail,
		spotify.ScopeUserReadPrivate,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopePlaylistModifyPublic,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserModifyPlaybackState}
	secretKey = os.Getenv("SECRET_KEY")
)

// Authenticator is an object that contains information about authentication to the music service.
type Authenticator interface{}

// GetSpotifyAuthenticator create new instance of Spotify Authenticator
func GetSpotifyAuthenticator() spotify.Authenticator {
	return spotify.NewAuthenticator(AuthCallback, scopes...)
}

// ProcessReceivedToken gets OAuth Token from the callback request.
func ProcessReceivedToken(r *http.Request, s *sessions.Session) (*oauth2.Token, error) {
	authenticator := GetSpotifyAuthenticator()
	state, ok := s.Values["CSRFState"].(string)
	if !ok {
		return nil, errors.New("auth: cannot receive state from session")
	}
	token, err := authenticator.Token(state, r)
	if err != nil {
		return nil, err
	}
	if st := r.FormValue("state"); st != state {
		return nil, fmt.Errorf("auth: possible CSRF detected: state mismatch")
	}
	return token, nil
}

// GetTokenCode pulls an authorization code from an HTTP request and returns it as a string.
func GetTokenCode(state string, r *http.Request) (string, error) {
	values := r.URL.Query()
	if e := values.Get("error"); e != "" {
		return "", errors.New("spotify: auth failed - " + e)
	}
	code := values.Get("code")
	if code == "" {
		return "", errors.New("spotify: didn't get access code")
	}
	actualState := values.Get("state")
	if actualState != state {
		return "", errors.New("spotify: redirect state parameter doesn't match")
	}
	return code, nil
}

// CreateRandomState generates random state that is used for CSRF protection when authorizing via OAuth2.
func CreateRandomState(session *sessions.Session) string {
	preState := []byte(session.ID + secretKey)
	state := md5.Sum(preState)
	return hex.EncodeToString(state[:md5.Size])
}

// GetTokenFromSession pull authorization information from user session and return OAuth2 token.
func GetTokenFromSession(session *sessions.Session) (*oauth2.Token, error) {

	token := new(oauth2.Token)

	accessToken, ok := session.Values["AccessToken"].(string)
	if !ok {
		return nil, errors.New("auth: cannot retrieve access token from session")
	}
	expiryValue, ok := session.Values["Expiry"]
	if !ok {
		return nil, errors.New("auth: cannot retrieve token expiration from session")
	}
	expiry := &time.Time{}
	expiry, ok = expiryValue.(*time.Time)
	if !ok {
		return nil, errors.New("auth: cannot process token expiration from session")
	}
	refreshToken, ok := session.Values["RefreshToken"].(string)
	if !ok {
		return nil, errors.New("auth: cannot retrieve refresh token from session")
	}
	tokenType, ok := session.Values["TokenType"].(string)
	if !ok {
		return nil, errors.New("auth: cannot retrieve token type from session")
	}

	token.AccessToken = accessToken
	token.Expiry = *expiry
	token.RefreshToken = refreshToken
	token.TokenType = tokenType

	return token, nil
}

// WriteTokenToSession writes token information to the session key-value storage.
func WriteTokenToSession(session *sessions.Session, token *oauth2.Token) {
	session.Values["AccessToken"] = token.AccessToken
	session.Values["Expiry"] = &token.Expiry
	session.Values["RefreshToken"] = token.RefreshToken
	session.Values["TokenType"] = token.TokenType
}

// InitializeLogin setup initial login operations and return authentication URL.
func InitializeLogin(ctx context.Context, session *sessions.Session) string {
	authenticator := GetSpotifyAuthenticator()
	state := CreateRandomState(session)
	session.Values["CSRFState"] = state
	return authenticator.AuthURL(state)
}

// FinishAuthentication completes OAuth flow and save token information
func FinishAuthentication(ctx context.Context, r *http.Request, session *sessions.Session, memberService cl.MemberService, roomService cl.RoomService) error {
	token, err := ProcessReceivedToken(r, session)
	if err != nil {
		return fmt.Errorf("authentication: %s", err)
	}

	WriteTokenToSession(session, token)

	spotifyToken := &storage.Token{
		AccessToken:  token.AccessToken,
		Expiry:       token.Expiry,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType}

	roomConfig := &storage.RoomConfig{
		SongsPerMember: 5,
		MaxMembers:     10}

	room, err := roomService.CreateRoom(spotifyToken, roomConfig)
	if err != nil {
		log.Printf("authentication: %s", err)
		return fmt.Errorf("authentication: %s", err)
	}
	member, err := memberService.CreateMember(cl.RawMember{
		Name:   "Host",
		RoomID: int(*room.ID),
		Role:   storage.RoleAdmin,
	})
	if err != nil {
		log.Printf("authentication: %s", err)
		return fmt.Errorf("authentication: %s", err)
	}
	session.Values["MemberID"] = int(*member.ID)

	return nil
}
