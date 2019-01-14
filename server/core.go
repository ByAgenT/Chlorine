package server

import (
	"encoding/gob"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// StartChlorineServer starts Chlorine to listen to HTTP connections on the given port.
func StartChlorineServer(port string) {
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/authcomplete", completeAuth)
	http.HandleFunc("/me/playlists", getMyPlaylists)
	http.ListenAndServe(port, nil)
}

func init() {
	gob.Register(&oauth2.Token{})
	gob.Register(&time.Time{})
}
