package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/cl"
	"chlorine/storage"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// MemberHandler serve endpoint for creating non-admin member for Chlorine.
type MemberHandler struct {
	auth.Session
	StorageHandler
}

func (h MemberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		jsonWriter.WriteJSONObject(member)
		return
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Printf("server: MemberHandler: error reading requset body: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		memberData := &struct {
			Name   string `json:"name,omitempty"`
			RoomID int    `json:"room_id,omitempty"`
		}{}
		err = json.Unmarshal(body, &memberData)
		if err != nil {
			log.Printf("server: MemberHandler: json parse error: %s", err)
			jsonWriter.Error(apierror.APIInvalidRequest, http.StatusBadRequest)
			return
		}
		member, err := cl.CreateMember(memberData.Name, memberData.RoomID, storage.RoleMember, h.storage)
		if err != nil {
			log.Printf("server: MemberHandler: cannot create member: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		token, err := h.storage.GetRoomToken(storage.ID(memberData.RoomID))
		if err != nil {
			log.Printf("server: MemberHandler: cannot retrieve token: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		oauthToken, err := token.ToOAuthToken()
		if err != nil {
			log.Printf("server: MemberHandler: cannot convert token: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		auth.WriteTokenToSession(session, oauthToken)
		session.Values["MemberID"] = member.ID
		err = session.Save(r, w)
		if err != nil {
			log.Printf("server: MemberHandler: error saving session: %s", err)
		}
		jsonWriter.WriteHeader(http.StatusCreated)
		return
	}
	jsonWriter.WriteHeader(http.StatusMethodNotAllowed)
}
