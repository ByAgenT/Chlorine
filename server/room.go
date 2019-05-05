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
			return
		}
		member, err := h.storage.GetMember(memberID)
		if err != nil {
			log.Printf("server: MemberHandler: cannot retrieve member: %s", err)
			return
		}
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
		if err != nil {
			log.Printf("server: RoomHandler: %s", err.Error())
			jsonWriter.Error(apierror.APIServerError, 500)
			return
		}
		jsonWriter.WriteJSONObject(room)
		return
	}

}

// RoomMembersHandler handle serving members of the current room
type RoomMembersHandler struct {
	auth.Session
	StorageHandler
}

func (h RoomMembersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	switch r.Method {
	case "GET":
		memberID, ok := session.Values["MemberID"].(storage.ID)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
			return
		}
		member, err := h.storage.GetMember(memberID)
		if err != nil {
			log.Printf("server: MemberHandler: cannot retrieve member: %s", err)
			return
		}
		room, err := h.storage.GetRoom(storage.ID(member.RoomID))
		if err != nil {
			log.Printf("server: RoomHandler: %s", err.Error())
			jsonWriter.Error(apierror.APIServerError, 500)
			return
		}
		members, err := room.GetMembers()
		if err != nil {
			log.Printf("server: RoomHandler: %s", err.Error())
			jsonWriter.Error(apierror.APIServerError, 500)
			return
		}
		jsonWriter.WriteJSONObject(members)
		return
	}
}
