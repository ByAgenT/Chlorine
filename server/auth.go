package server

import (
	"akovalyov/chlorine/apierror"
	"akovalyov/chlorine/auth"
	"log"
	"net/http"
)

// LoginHandler initiates Chlorine room and start OAuth2 authentication flow for Spotify.
type LoginHandler SessionedHandler

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.InitSession(r)

	authenticator := auth.GetSpotifyAuthenticator()
	state := auth.CreateRandomState(h.GetSession())
	h.GetSession().Values["CSRFState"] = state
	err := h.GetSession().Save(r, w)
	if err != nil {
		log.Printf("server: handleLogin: error saving session: %s", err)
	}
	http.Redirect(w, r, authenticator.AuthURL(state), http.StatusFound)
}

// CompleteAuthHandler receives result from Spotify authorization and finishes authentication flow.
type CompleteAuthHandler SessionedHandler

func (h CompleteAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.InitSession(r)

	token, err := auth.ProcessReceivedToken(r, h.GetSession())
	if err != nil {
		log.Printf("server: completeAuth: token process error: %s", err)
		http.Error(w, "Authentication error", http.StatusForbidden)
		return
	}

	auth.WriteTokenToSession(h.GetSession(), token)
	err = h.GetSession().Save(r, w)
	if err != nil {
		log.Printf("server: completeAuth: error saving session: %s", err)
	}

	http.Redirect(w, r, "me/playlists", http.StatusFound)
}

// SpotifyTokenHandler returns Spotify authentication token from authorized user.
type SpotifyTokenHandler SessionedHandler

func (h SpotifyTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	token, err := auth.GetTokenFromSession(h.GetSession())
	if err != nil {
		log.Printf("server: spotifyToken: error retrieving token from session: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}

	err = h.GetSession().Save(r, w)
	if err != nil {
		log.Printf("server: spotifyToken: cannot save session: %s", err)
	}

	jsonWriter.WriteJSONObject(token)
}
