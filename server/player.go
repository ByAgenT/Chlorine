package server

import (
	"chlorine/apierror"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zmb3/spotify"
)

// AvailableDevicesHandler is a handler of a list of user's available devices.
type AvailableDevicesHandler struct {
	ExternalMusicHandler
}

func (h AvailableDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	client, err := h.GetClient(session)
	if err != nil {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	devices, err := client.PlayerDevices()
	if err != nil {
		log.Printf("server: AvailableDevicesHandler: error retrieving devices from spotify: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	jsonWriter.WriteJSONObject(devices)
}

// PlaybackHandler is a handler for Spotify playback actions.
type PlaybackHandler struct {
	ExternalMusicHandler
}

func (h PlaybackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	client, err := h.GetClient(session)
	if err != nil {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}

	switch r.Method {
	case "GET":
		playerState, err := client.PlayerState()
		if err != nil {
			log.Printf("server: PlaybackHandler: error retrieving PlayerState: %s", err)
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
			return
		}
		jsonWriter.WriteJSONObject(playerState)
	case "PUT":
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("server: PlaybackHandler: error reading requset body: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		parsedReq := &struct {
			DeviceID spotify.ID `json:"device_id"`
			Play     bool       `json:"play,omitempty"`
		}{}
		err = json.Unmarshal(body, &parsedReq)
		if err != nil {
			jsonWriter.Error(apierror.APIInvalidRequest, http.StatusBadRequest)
			return
		}
		log.Printf("Tranferring to device ID: %s", parsedReq.DeviceID)
		err = client.TransferPlayback(parsedReq.DeviceID, parsedReq.Play)
		if err != nil {
			log.Printf("server: PlaybackHandler: error transferring playback: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusBadRequest)
			return
		}
		jsonWriter.WriteHeader(http.StatusNoContent)
	}
}
