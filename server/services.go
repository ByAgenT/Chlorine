package server

import (
	"chlorine/auth"
	"chlorine/cl"
	"chlorine/music"
	"chlorine/music/spotify"
)

var (
	// Music service binding
	musicService music.Service

	// Authentication provider binding
	authenticationProvider auth.SessionAuthentication

	songService   cl.SongService
	memberService cl.MemberService
	roomService   cl.RoomService
	tokenService  cl.TokenService
)

func initServices() {
	// Music service binding
	musicService = &spotify.Service{}

	// Authentication provider binding
	authenticationProvider = &spotify.SessionAuthentication{}

	songService = cl.ChlorineSongService{Repository: songRepository}
	memberService = cl.ChlorineMemberService{Repository: memberRepository}
	roomService = cl.ChlorineRoomService{Repository: roomRepository, TokenRepository: tokenRepository}
	tokenService = cl.ChlorineTokenService{Repository: tokenRepository}
}
