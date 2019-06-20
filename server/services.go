package server

import "chlorine/music"

var (
	// Music service binding
	musicService = &music.SpotifyService{}

	// Authentication provider binding
	authenticationProvider = &music.SpotifySessionAuthentication{}
)
