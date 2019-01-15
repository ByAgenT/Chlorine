package server

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"
)

// StartChlorineServer starts Chlorine to listen to HTTP connections on the given port.
func StartChlorineServer(port string) {
	http.HandleFunc("/login", logHandler(handleLogin))
	http.HandleFunc("/authcomplete", completeAuth)
	http.HandleFunc("/me/playlists", logHandler(myPlaylists))
	http.HandleFunc("/token", logHandler(spotifyToken))
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gob.Register(&time.Time{})
	gob.Register(&time.Location{})
}
