package server

import (
	"akovalyov/chlorine/middleware"
	"net/http"
)

// GetApplicationHandler create ServeMux instance with all applicaion routes.
func GetApplicationHandler() *http.ServeMux {
	handler := http.NewServeMux()

	handler.HandleFunc("/login", middleware.ApplyMiddlewares(HandleLogin, middleware.LogMiddleware))
	handler.HandleFunc("/authcomplete", CompleteAuth)
	handler.HandleFunc("/token", middleware.ApplyMiddlewares(SpotifyToken, middleware.LogMiddleware))

	handler.HandleFunc("/me/playlists", middleware.ApplyMiddlewares(MyPlaylists, middleware.LogMiddleware))
	handler.HandleFunc("/me/player/devices", middleware.ApplyMiddlewares(AvailableDevices, middleware.LogMiddleware))

	return handler
}
