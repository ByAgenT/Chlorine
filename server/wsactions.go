package server

import "chlorine/ws"

func healthCheckAction(message *ws.ClientMessage) *ws.Response {
	return &ws.Response{
		Type:   ws.TypeResponse,
		Status: ws.StatusOK,
		Body:   map[string]interface{}{},
	}
}

func registerAction(message *ws.ClientMessage) *ws.Response {
	return &ws.Response{
		Type:   ws.TypeResponse,
		Status: ws.StatusOK,
		Body:   map[string]interface{}{},
	}
}
