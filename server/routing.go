package server

import (
	"github.com/gorilla/mux"
	"net/http"
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
	router.Handle("/login", createDispatchableHandler(loginHandler)).Methods(http.MethodGet)
	router.Handle("/authcomplete", createDispatchableHandler(completeAuthHandler)).Methods(http.MethodGet)
	router.Handle("/token", createDispatchableHandler(spotifyTokenHandler)).Methods(http.MethodGet)
}

func spotifyRouting(router *mux.Router) {
	router.Handle("/me/playlists", createDispatchableHandler(playlistsHandler)).Methods(http.MethodGet)
	router.Handle("/me/player/devices", createDispatchableHandler(availableDevicesHandler)).Methods(http.MethodGet)
	router.Handle("/me/player/", createDispatchableHandler(playbackHandler)).Methods(http.MethodGet, http.MethodPut)
	router.Handle("/play", createDispatchableHandler(spotifyPlayHandler)).Methods(http.MethodPost)
	router.Handle("/search", createDispatchableHandler(searchSongHandler)).Methods(http.MethodGet)
}

func chlorineRouting(router *mux.Router) {
	router.Handle("/room", createDispatchableHandler(roomHandler)).Methods(http.MethodGet)
	router.Handle("/room/members", createDispatchableHandler(roomMembersHandler)).Methods(http.MethodGet)
	router.Handle("/room/songs", createDispatchableHandler(roomSongsHandler)).Methods(http.MethodGet, http.MethodPost, http.MethodPut)
	router.Handle("/room/songs/{id:[0-9]+}", createDispatchableHandler(roomSongsDetailHandler)).Methods(http.MethodDelete)
	router.Handle("/room/songs/spotify", createDispatchableHandler(roomSongsSpotifiedHandler)).Methods(http.MethodGet)
	router.Handle("/member", createDispatchableHandler(memberHandler)).Methods(http.MethodGet, http.MethodPost)
}

func wsRouting(router *mux.Router) {
	router.Handle("/ws", wsHandler)
}
