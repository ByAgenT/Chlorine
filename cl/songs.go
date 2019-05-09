package cl

import (
	"chlorine/storage"
	"fmt"
)

// CreateSong creates new song.
func CreateSong(spotifyID string, roomID int, prevSongID int, nextSongID int, memberCreated *storage.Member, s *storage.DBStorage) (*storage.Song, error) {
	var prevSong, nextSong *storage.Reference
	if prevSongID != 0 {
		ref := storage.Reference(prevSongID)
		prevSong = &ref
	}
	if nextSongID != 0 {
		ref := storage.Reference(nextSongID)
		nextSong = &ref
	}
	song := &storage.Song{SpotifyID: spotifyID, RoomID: storage.Reference(roomID),
		PreviousSongID: prevSong, NextSongID: nextSong,
		MemberAddedID: storage.Reference(*memberCreated.ID)}
	err := s.SaveSong(song)
	if err != nil {
		return nil, fmt.Errorf("chlorine: cannot create song: %s", err)
	}
	return song, nil
}

// UpdateSong updates new song.
func UpdateSong(id int, spotifyID string, roomID int, prevSongID int, nextSongID int, memberCreated *storage.Member, s *storage.DBStorage) (*storage.Song, error) {
	var prevSong, nextSong *storage.Reference
	identifier := storage.ID(id)
	if prevSongID != 0 {
		ref := storage.Reference(prevSongID)
		prevSong = &ref
	}
	if nextSongID != 0 {
		ref := storage.Reference(nextSongID)
		nextSong = &ref
	}
	song := &storage.Song{ID: &identifier, SpotifyID: spotifyID, RoomID: storage.Reference(roomID),
		PreviousSongID: prevSong, NextSongID: nextSong,
		MemberAddedID: storage.Reference(*memberCreated.ID)}
	err := s.SaveSong(song)
	if err != nil {
		return nil, fmt.Errorf("chlorine: cannot create song: %s", err)
	}
	return song, nil
}
