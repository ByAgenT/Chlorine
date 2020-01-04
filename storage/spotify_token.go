package storage

import (
	"golang.org/x/oauth2"
)

type PGTokenRepository struct {
	Storage *DBStorage
}

func (r PGTokenRepository) GetToken(id ID) (*Token, error) {
	token := new(Token)
	err := r.Storage.QueryRow("SELECT id, access_token, expiry, refresh_token, token_type FROM room WHERE id = $1", id).Scan(
		&token.ID, &token.AccessToken, &token.Expiry, &token.RefreshToken, &token.TokenType)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r PGTokenRepository) SaveToken(token *Token) error {
	if token.ID == nil {
		var id ID
		err := r.Storage.QueryRow(
			"INSERT INTO spotify_token (access_token, expiry, refresh_token, token_type) VALUES ($1, $2, $3, $4) RETURNING id",
			token.AccessToken, token.Expiry, token.RefreshToken, token.TokenType).Scan(&id)
		if err != nil {
			return err
		}
		token.ID = &id
		return nil
	}
	_, err := r.Storage.Exec("UPDATE spotify_token SET access_token=$2, expiry=$3, refresh_token=$4, token_type=$5 WHERE id = $1;",
		token.ID, token.AccessToken, token.Expiry, token.RefreshToken, token.TokenType)
	if err != nil {
		return err
	}
	return nil
}

func (r PGTokenRepository) GetRoomToken(roomID ID) (*Token, error) {
	token := new(Token)
	err := r.Storage.QueryRow("SELECT spotify_token.id, spotify_token.access_token, spotify_token.expiry, spotify_token.refresh_token, spotify_token.token_type FROM room INNER JOIN spotify_token ON room.spotify_token = spotify_token.id WHERE room.id = $1", roomID).Scan(
		&token.ID, &token.AccessToken, &token.Expiry, &token.RefreshToken, &token.TokenType)
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ToOAuthToken converts Spotify token to OAuth Token object.
func (st Token) ToOAuthToken() (*oauth2.Token, error) {
	token := new(oauth2.Token)
	token.AccessToken = st.AccessToken
	token.Expiry = st.Expiry
	token.RefreshToken = st.RefreshToken
	token.TokenType = st.TokenType
	return token, nil
}
