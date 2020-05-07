package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zmb3/spotify"
)

// SpotifyPlayHandler serve play call for Spotify API.
type SpotifyPlayHandler struct {
	auth.Session
	ExternalMusicHandler
}

func (h SpotifyPlayHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	switch r.Method {
	case "POST":
		client, err := h.GetClient(session)
		if err != nil {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("server: SpotifyPlayHandler: error reading requset body: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		bodyData := &struct {
			URIs []string `json:"uris,omitempty"`
		}{}
		err = json.Unmarshal(body, &bodyData)
		if err != nil {
			log.Printf("server: SpotifyPlayHandler: json parse error: %s", err)
			jsonWriter.Error(apierror.APIInvalidRequest, http.StatusBadRequest)
			return
		}
		playUris := make([]spotify.URI, 0)
		for _, uri := range bodyData.URIs {
			playUris = append(playUris, spotify.URI(uri))
		}
		err = client.PlayOpt(&spotify.PlayOptions{URIs: playUris})
		if err != nil {
			log.Printf("server: SpotifyPlayHandler: error start playing: %s", err)
			jsonWriter.Error(apierror.APIInvalidRequest, http.StatusBadRequest)
			return
		}
		jsonWriter.WriteHeader(http.StatusOK)
		return
	}

}
