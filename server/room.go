package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"log"
	"net/http"
)

// RoomHandler handle room creation and retrieving inforamtion about rooms.
type RoomHandler struct {
	auth.Session
	StorageHandler
}

func (h RoomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	rooms, err := h.storage.GetRooms()
	if err != nil {
		log.Printf("server: RoomHandler: %s", err.Error())
		jsonWriter.Error(apierror.APIServerError, 500)
	}
	jsonWriter.WriteJSONObject(rooms)
}
