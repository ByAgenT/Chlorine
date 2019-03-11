package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/cl"
	"chlorine/storage"
	"log"
	"net/http"
)

// LoginHandler initiates Chlorine room and start OAuth2 authentication flow for Spotify.
type LoginHandler struct {
	auth.Session
	StorageHandler
}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)

	authenticator := auth.GetSpotifyAuthenticator()
	state := auth.CreateRandomState(session)
	session.Values["CSRFState"] = state
	err := session.Save(r, w)
	if err != nil {
		log.Printf("server: handleLogin: error saving session: %s", err)
	}
	http.Redirect(w, r, authenticator.AuthURL(state), http.StatusFound)
}

// CompleteAuthHandler receives result from Spotify authorization and finishes authentication flow.
type CompleteAuthHandler struct {
	auth.Session
	StorageHandler
}

func (h CompleteAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)

	token, err := auth.ProcessReceivedToken(r, session)
	if err != nil {
		log.Printf("server: completeAuth: token process error: %s", err)
		http.Error(w, "Authentication error", http.StatusForbidden)
		return
	}

	auth.WriteTokenToSession(session, token)

	spotifyToken := &storage.SpotifyToken{
		AccessToken:  token.AccessToken,
		Expiry:       token.Expiry,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType}

	roomConfig := &storage.RoomConfig{
		SongsPerMember: 5,
		MaxMembers:     10}

	_, err = cl.CreateRoom(spotifyToken, roomConfig, h.storage)
	if err != nil {
		log.Printf("server: completeAuth: %s", err)
	}

	err = session.Save(r, w)
	if err != nil {
		log.Printf("server: completeAuth: error saving session: %s", err)
	}

	http.Redirect(w, r, "me/playlists", http.StatusFound)
}

// SpotifyTokenHandler returns Spotify authentication token from authorized user.
type SpotifyTokenHandler struct {
	auth.Session
}

func (h SpotifyTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	token, err := auth.GetTokenFromSession(session)
	if err != nil {
		log.Printf("server: spotifyToken: error retrieving token from session: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}

	err = session.Save(r, w)
	if err != nil {
		log.Printf("server: spotifyToken: cannot save session: %s", err)
	}

	jsonWriter.WriteJSONObject(token)
}
