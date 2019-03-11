package storage

import "time"

// Room is a structural representation of a party room object. Party room have
// own configuration and have Spotify token from owner to allow control music.
type Room struct {
	ID             *ID       `json:"id,omitempty"`
	SpotifyTokenID Reference `json:"spotify_token,omitempty"`
	ConfigID       Reference `json:"config_id,omitempty"`
	CreatedDate    time.Time `json:"created_date,omitempty"`
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
		err := rows.Scan(&room.ID, &room.SpotifyTokenID, &room.ConfigID, &room.CreatedDate)
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

// SaveRoom performs inserting of a new entry into database if ID is not present
// or performs update of an entry with the given ID in the Room object.
func (s DBStorage) SaveRoom(room *Room) error {
	if room.ID == nil {
		var id ID
		room.CreatedDate = time.Now().UTC()
		err := s.QueryRow(
			"INSERT INTO room (config_id, spotify_token, created_date) VALUES ($1, $2, $3) RETURNING id",
			room.ConfigID, room.SpotifyTokenID, room.CreatedDate).Scan(&id)
		if err != nil {
			return err
		}
		room.ID = &id
		return nil
	}
	_, err := s.Exec("UPDATE room SET config_id=$1, spotify_token=$2 WHERE id = 1", room.ConfigID, room.SpotifyTokenID)
	if err != nil {
		return err
	}
	return nil
}

// SaveRoomConfig performs inserting of a new entry into database if ID is not present
// or performs update of an entry with the given ID in the RoomConfig object.
func (s DBStorage) SaveRoomConfig(rc *RoomConfig) error {
	if rc.ID == nil {
		var id ID
		rc.CreatedDate = time.Now().UTC()
		err := s.QueryRow(
			"INSERT INTO room_config (songs_per_member, max_members, created_date) VALUES ($1, $2, $3) RETURNING id",
			rc.SongsPerMember, rc.MaxMembers, rc.CreatedDate).Scan(&id)
		if err != nil {
			return err
		}
		rc.ID = &id
		return nil
	}
	_, err := s.Exec("UPDATE room_config SET songs_per_member=$1, max_members=$2 WHERE id = 1", rc.SongsPerMember, rc.MaxMembers)
	if err != nil {
		return err
	}
	return nil
}
