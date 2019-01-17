package server

import (
	"akovalyov/chlorine/auth"
	"log"
	"net/http"
)

// AvailableDevicesHandler is a handler of a list of user's available devices.
type AvailableDevicesHandler struct {
	Session
}

func (h AvailableDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.session = auth.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	authenticator := auth.GetSpotifyAuthenticator()

	token, err := auth.GetTokenFromSession(h.session)
	if err != nil {
		log.Printf("server: AvailableDevices: error retrieving token from session: %s", err)
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	client := authenticator.NewClient(token)
	devices, err := client.PlayerDevices()
	if err != nil {
		log.Printf("server: AvailableDevices: error retrieving devices from spotify: %s", err)
		http.Error(jsonWriter, "Cannot retrieve devices", http.StatusForbidden)
		return
	}
	jsonWriter.WriteJSONObject(devices)
}
