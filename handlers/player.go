package handlers

import (
	"akovalyov/chlorine/server"
	"net/http"
)

func AvailableDevices(w http.ResponseWriter, r *http.Request) {
	jsonWriter := server.JSONResponseWriter{w}
	jsonWriter.WriteJSON([]byte("{}"))
}
