package server

import (
	"chlorine/ws"
	"net/http"
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws.ServeWSConnection(webSocketHub, w, r)
}
