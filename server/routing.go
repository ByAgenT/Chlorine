package server

import (
	"github.com/gorilla/mux"
)

// GetApplicationHandler create dispatcher router with all application routes.
func GetApplicationHandler() *mux.Router {
	router := mux.NewRouter()

	setupMiddlewares(router)

	// Connect route sets to the main router.
	authRouting(router)
	spotifyRouting(router)
	chlorineRouting(router)
	wsRouting(router)

	return router
}

func setupMiddlewares(router *mux.Router) {
	router.Use(LogMiddleware)
}

func authRouting(router *mux.Router) {
	router.Handle("/login", loginHandler).Methods("GET")
	router.Handle("/authcomplete", completeAuthHandler).Methods("GET")
	router.Handle("/token", spotifyTokenHandler).Methods("GET")
}

func spotifyRouting(router *mux.Router) {
	router.Handle("/me/playlists", playlistsHandler).Methods("GET")
	router.Handle("/me/player/devices", availableDevicesHandler).Methods("GET")
	router.Handle("/me/player/", playbackHandler).Methods("GET", "PUT")
	router.Handle("/play", spotifyPlayHandler).Methods("POST")
	router.Handle("/search", searchSongHandler).Methods("GET")
}

func chlorineRouting(router *mux.Router) {
	router.Handle("/room", roomHandler).Methods("GET")
	router.Handle("/room/members", roomMembersHandler).Methods("GET")
	router.Handle("/room/songs", roomSongsHandler).Methods("GET", "POST", "PUT")
	router.Handle("/room/songs/spotify", roomSongsSpotifiedHandler).Methods("GET")
	router.Handle("/member", memberHandler).Methods("GET", "POST")
}

func wsRouting(router *mux.Router) {
	router.HandleFunc("/ws", WebSocketHandler)
}
