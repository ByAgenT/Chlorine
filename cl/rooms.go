package cl

import (
	"chlorine/storage"
	"fmt"
)

type RoomService interface {
	CreateRoom(token *storage.Token, config *storage.RoomConfig) (*storage.Room, error)
	GetRoom(roomID int) (*storage.Room, error)
	GetRoomConfig(roomID int) (*storage.RoomConfig, error)
}

type ChlorineRoomService struct {
	Repository      storage.RoomRepository
	TokenRepository storage.TokenRepository
}

// CreateRoom creates new Room with provided Spotify token and room config to the storage passed as an argument.
// If token and/or config are not saved to database, they will be automatically saved and assigned
// to the created room. Currently, this is not transactional operation.
// TODO: Make room creation process in one transaction and revert on error.
func (s ChlorineRoomService) CreateRoom(token *storage.Token, config *storage.RoomConfig) (*storage.Room, error) {
	if token.ID == nil {
		err := s.TokenRepository.SaveToken(token)
		if err != nil {
			return nil, fmt.Errorf("chlorine: cannot create room: cannot save token: %s", err)
		}
	}
	if config.ID == nil {
		err := s.Repository.SaveRoomConfig(config)
		if err != nil {
			return nil, fmt.Errorf("chlorine: cannot create room: cannot save config: %s", err)
		}
	}
	room := &storage.Room{SpotifyTokenID: *token.ID, ConfigID: *config.ID}
	err := s.Repository.SaveRoom(room)
	if err != nil {
		return nil, fmt.Errorf("chlorine: cannot create room: %s", err)
	}
	return room, nil
}

func (s ChlorineRoomService) GetRoom(roomID int) (*storage.Room, error) {
	return s.Repository.GetRoom(roomID)
}

func (s ChlorineRoomService) GetRoomConfig(roomID int) (*storage.RoomConfig, error) {
	panic("implement me")
}
