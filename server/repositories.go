package server

import "chlorine/storage"

var (
	songRepository = storage.PGSongRepository{Storage: dbStorage}
)
