package server

import (
	"akovalyov/chlorine/middleware"
	"net/http"
)

// GetApplicationHandler create ServeMux instance with all applicaion routes.
func GetApplicationHandler() *http.ServeMux {
	handler := http.NewServeMux()
	authRouting(handler)
	spotifyRouting(handler)
	return handler
}

func authRouting(handler *http.ServeMux) {
	handler.Handle("/login", middleware.ApplyMiddlewares(LoginHandler{}, LogMiddleware))
	handler.Handle("/authcomplete", CompleteAuthHandler{})
	handler.Handle("/token", middleware.ApplyMiddlewares(SpotifyTokenHandler{}, LogMiddleware))
}

func spotifyRouting(handler *http.ServeMux) {
	handler.Handle("/me/playlists", middleware.ApplyMiddlewares(MyPlaylistsHandler{}, LogMiddleware))
	handler.Handle("/me/player/devices", middleware.ApplyMiddlewares(AvailableDevicesHandler{}, LogMiddleware))
	handler.Handle("/me/player/", middleware.ApplyMiddlewares(PlaybackHandler{}, LogMiddleware))
}
