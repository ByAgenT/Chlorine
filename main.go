package main

import (
	"chlorine/server"
	"fmt"
	"log"
)

const (
	// ServerPort when Chlorine server will serve HTTP connections.
	ServerPort = ":8080"
)

func main() {
	log.Println("Server starting...")
	server.StartChlorineServer(ServerPort)
	fmt.Println("Server shut down.")
}
