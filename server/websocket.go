package server

import (
	"chlorine/auth"
	"chlorine/cl"
	"chlorine/ws"
	"net/http"
)

var (
	roomWSConnections = make(map[int][]*ws.Client)
)

func initWebSocketActions(hub *ws.Hub) {
	dispatcher := ws.NewDispatcher()
	hub.AttachDispatcher(dispatcher)
}

type WebSocketInitHandler struct {
	auth.Session
	MemberService cl.MemberService
}

func (h WebSocketInitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session := h.InitSession(r)

	memberID, ok := session.Values["MemberID"].(int)
	member, err := h.MemberService.GetMember(memberID)
	if !ok || err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	client := ws.ServeWSConnection(webSocketHub, w, r)
	roomWSConnections[int(member.RoomID)] = append(roomWSConnections[int(member.RoomID)], client)
}
