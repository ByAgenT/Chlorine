package cl

import (
	"chlorine/storage"
	"time"
)

type RawToken struct {
	AccessToken  string
	Expiry       time.Time
	RefreshToken string
	TokenType    string
}

type TokenService interface {
	GetToken(id int) (*storage.Token, error)
	SaveToken(token RawToken) error
	GetRoomToken(roomID int) (*storage.Token, error)
}

type ChlorineTokenService struct {
	Repository storage.TokenRepository
}

func (s ChlorineTokenService) GetToken(id int) (*storage.Token, error) {
	return s.Repository.GetToken(id)
}

func (s ChlorineTokenService) SaveToken(token RawToken) error {
	tokenToWrite := &storage.Token{
		AccessToken:  token.AccessToken,
		Expiry:       token.Expiry,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
	}
	return s.Repository.SaveToken(tokenToWrite)
}

func (s ChlorineTokenService) GetRoomToken(roomID int) (*storage.Token, error) {
	return s.Repository.GetRoomToken(roomID)
}
