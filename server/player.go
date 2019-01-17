package server

import (
	"akovalyov/chlorine/apierror"
	"akovalyov/chlorine/auth"
	"log"
	"net/http"
)

// AvailableDevicesHandler is a handler of a list of user's available devices.
type AvailableDevicesHandler struct {
	auth.Session
}

func (h AvailableDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	authenticator := auth.GetSpotifyAuthenticator()

	token, err := auth.GetTokenFromSession(h.GetSession())
	if err != nil {
		log.Printf("server: AvailableDevicesHandler: error retrieving token from session: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}

	client := authenticator.NewClient(token)
	devices, err := client.PlayerDevices()
	if err != nil {
		log.Printf("server: AvailableDevicesHandler: error retrieving devices from spotify: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}
	jsonWriter.WriteJSONObject(devices)
}
