package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader            = websocket.Upgrader{CheckOrigin: checkOriginFunction}
	checkOriginFunction = func(r *http.Request) bool {
		return true
	}
)

// ServeWSConnection initiate websocket handshake and register new connection in websocket hub.
func ServeWSConnection(hub *Hub, w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic("Unhandled upgrade error.")
	}
	client := &Client{
		hub:     hub,
		conn:    connection,
		send:    make(chan []byte, 512),
		receive: make(chan []byte, 512),
	}
	hub.register <- client
	go client.serve()
}
