package cl

import (
	"chlorine/auth"
	"context"

	"github.com/gorilla/sessions"
)

// InitializeLogin setup initial login operations and return authentication URL.
func InitializeLogin(ctx context.Context, session *sessions.Session) string {
	authenticator := auth.GetSpotifyAuthenticator()
	state := auth.CreateRandomState(session)
	session.Values["CSRFState"] = state
	return authenticator.AuthURL(state)
}
