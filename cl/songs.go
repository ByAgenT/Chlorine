package cl

import (
	"chlorine/storage"
	"fmt"
)

// RawSong is used to create or update song via raw parameters.
type RawSong struct {
	SpotifyID      string
	RoomID         int
	PreviousSongID int
	NextSongID     int
	MemberCreated  *storage.Member
}

type SongService interface {
	CreateSong(songData RawSong) (*storage.Song, error)
	UpdateSong(id int, songData RawSong) (*storage.Song, error)
	GetRoomSongs(roomID int) ([]storage.Song, error)
	DeleteSong(id int) error
}

type ChlorineSongService struct {
	Repository storage.SongRepository
}

func (s ChlorineSongService) DeleteSong(id int) error {
	return s.Repository.DeleteSong(id)
}

// GetRoomSongs return all songs that belongs to the provided room ID.
func (s ChlorineSongService) GetRoomSongs(roomID int) ([]storage.Song, error) {
	return s.Repository.GetRoomSongs(roomID)
}

// CreateSong creates a new song from the RawSong data and return Song instance written to the database.
func (s ChlorineSongService) CreateSong(songData RawSong) (*storage.Song, error) {
	var prevSong, nextSong *storage.Reference
	if songData.PreviousSongID != 0 {
		ref := storage.Reference(songData.PreviousSongID)
		prevSong = &ref
	}
	if songData.NextSongID != 0 {
		ref := storage.Reference(songData.NextSongID)
		nextSong = &ref
	}
	song := &storage.Song{SpotifyID: songData.SpotifyID, RoomID: storage.Reference(songData.RoomID),
		PreviousSongID: prevSong, NextSongID: nextSong,
		MemberAddedID: storage.Reference(*songData.MemberCreated.ID)}
	err := s.Repository.SaveSong(song)
	if err != nil {
		return nil, fmt.Errorf("chlorine: cannot create song: %s", err)
	}
	return song, nil
}

// UpdateSong updates song by it's ID with the values provided in RawSong.
func (s ChlorineSongService) UpdateSong(id int, songData RawSong) (*storage.Song, error) {
	var prevSong, nextSong *storage.Reference
	identifier := storage.ID(id)
	if songData.PreviousSongID != 0 {
		ref := storage.Reference(songData.PreviousSongID)
		prevSong = &ref
	}
	if songData.NextSongID != 0 {
		ref := storage.Reference(songData.NextSongID)
		nextSong = &ref
	}
	song := &storage.Song{ID: &identifier, SpotifyID: songData.SpotifyID, RoomID: storage.Reference(songData.RoomID),
		PreviousSongID: prevSong, NextSongID: nextSong,
		MemberAddedID: storage.Reference(*songData.MemberCreated.ID)}
	err := s.Repository.SaveSong(song)
	if err != nil {
		return nil, fmt.Errorf("chlorine: cannot create song: %s", err)
	}
	return song, nil
}
