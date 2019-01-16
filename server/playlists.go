package server

import (
	"akovalyov/chlorine/auth"
	"encoding/json"
	"log"
	"net/http"
)

// MyPlaylists returns list of personal playlists
func MyPlaylists(w http.ResponseWriter, r *http.Request) {
	session := auth.InitSession(r)
	if session.IsNew {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	authenticator := auth.GetSpotifyAuthenticator()

	token, err := auth.GetTokenFromSession(session)
	if err != nil {
		log.Printf("server: MyPlaylists: error retrieving token from session: %s", err)
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	client := authenticator.NewClient(token)
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Printf("server: MyPlaylists: cannot get top tracks: %s", err)
		http.Error(w, "Error retrieving songs.", http.StatusForbidden)
		return
	}
	serializedPlaylists, err := json.Marshal(playlists)
	if err != nil {
		log.Printf("server: MyPlaylists: %s", err)
		http.Error(w, "Error processing songs.", http.StatusForbidden)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(serializedPlaylists)
}
