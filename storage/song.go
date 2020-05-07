package storage

import (
	"log"
	"time"
)

// Song is a struct representation of a Spotify song within Chlorine.
type Song struct {
	Model
	ID             *ID        `json:"id"`
	SpotifyID      string     `json:"spotify_id"`
	RoomID         Reference  `json:"room_id"`
	PreviousSongID *Reference `json:"previous_song_id"`
	NextSongID     *Reference `json:"next_song_id"`
	MemberAddedID  Reference  `json:"member_added_id"`
	CreatedDate    time.Time  `json:"created_date"`
}

type SongRepository interface {
	GetRoomSongs(roomID int) ([]Song, error)
	SaveSong(song *Song) error
}

type PGSongRepository struct {
	Storage *DBStorage
}

func (s PGSongRepository) GetRoomSongs(roomID int) ([]Song, error) {
	query := "SELECT id, spotify_id, room_id, prev_song_id, next_song_id, member_added_id FROM song WHERE room_id = $1 ORDER BY next_song_id"
	rows, err := s.Storage.Query(query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	songs := make([]Song, 0)
	for rows.Next() {
		song := Song{}
		err := rows.Scan(&song.ID, &song.SpotifyID, &song.RoomID, &song.PreviousSongID, &song.NextSongID, &song.MemberAddedID)
		if err != nil {
			log.Printf("error fetching lines: %s", err)
			return nil, err
		}
		songs = append(songs, song)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return songs, nil
}

func (s PGSongRepository) SaveSong(song *Song) error {
	if song.ID == nil {
		var id ID

		song.CreatedDate = time.Now().UTC()
		var err error
		err = s.Storage.QueryRow("INSERT INTO song (spotify_id, room_id, prev_song_id, next_song_id, member_added_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			song.SpotifyID, song.RoomID, song.PreviousSongID, song.NextSongID, song.MemberAddedID).Scan(&id)
		if err != nil {
			return err
		}
		song.ID = &id
		return nil
	}
	_, err := s.Storage.Exec("UPDATE song SET spotify_id=$2, room_id=$3, prev_song_id=$4, next_song_id=$5, member_added_id=$6 WHERE id = $1",
		song.ID, song.SpotifyID, song.RoomID, song.PreviousSongID, song.NextSongID, song.MemberAddedID)
	if err != nil {
		return err
	}
	return nil
}
