package server

import (
	"akovalyov/chlorine/auth"
	"log"
	"net/http"
)

// LoginHandler initiates Chlorine room and start OAuth2 authentication flow for Spotify.
type LoginHandler struct {
	Session
}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.session = auth.InitSession(r)

	authenticator := auth.GetSpotifyAuthenticator()
	state := auth.CreateRandomState(h.session)
	h.session.Values["CSRFState"] = state
	err := h.session.Save(r, w)
	if err != nil {
		log.Printf("server: handleLogin: error saving session: %s", err)
	}
	http.Redirect(w, r, authenticator.AuthURL(state), http.StatusFound)
}

// CompleteAuthHandler receives result from Spotify authorization and finishes authentication flow.
type CompleteAuthHandler struct {
	Session
}

func (h CompleteAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.session = auth.InitSession(r)

	token, err := auth.ProcessReceivedToken(r, h.session)
	if err != nil {
		log.Printf("server: completeAuth: token process error: %s", err)
		http.Error(w, "Authentication error", http.StatusForbidden)
		return
	}

	auth.WriteTokenToSession(h.session, token)
	err = h.session.Save(r, w)
	if err != nil {
		log.Printf("server: completeAuth: error saving session: %s", err)
	}

	http.Redirect(w, r, "me/playlists", http.StatusFound)
}

// SpotifyTokenHandler returns Spotify authentication token from authorized user.
type SpotifyTokenHandler struct {
	Session
}

func (h SpotifyTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.session = auth.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	token, err := auth.GetTokenFromSession(h.session)
	if err != nil {
		log.Printf("server: spotifyToken: error retrieving token from session: %s", err)
		http.Error(w, "Error retrieving token", http.StatusForbidden)
		return
	}

	err = h.session.Save(r, w)
	if err != nil {
		log.Printf("server: spotifyToken: cannot save session: %s", err)
	}

	jsonWriter.WriteJSONObject(token)
}
