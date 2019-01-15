package server

import (
	"akovalyov/chlorine/auth"
	"encoding/json"
	"log"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	authenticator := auth.GetSpotifyAuthenticator()
	session := initSession(r)
	state := auth.CreateRandomState(session)
	session.Values["CSRFState"] = state
	err := session.Save(r, w)
	if err != nil {
		log.Printf("server: handleLogin: error saving session: %s", err)
	}
	http.Redirect(w, r, authenticator.AuthURL(state), http.StatusFound)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	session := initSession(r)

	token, err := auth.ProcessReceivedToken(r, session)
	if err != nil {
		log.Printf("server: completeAuth: token process error: %s", err)
		http.Error(w, "Authentication error", http.StatusForbidden)
	}

	auth.WriteTokenToSession(session, token)
	err = session.Save(r, w)
	if err != nil {
		log.Printf("server: completeAuth: error saving session: %s", err)

	}

	http.Redirect(w, r, "me/playlists", http.StatusFound)
}

func getMyPlaylists(w http.ResponseWriter, r *http.Request) {
	session := initSession(r)
	if session.IsNew {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	authenticator := auth.GetSpotifyAuthenticator()

	token, err := auth.GetTokenFromSession(session)
	if err != nil {
		log.Printf("server: getMyPlaylists: error retrieving token from session: %s", err)
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	client := authenticator.NewClient(token)
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Printf("server: getMyPlaylists: cannot get top tracks: %s", err)
		http.Error(w, "Error retrieving songs.", http.StatusForbidden)
	}
	serializedPlaylists, err := json.Marshal(playlists)
	if err != nil {
		log.Printf("server: getMyPlaylists: %s", err)
		http.Error(w, "Error processing songs.", http.StatusForbidden)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(serializedPlaylists)
}
