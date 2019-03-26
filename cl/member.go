package cl

import (
	"chlorine/storage"
)

// CreateMember create member object and write it to the provided DBStorage.
// TODO: handle the case of invalid room ID.
func CreateMember(name string, roomID int, isAdmin bool, s *storage.DBStorage) (*storage.Member, error) {
	member := &storage.Member{
		Name:    name,
		RoomID:  storage.Reference(roomID),
		IsAdmin: isAdmin}
	err := s.SaveMember(member)
	if err != nil {
		return nil, err
	}
	return member, nil
}
