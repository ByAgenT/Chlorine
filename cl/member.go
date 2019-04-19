package cl

import (
	"chlorine/storage"
)

// CreateMember create member object and write it to the provided DBStorage.
// TODO: handle the case of invalid room ID.
func CreateMember(name string, roomID int, role int, s *storage.DBStorage) (*storage.Member, error) {
	member := &storage.Member{
		Name:   name,
		RoomID: storage.Reference(roomID),
		Role:   storage.Reference(role)}
	err := s.SaveMember(member)
	if err != nil {
		return nil, err
	}
	return member, nil
}
