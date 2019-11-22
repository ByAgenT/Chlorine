package server

import (
	"net/http"
)

// GetApplicationHandler create dispatcher handler with all application routes.
func GetApplicationHandler() *http.ServeMux {
	handler := http.NewServeMux()

	// Connect route sets to the main handler.
	authRouting(handler)
	spotifyRouting(handler)
	chlorineRouting(handler)
	wsRouting(handler)

	return handler
}

func authRouting(handler *http.ServeMux) {
	handler.Handle("/login", injectMiddlewares(loginHandler))
	handler.Handle("/authcomplete", injectMiddlewares(completeAuthHandler))
	handler.Handle("/token", injectMiddlewares(spotifyTokenHandler))
}

func spotifyRouting(handler *http.ServeMux) {
	handler.Handle("/me/playlists", injectMiddlewares(playlistsHandler))
	handler.Handle("/me/player/devices", injectMiddlewares(availableDevicesHandler))
	handler.Handle("/me/player/", injectMiddlewares(playbackHandler))
	handler.Handle("/play", injectMiddlewares(spotifyPlayHandler))
	handler.Handle("/search", injectMiddlewares(searchSongHandler))
}

func chlorineRouting(handler *http.ServeMux) {
	handler.Handle("/room", injectMiddlewares(roomHandler))
	handler.Handle("/room/members", injectMiddlewares(roomMembersHandler))
	handler.Handle("/room/songs", injectMiddlewares(roomSongsHandler))
	handler.Handle("/room/songs/spotify", injectMiddlewares(roomSongsSpotifiedHandler))
	handler.Handle("/member", injectMiddlewares(memberHandler))
}

func wsRouting(handler *http.ServeMux) {
	handler.HandleFunc("/ws", WebSocketHandler)
}
