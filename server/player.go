package server

import (
	"akovalyov/chlorine/apierror"
	"log"
	"net/http"
)

// AvailableDevicesHandler is a handler of a list of user's available devices.
type AvailableDevicesHandler SessionedHandler

func (h AvailableDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	client, err := InitSpotifyClientFromSession(h.GetSession())
	if err != nil {
		log.Printf("server: AvailableDevicesHandler: error initializing Spotify client: %s", err)
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
type PlaybackHandler SessionedHandler

func (h PlaybackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
