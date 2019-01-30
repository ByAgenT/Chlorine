package server

import (
	"chlorine/apierror"
	"chlorine/music"
	"log"
	"net/http"
)

// MyPlaylistsHandler is a handler for user's personal playlists in Spotify
type MyPlaylistsHandler struct {
	SessionedHandler
	MusicService           music.Service
	AuthenticationProvider music.SessionMusicAuthenticaton
}

func (h MyPlaylistsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	authenticator, err := h.AuthenticationProvider.GetAuth(h.GetSession())
	if err != nil {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
	}
	client, err := h.MusicService.Authenticate(authenticator)
	if err != nil {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
	}
	playlists, err := client.CurrentUsersPlaylists()
	if err != nil {
		log.Printf("server: MyPlaylistsHandler: cannot get top tracks: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	jsonWriter.WriteJSONObject(playlists)
}
