package cl

import (
	"chlorine/storage"
	"fmt"
)

// CreateRoom creates new Room with provided Spotify token and room config to the storage passed as an argument.
// If token and/or config are not saved to database, they will be automatically saved and assigned
// to the created room. Currently, this is not transactional operation.
// TODO: Make room creation process in one transaction and revert on error.
func CreateRoom(token *storage.SpotifyToken, config *storage.RoomConfig, s *storage.DBStorage) (*storage.Room, error) {
	if token.ID == nil {
		err := s.SaveToken(token)
		if err != nil {
			return nil, fmt.Errorf("chlorine: cannot create room: cannot save token: %s", err)
		}
	}
	if config.ID == nil {
		err := s.SaveRoomConfig(config)
		if err != nil {
			return nil, fmt.Errorf("chlorine: cannot create room: cannot save config: %s", err)
		}
	}
	room := &storage.Room{SpotifyTokenID: storage.Reference(*token.ID), ConfigID: storage.Reference(*config.ID)}
	err := s.SaveRoom(room)
	if err != nil {
		return nil, fmt.Errorf("chlorine: cannot create room: %s", err)
	}
	return room, nil
}
