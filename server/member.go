package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/cl"
	"chlorine/storage"
	"chlorine/ws"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// MemberHandler serve endpoint for creating non-admin member for Chlorine.
type MemberHandler struct {
	auth.Session
	MemberService cl.MemberService
	TokenService  cl.TokenService
}

func (h MemberHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
	member, ok := auth.GetMemberIfAuthorized(h.MemberService, session)
	if !ok {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	jsonWriter.WriteJSONObject(member)
}

func (h MemberHandler) Post(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
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
	session.Values["MemberID"] = *member.ID
	err = session.Save(r, w)
	if err != nil {
		log.Printf("server: MemberHandler: error saving session: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
	}
	jsonWriter.WriteHeader(http.StatusCreated)
	ws.Broadcast(roomWSConnections[member.RoomID], &ws.Response{
		Type:        ws.TypeBroadcast,
		Status:      ws.StatusOK,
		Description: "MemberAdded",
		Body: map[string]interface{}{
			"member": member,
		},
	})
}

type MemberDetailHandler struct {
	auth.Session
	MemberService cl.MemberService
	TokenService  cl.TokenService
}

func (h MemberDetailHandler) Delete(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}
	member, ok := auth.GetMemberIfAuthorized(h.MemberService, session)
	if !(ok && auth.IsMemberAdministrator(memberService, member)) {
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusUnauthorized)
		return
	}
	memberID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Printf("server: MemberDetailHandler: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	err = h.MemberService.Delete(memberID, false)
	if err != nil {
		if errors.Is(err, cl.ErrorDeleteProtected) {
			jsonWriter.Error(apierror.APIError{
				Description: "Administrator cannot be deleted",
				ErrorCode:   apierror.StatusInvalidRequest,
			}, http.StatusBadRequest)
			return
		}
		log.Printf("server: MemberDetailHandler: %s", err)
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		return
	}
	jsonWriter.WriteHeader(http.StatusOK)
	ws.Broadcast(roomWSConnections[member.RoomID], &ws.Response{
		Type:        ws.TypeBroadcast,
		Status:      ws.StatusOK,
		Description: "MemberDeleted",
		Body: map[string]interface{}{
			"member": member,
		},
	})
}
