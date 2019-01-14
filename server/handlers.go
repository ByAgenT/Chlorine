package server

import (
	"akovalyov/chlorine/auth"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	authenticator := auth.GetSpotifyAuthenticator()
	session := initSession(r)
	session.Save(r, w)
	http.Redirect(w, r, authenticator.AuthURL(auth.CSRFState), http.StatusFound)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	// authenticator := auth.GetSpotifyAuthenticator()
	token, err := auth.ProcessReceivedToken(r)
	if err != nil {
		http.Error(w, "Token obtaining error", http.StatusForbidden)
		log.Printf("Token receive error: %s", err)
	}
	// client := authenticator.NewClient(token)
	// log.Printf("New client has been added: %#v", client)
	session := initSession(r)
	log.Printf("Token value in playlists: %#v", token)
	session.Values["token"] = token
	session.Save(r, w)

	http.Redirect(w, r, "me/playlists", http.StatusFound)
}

func getMyPlaylists(w http.ResponseWriter, r *http.Request) {
	session := initSession(r)
	val := session.Values["token"]
	log.Printf("Token value in playlists: %#v", val)
	var token = &oauth2.Token{}
	if token, ok := val.(*oauth2.Token); !ok {
		log.Printf("Failed to cast value to oauth token %#v", token)
		return
	}
	authenticator := auth.GetSpotifyAuthenticator()
	client := authenticator.NewClient(token)
	// log.Printf("New client has been added: %#v", client)
	log.Printf("Client: %#v", client)
	if session.IsNew {
		// log.Printf("Client: %#v", client)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Fatalf("Cannot get top tracks: %s", err)
		http.Error(w, "Error retrieving songs.", http.StatusExpectationFailed)
	}
	serializedPlaylists, err := json.Marshal(playlists)
	w.Header().Add("Content-Type", "application/json")
	w.Write(serializedPlaylists)
	log.Printf("Finished executing")
}
