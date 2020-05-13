package server

import (
	"chlorine/apierror"
	"net/http"

	"github.com/zmb3/spotify"
)

// SearchSongHandler serves search within Spotify.
type SearchSongHandler struct {
	ExternalMusicHandler
}

func (h SearchSongHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	client, err := h.GetClient(session)
	if err != nil {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	queries, ok := r.URL.Query()["q"]
	if !ok || len(queries[0]) < 1 {
		jsonWriter.Error(apierror.APIInvalidRequest, http.StatusBadRequest)
		return
	}
	query := queries[0]
	result, err := client.Search(query, spotify.SearchTypeTrack)
	if err != nil {
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
	}
	jsonWriter.WriteJSONObject(result)

}
