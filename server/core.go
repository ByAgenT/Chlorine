package server

import (
	"chlorine/apierror"
	"chlorine/storage"
	"chlorine/ws"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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
}
