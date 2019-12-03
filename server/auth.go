package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/cl"
	"context"
	"log"
	"net/http"
)

// LoginHandler initiates Chlorine room and start OAuth2 authentication flow for Spotify.
type LoginHandler struct {
	auth.Session
	StorageHandler
}

func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	StorageHandler
	MemberService cl.MemberService
	RoomService   cl.RoomService
}

func (h CompleteAuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h SpotifyTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
