package storage

import "time"

// Room is a structural representation of a party room object. Party room have
// own configuration and have Spotify token from owner to allow control music.
type Room struct {
	ID           *ID       `json:"id,omitempty"`
	SpotifyToken string    `json:"spotify_token,omitempty"`
	ConfigID     int       `json:"config_id,omitempty"`
	CreatedDate  time.Time `json:"created_date,omitempty"`
}

// RoomConfig is a config structure for rooms that contains meta information about
// party room.
type RoomConfig struct {
	ID             *ID       `json:"id,omitempty"`
	SongsPerMember int       `json:"songs_per_member,omitempty"`
	MaxMembers     int       `json:"max_members,omitempty"`
	CreatedDate    time.Time `json:"created_date,omitempty"`
}

// GetRooms fetches all room entries from a database and return slice of Room objects.
func (s DBStorage) GetRooms() ([]Room, error) {
	rows, err := s.Query("SELECT id, spotify_token, config_id, created_date FROM room")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rooms := make([]Room, 0)
	for rows.Next() {
		room := Room{}
		err := rows.Scan(&room.ID, &room.SpotifyToken, &room.ConfigID, &room.CreatedDate)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}
