package server

import (
	"github.com/gorilla/sessions"
)

// Session is a structure for creating handlers with session.
type Session struct {
	session *sessions.Session
}
