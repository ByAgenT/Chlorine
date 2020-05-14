package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/cl"
	"chlorine/storage"
	"context"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

func getMemberIfAuthorized(service cl.MemberService,
	session *sessions.Session) (*storage.Member, bool) {
	memberID, ok := session.Values["MemberID"].(int)
	if !ok {
		return nil, false
	}
	member, err := service.GetMember(memberID)
	if err != nil {
		// TODO: return false only if NotFound error
		log.Printf("server: cannot get member for auth check: %s", err)
		return nil, false
	}
	return member, true
}

// LoginHandler initiates Chlorine room and start OAuth2 authentication flow for Spotify.
type LoginHandler struct {
	auth.Session
}

func (h LoginHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	authURL := auth.InitializeLogin(context.Background(), session)
	err := session.Save(r, w)
	panicIfErr(jsonWriter, err, "unable to save session")
	http.Redirect(w, r, authURL, http.StatusFound)
}

// CompleteAuthHandler receives result from Spotify authorization and finishes authentication flow.
type CompleteAuthHandler struct {
	auth.Session
	MemberService cl.MemberService
	RoomService   cl.RoomService
}

func (h CompleteAuthHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)

	err := auth.FinishAuthentication(context.Background(), r, session, h.MemberService, h.RoomService)
	if err != nil {
		log.Printf("unable to finish authorization: %s", err)
	}
	err = session.Save(r, w)
	panicIfErr(JSONResponseWriter{w}, err, "server: complete auth: cannot save session")

	http.Redirect(w, r, "/player", http.StatusFound)
}

// SpotifyTokenHandler returns Spotify authentication token from authorized user.
type SpotifyTokenHandler struct {
	auth.Session
}

func (h SpotifyTokenHandler) Get(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)
	jsonWriter := JSONResponseWriter{w}

	token, err := auth.GetTokenFromSession(session)
	if err != nil {
		log.Printf("server: spotifyToken: error retrieving token from session: %s", err)
		jsonWriter.Error(apierror.APIErrorUnauthorized, http.StatusForbidden)
		return
	}

	err = session.Save(r, w)
	panicIfErr(jsonWriter, err, "server: spotifyToken: cannot save session")

	jsonWriter.WriteJSONObject(token)
}
