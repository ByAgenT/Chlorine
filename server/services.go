package server

import (
	"chlorine/music/spotify"
)

var (
	// Music service binding
	musicService = &spotify.SpotifyService{}

	// Authentication provider binding
	authenticationProvider = &spotify.SpotifySessionAuthentication{}
)
