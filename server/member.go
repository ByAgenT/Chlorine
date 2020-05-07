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
	MemberService cl.MemberService
	TokenService  cl.TokenService
}

func (h MemberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	switch r.Method {
	case "GET":
		memberID, ok := session.Values["MemberID"].(int)
		if !ok {
			jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		}
		member, err := h.MemberService.GetMember(memberID)
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
		member, err := h.MemberService.CreateMember(cl.RawMember{
			Name:   memberData.Name,
			RoomID: memberData.RoomID,
			Role:   storage.RoleMember,
		})
		if err != nil {
			log.Printf("server: MemberHandler: cannot create member: %s", err)
			jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
			return
		}
		token, err := h.TokenService.GetRoomToken(memberData.RoomID)
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
		session.Values["MemberID"] = int(*member.ID)
		err = session.Save(r, w)
		if err != nil {
			log.Printf("server: MemberHandler: error saving session: %s", err)
		}
		jsonWriter.WriteHeader(http.StatusCreated)
		return
	}
	jsonWriter.WriteHeader(http.StatusMethodNotAllowed)
}
