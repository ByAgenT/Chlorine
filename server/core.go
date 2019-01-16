package server

import (
	"akovalyov/chlorine/handlers"
	"encoding/gob"
	"log"
	"net/http"
	"time"
)

// StartChlorineServer starts Chlorine to listen to HTTP connections on the given port.
func StartChlorineServer(port string) {
	handler := handlers.GetApplicationHandler()
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gob.Register(&time.Time{})
	gob.Register(&time.Location{})
}
