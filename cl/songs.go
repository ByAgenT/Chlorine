package cl

import (
	"chlorine/storage"
	"fmt"
)

// RawSong is used to create or update song via raw parameters.
type RawSong struct {
	SpotifyID      string
	RoomID         int
	PreviousSongID *int
	NextSongID     *int
	MemberCreated  int
}

type SongService interface {
	CreateSong(songData RawSong) (*storage.Song, error)
	UpdateSong(id int, songData RawSong) (*storage.Song, error)
	GetSong(id int) (*storage.Song, error)
	GetRoomSongs(roomID int) ([]storage.Song, error)
	DeleteSong(id int) error
}

type ChlorineSongService struct {
	Repository storage.SongRepository
}

func (s ChlorineSongService) DeleteSong(id int) error {
	song, err := s.GetSong(id)
	if err != nil {
		return err
	}
	if song.NextSongID != nil {
		nextSong, err := s.GetSong(*song.NextSongID)
		if err != nil {
			return err
		}
		_, err = s.UpdateSong(*nextSong.ID, RawSong{
			SpotifyID:      nextSong.SpotifyID,
			RoomID:         nextSong.RoomID,
			PreviousSongID: song.PreviousSongID,
			NextSongID:     nextSong.NextSongID,
			MemberCreated:  nextSong.MemberAddedID,
		})
		if err != nil {
			return err
		}
	}
	if song.PreviousSongID != nil {
		prevSong, err := s.GetSong(*song.PreviousSongID)
		if err != nil {
			return err
		}
		_, err = s.UpdateSong(*prevSong.ID, RawSong{
			SpotifyID:      prevSong.SpotifyID,
			RoomID:         prevSong.RoomID,
			PreviousSongID: prevSong.PreviousSongID,
			NextSongID:     song.NextSongID,
			MemberCreated:  prevSong.MemberAddedID,
		})
		if err != nil {
			return err
		}
	}
	return s.Repository.DeleteSong(id)
}

func (s ChlorineSongService) GetSong(id int) (*storage.Song, error) {
	return s.Repository.GetSong(id)
}

// GetRoomSongs return all songs that belongs to the provided room ID.
func (s ChlorineSongService) GetRoomSongs(roomID int) ([]storage.Song, error) {
	return s.Repository.GetRoomSongs(roomID)
}

// CreateSong creates a new song from the RawSong data and return Song instance written to the database.
func (s ChlorineSongService) CreateSong(songData RawSong) (*storage.Song, error) {
	song := &storage.Song{SpotifyID: songData.SpotifyID, RoomID: songData.RoomID,
		PreviousSongID: songData.PreviousSongID, NextSongID: songData.NextSongID,
		MemberAddedID: songData.MemberCreated}
	err := s.Repository.SaveSong(song)
	if err != nil {
		return nil, fmt.Errorf("chlorine: cannot create song: %s", err)
	}
	return song, nil
}

// UpdateSong updates song by it's ID with the values provided in RawSong.
func (s ChlorineSongService) UpdateSong(id int, songData RawSong) (*storage.Song, error) {
	identifier := id
	song := &storage.Song{ID: &identifier, SpotifyID: songData.SpotifyID, RoomID: songData.RoomID,
		PreviousSongID: songData.PreviousSongID, NextSongID: songData.NextSongID,
		MemberAddedID: songData.MemberCreated}
	err := s.Repository.SaveSong(song)
	if err != nil {
		return nil, fmt.Errorf("chlorine: cannot create song: %s", err)
	}
	return song, nil
}
