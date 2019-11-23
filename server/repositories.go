package server

import "chlorine/storage"

var (
	songRepository   = storage.PGSongRepository{Storage: dbStorage}
	memberRepository = storage.PGMemberRepository{Storage: dbStorage}
)
