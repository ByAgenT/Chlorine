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
	storageHandler         = StorageHandler{}
)

// GetApplicationHandler injects application routes and services into ServeMux and returns it.
func GetApplicationHandler() *http.ServeMux {
	handler := http.NewServeMux()
	storageHandler.storage = dbStorage

	authRouting(handler)
	spotifyRouting(handler)
	chlorineRouting(handler)

	return handler
}

func authRouting(handler *http.ServeMux) {
	loginHandler := LoginHandler{StorageHandler: storageHandler}
	completeAuthHandler := CompleteAuthHandler{StorageHandler: storageHandler}
	spotifyTokenHandler := SpotifyTokenHandler{}

	handler.Handle("/login", injectMiddlewares(loginHandler))
	handler.Handle("/authcomplete", injectMiddlewares(completeAuthHandler))
	handler.Handle("/token", injectMiddlewares(spotifyTokenHandler))
}

func spotifyRouting(handler *http.ServeMux) {
	playlistsHandler := MyPlaylistsHandler{ExternalMusicHandler: externalMusicHandler}
	availableDevicesHandler := AvailableDevicesHandler{ExternalMusicHandler: externalMusicHandler}
	playbackHandler := PlaybackHandler{ExternalMusicHandler: externalMusicHandler}

	handler.Handle("/me/playlists", injectMiddlewares(playlistsHandler))
	handler.Handle("/me/player/devices", injectMiddlewares(availableDevicesHandler))
	handler.Handle("/me/player/", injectMiddlewares(playbackHandler))
}

func chlorineRouting(handler *http.ServeMux) {
	roomHandler := RoomHandler{StorageHandler: storageHandler}
	memberHandler := MemberHandler{StorageHandler: storageHandler}
	roomMembersHandler := RoomMembersHandler{StorageHandler: storageHandler}

	handler.Handle("/room", injectMiddlewares(roomHandler))
	handler.Handle("/room/members", injectMiddlewares(roomMembersHandler))
	handler.Handle("/member", injectMiddlewares(memberHandler))
}

func injectMiddlewares(h http.Handler) http.Handler {
	return middleware.ApplyMiddlewares(h, LogMiddleware)
}
