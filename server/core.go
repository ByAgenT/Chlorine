package server

import (
	"chlorine/auth"
	"chlorine/music"
	"chlorine/storage"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/zmb3/spotify"
)

var (
	dbStorage *storage.DBStorage
	dbConfig  = storage.DatabaseConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     os.Getenv("POSTGRES_DATABASE")}
)

// ExternalMusicHandler contains external MusicService and authentication provider for it to retrieve music information.
type ExternalMusicHandler struct {
	auth.Session
	MusicService           music.Service
	AuthenticationProvider auth.SessionAuthenticaton
}

// StorageHandler contains database provider and allow handlers to work with storage.
type StorageHandler struct {
	storage *storage.DBStorage
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
	handler := GetApplicationHandler()
	err := http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err)
	}
}

// InitSpotifyClientFromSession doing client initialization from session storage.
func InitSpotifyClientFromSession(s *sessions.Session) (*spotify.Client, error) {
	authenticator := auth.GetSpotifyAuthenticator()
	token, err := auth.GetTokenFromSession(s)
	if err != nil {
		return nil, err
	}
	client := authenticator.NewClient(token)
	return &client, nil
}

func init() {
	gob.Register(&time.Time{})
	gob.Register(&time.Location{})
	gob.Register(storage.ID(0))
}
