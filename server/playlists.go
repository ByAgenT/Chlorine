package server

import (
	"chlorine/apierror"
	"log"
	"net/http"
)

// MyPlaylistsHandler is a handler for user's personal playlists in Spotify
type MyPlaylistsHandler struct {
	ExternalMusicHandler
}

func (h MyPlaylistsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	client, err := h.GetClient(session)
	if err != nil {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Printf("server: MyPlaylistsHandler: cannot get top tracks: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	jsonWriter.WriteJSONObject(playlists)
}
