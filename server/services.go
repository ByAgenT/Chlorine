package server

import (
	"chlorine/cl"
	"chlorine/music/spotify"
)

var (
	// Music service binding
	musicService = &spotify.Service{}

	// Authentication provider binding
	authenticationProvider = &spotify.SessionAuthentication{}

	songService = cl.ChlorineSongService{Repository: songRepository}
)
