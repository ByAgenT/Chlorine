package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"context"
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
	jsonWriter := JSONResponseWriter{w}

	authURL := auth.InitializeLogin(context.Background(), session)
	err := session.Save(r, w)
	if err != nil {
		log.Printf("unable to save session: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, authURL, http.StatusFound)
}

// CompleteAuthHandler receives result from Spotify authorization and finishes authentication flow.
type CompleteAuthHandler struct {
	auth.Session
	StorageHandler
}

func (h CompleteAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)

	err := auth.FinishAuthentication(context.Background(), r, session, h.storage)
	if err != nil {
		log.Printf("server: authentication finishing error: %s", err)
		http.Error(w, "Authentication error", http.StatusForbidden)
		return
	}
	err = session.Save(r, w)
	if err != nil {
		log.Printf("server: completeAuth: error saving session: %s", err)
	}
	http.Redirect(w, r, "/player", http.StatusFound)
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
