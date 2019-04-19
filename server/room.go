package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/storage"
	"log"
	"net/http"
)

// RoomHandler handle room creation and retrieving inforamtion about rooms.
type RoomHandler struct {
	auth.Session
	StorageHandler
}

func (h RoomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	switch r.Method {
	case "GET":
		memberID, ok := session.Values["MemberID"].(storage.ID)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		}
		member, err := h.storage.GetMember(memberID)
		if err != nil {
			log.Printf("server: MemberHandler: cannot retrieve member: %s", err)
		}
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
		if err != nil {
			log.Printf("server: RoomHandler: %s", err.Error())
			jsonWriter.Error(apierror.APIServerError, 500)
		}
		jsonWriter.WriteJSONObject(room)
		return
	}

}
