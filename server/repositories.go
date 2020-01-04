package server

import "chlorine/storage"

var (
	songRepository   storage.SongRepository
	memberRepository storage.MemberRepository
	roomRepository   storage.RoomRepository
	tokenRepository  storage.TokenRepository
)

func initRepositories() {
	songRepository = storage.PGSongRepository{Storage: dbStorage}
	memberRepository = storage.PGMemberRepository{Storage: dbStorage}
	roomRepository = storage.PGRoomRepository{Storage: dbStorage}
	tokenRepository = storage.PGTokenRepository{Storage: dbStorage}
}
