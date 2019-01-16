package server

import (
	"akovalyov/chlorine/auth"
	"encoding/json"
	"log"
	"net/http"
)

// HandleLogin initiates Chlorine session and start OAuth2 authentication flow for Spotify.
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	authenticator := auth.GetSpotifyAuthenticator()
	session := auth.InitSession(r)
	state := auth.CreateRandomState(session)
	session.Values["CSRFState"] = state
	err := session.Save(r, w)
	if err != nil {
		log.Printf("server: handleLogin: error saving session: %s", err)
	}
	http.Redirect(w, r, authenticator.AuthURL(state), http.StatusFound)
}

// CompleteAuth receives result from Spotify authorization and finishes authentication flow.
func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	session := auth.InitSession(r)

	token, err := auth.ProcessReceivedToken(r, session)
	if err != nil {
		log.Printf("server: completeAuth: token process error: %s", err)
		http.Error(w, "Authentication error", http.StatusForbidden)
		return
	}

	auth.WriteTokenToSession(session, token)
	err = session.Save(r, w)
	if err != nil {
		log.Printf("server: completeAuth: error saving session: %s", err)
	}

	http.Redirect(w, r, "me/playlists", http.StatusFound)
}

// SpotifyToken return Spotify authentication token from authorized user
func SpotifyToken(w http.ResponseWriter, r *http.Request) {
	session := auth.InitSession(r)
	token, err := auth.GetTokenFromSession(session)
	if err != nil {
		log.Printf("server: spotifyToken: error retrieving token from session: %s", err)
		http.Error(w, "Error retrieving token", http.StatusForbidden)
		return
	}

	serializedToken, err := json.Marshal(token)
	if err != nil {
		log.Printf("server: spotifyToken: cannot serialize token: %s", err)
		http.Error(w, "Error retrieving token.", http.StatusInternalServerError)
		return
	}

	err = session.Save(r, w)
	if err != nil {
		log.Printf("server: spotifyToken: cannot save session: %s", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(serializedToken)
}
