package server

import (
	"akovalyov/chlorine/auth"
	"log"
	"net/http"
)

// MyPlaylistsHandler is a handler for user's personal playlists in Spotify
type MyPlaylistsHandler struct {
	Session
}

func (h MyPlaylistsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.session = auth.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	if h.session.IsNew {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	authenticator := auth.GetSpotifyAuthenticator()

	token, err := auth.GetTokenFromSession(h.session)
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
	jsonWriter.WriteJSONObject(playlists)
}
