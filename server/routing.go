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
	router.Handle("/login", createDispatchableHandler(loginHandler)).Methods("GET")
	router.Handle("/authcomplete", createDispatchableHandler(completeAuthHandler)).Methods("GET")
	router.Handle("/token", createDispatchableHandler(spotifyTokenHandler)).Methods("GET")
}

func spotifyRouting(router *mux.Router) {
	router.Handle("/me/playlists", createDispatchableHandler(playlistsHandler)).Methods("GET")
	router.Handle("/me/player/devices", createDispatchableHandler(availableDevicesHandler)).Methods("GET")
	router.Handle("/me/player/", createDispatchableHandler(playbackHandler)).Methods("GET", "PUT")
	router.Handle("/play", createDispatchableHandler(spotifyPlayHandler)).Methods("POST")
	router.Handle("/search", createDispatchableHandler(searchSongHandler)).Methods("GET")
}

func chlorineRouting(router *mux.Router) {
	router.Handle("/room", createDispatchableHandler(roomHandler)).Methods("GET")
	router.Handle("/room/members", createDispatchableHandler(roomMembersHandler)).Methods("GET")
	router.Handle("/room/songs", createDispatchableHandler(roomSongsHandler)).Methods("GET", "POST", "PUT")
	router.Handle("/room/songs/{id:[0-9]+}", createDispatchableHandler(roomSongsDetailHandler)).Methods("DELETE")
	router.Handle("/room/songs/spotify", createDispatchableHandler(roomSongsSpotifiedHandler)).Methods("GET")
	router.Handle("/member", createDispatchableHandler(memberHandler)).Methods("GET", "POST")
}

func wsRouting(router *mux.Router) {
	router.Handle("/ws", wsHandler)
}
