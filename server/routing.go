package server

import (
	"chlorine/middleware"
	"chlorine/music"
	"net/http"
)

var (
	musicService           = &music.SpotifyService{}
	authenticationProvider = &music.SpotifySessionAuthentication{}
	externalMusicHandler   = ExternalMusicHandler{MusicService: musicService, AuthenticationProvider: authenticationProvider}
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
	playlistsHandler := MyPlaylistsHandler{ExternalMusicHandler: externalMusicHandler}
	availableDevicesHandler := AvailableDevicesHandler{ExternalMusicHandler: externalMusicHandler}
	playbackHandler := PlaybackHandler{ExternalMusicHandler: externalMusicHandler}

	handler.Handle("/me/playlists", injectMiddlewares(playlistsHandler))
	handler.Handle("/me/player/devices", injectMiddlewares(availableDevicesHandler))
	handler.Handle("/me/player/", injectMiddlewares(playbackHandler))
}

func injectMiddlewares(h http.Handler) http.Handler {
	return middleware.ApplyMiddlewares(h, LogMiddleware)
}
