package server

import (
	"akovalyov/chlorine/apierror"
	"akovalyov/chlorine/auth"
	"log"
	"net/http"
)

// MyPlaylistsHandler is a handler for user's personal playlists in Spotify
type MyPlaylistsHandler struct {
	auth.Session
}

func (h MyPlaylistsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
	authenticator := auth.GetSpotifyAuthenticator()

	token, err := auth.GetTokenFromSession(h.GetSession())
	if err != nil {
		log.Printf("server: MyPlaylistsHandler: error retrieving token from session: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	client := authenticator.NewClient(token)
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Printf("server: MyPlaylistsHandler: cannot get top tracks: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	jsonWriter.WriteJSONObject(playlists)
}
