package server

import (
	"chlorine/ws"
	"net/http"
)

func initWebSocketActions(hub *ws.Hub) {
	dispatcher := ws.NewDispatcher()
	dispatcher.AttachAction("helloworld", func(message *ws.ClientMessage) *ws.Response {
		return &ws.Response{
			Type:   ws.TypeResponse,
			Status: ws.StatusOK,
			Body:   map[string]interface{}{"hello": "world"},
		}
	})

	hub.AttachDispatcher(dispatcher)
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws.ServeWSConnection(webSocketHub, w, r)
}
