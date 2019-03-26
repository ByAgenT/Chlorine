package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/cl"
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
		member, err := cl.CreateMember(memberData.Name, memberData.RoomID, false, h.storage)
		if err != nil {
			log.Printf("server: MemberHandler: cannot create member: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		session.Values["MemberID"] = member.ID
		session.Save(r, w)
		jsonWriter.WriteHeader(http.StatusCreated)
		return
	}
	jsonWriter.WriteHeader(http.StatusMethodNotAllowed)
}
