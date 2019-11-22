package server

import (
	"chlorine/ws"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: checkOriginFunction}

	checkOriginFunction = func(r *http.Request) bool {
		return true
	}
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws.ServeWSConnection(webSocketHub, w, r)
}
