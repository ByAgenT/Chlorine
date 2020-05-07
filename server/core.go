package server

import (
	"chlorine/apierror"
	"chlorine/auth"
	"chlorine/music"
	"chlorine/storage"
	"chlorine/ws"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
)

var (
	dbStorage *storage.DBStorage
	dbConfig  = storage.DatabaseConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     os.Getenv("POSTGRES_DATABASE")}
	webSocketHub = ws.CreateHub()
)

// ExternalMusicHandler contains external MusicService and authentication provider for it to retrieve music information.
type ExternalMusicHandler struct {
	auth.Session
	MusicService           music.Service
	AuthenticationProvider auth.SessionAuthentication
}

// GetClient return authenticate music service and return client instance.
func (h ExternalMusicHandler) GetClient(session *sessions.Session) (music.Client, error) {
	authenticator, err := h.AuthenticationProvider.GetAuth(session)
	if err != nil {
		return nil, err
	}
	client, err := h.MusicService.Authenticate(authenticator)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// StartChlorineServer starts Chlorine to listen to HTTP connections on the given port.
func StartChlorineServer(port string) {
	dbStorage = storage.ConnectDatabase(dbConfig)
	initWebSocketActions(webSocketHub)
	go webSocketHub.Run()
	initRepositories()
	initServices()
	initHandlers()
	handler := GetApplicationHandler()
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}

func panicIfErr(jsonWriter JSONResponseWriter, err error, pretext string) {
	if err != nil {
		jsonWriter.Error(apierror.APIServerError, http.StatusInternalServerError)
		panic(fmt.Sprintf("%s: %s", pretext, err.Error()))
	}
}

func init() {
	gob.Register(&time.Time{})
	gob.Register(&time.Location{})
	gob.Register(storage.ID(0))
}
