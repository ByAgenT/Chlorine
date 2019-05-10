package storage

import (
	"time"

	"golang.org/x/oauth2"
)

// SpotifyToken contains Spotify OAuth 2.0 token information.
type SpotifyToken struct {
	ID           *ID       `json:"id,omitempty"`
	AccessToken  string    `json:"access_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type,omitempty"`
}

// SaveToken create a new entry in database and populate ID of a token object
// or update database entry if SpotifyToken object contains ID.
func (s DBStorage) SaveToken(token *SpotifyToken) error {
	if token.ID == nil {
		var id ID
		err := s.QueryRow(
			"INSERT INTO spotify_token (access_token, expiry, refresh_token, token_type) VALUES ($1, $2, $3, $4) RETURNING id",
			token.AccessToken, token.Expiry, token.RefreshToken, token.TokenType).Scan(&id)
		if err != nil {
			return err
		}
		token.ID = &id
		return nil
	}
	_, err := s.Exec("UPDATE spotify_token SET access_token=$2, expiry=$3, refresh_token=$4, token_type=$5 WHERE id = $1;",
		token.ID, token.AccessToken, token.Expiry, token.RefreshToken, token.TokenType)
	if err != nil {
		return err
	}
	return nil
}

// RetrieveToken fetches token entry with the provided ID from the database.
func (s DBStorage) RetrieveToken(id ID) (*SpotifyToken, error) {
	token := new(SpotifyToken)
	err := s.QueryRow("SELECT id, access_token, expiry, refresh_token, token_type FROM room WHERE id = $1", id).Scan(
		&token.ID, &token.AccessToken, &token.Expiry, &token.RefreshToken, &token.TokenType)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// GetRoomToken returns token for the given room.
func (s DBStorage) GetRoomToken(roomID ID) (*SpotifyToken, error) {
	token := new(SpotifyToken)
	err := s.QueryRow("SELECT spotify_token.id, spotify_token.access_token, spotify_token.expiry, spotify_token.refresh_token, spotify_token.token_type FROM room INNER JOIN spotify_token ON room.spotify_token = spotify_token.id WHERE room.id = $1", roomID).Scan(
		&token.ID, &token.AccessToken, &token.Expiry, &token.RefreshToken, &token.TokenType)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ToOAuthToken converts Spotify token to OAuth Token object.
func (st SpotifyToken) ToOAuthToken() (*oauth2.Token, error) {
	token := new(oauth2.Token)
	token.AccessToken = st.AccessToken
	token.Expiry = st.Expiry
	token.RefreshToken = st.RefreshToken
	token.TokenType = st.TokenType
	return token, nil
}
