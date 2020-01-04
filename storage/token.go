package storage

import (
	"golang.org/x/oauth2"
	"time"
)

// Token structure contains OAuth 2.0 token information.
type Token struct {
	ID           *ID       `json:"id,omitempty"`
	AccessToken  string    `json:"access_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type,omitempty"`
}

// TokenRepository represents interface for data access to Token object.
type TokenRepository interface {
	GetToken(id ID) (*Token, error)
	SaveToken(token *Token) error
	GetRoomToken(roomID ID) (*Token, error)
}

// OAuthConvertible represents objects that could be converted to the OAuth2 token
// type from "golang.org/x/oauth2" package
type OAuthConvertible interface {
	ToOAuthToken() (*oauth2.Token, error)
}
