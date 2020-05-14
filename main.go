package main

import (
	"chlorine/server"
	"fmt"
	"log"
	"os"
)

const (
	// ServerHost is a default value for host.
	ServerHost = "localhost"

	// ServerPort is a default value for port.
	ServerPort = "8080"
)

func main() {
	host, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		log.Printf("server: hostname not defined, using default (%s).", ServerHost)
		host = ServerHost
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Printf("server: port not defined, using default (%s).", ServerPort)
		port = ServerPort
	}
	serveString := fmt.Sprintf("%s:%s", host, port)
	log.Printf("server: start listening on %s.", serveString)
	server.StartChlorineServer(serveString)
	fmt.Println("Server shut down.")
}
