package model

import (
	"math/rand"
	"strconv"
	"time"
)

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

type OAuth struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	AccessToken string    `bson:"accessToken" json:"accessToken"`
	TokenType   string    `bson:"tokenType" json:"tokenType"`
	ExpiresIn   time.Time `bson:"expiresIn" json:"expiresIn"`
}

func NewOAuth(response OAuthResponse) *OAuth{

	oAuth := new(OAuth)

	oAuth.ID = "TOK" + strconv.Itoa(rand.Intn(10000000))
	oAuth.AccessToken = response.AccessToken
	oAuth.TokenType = response.TokenType
	oAuth.ExpiresIn = time.Now().Add(time.Second * time.Duration(response.ExpiresIn))

	return oAuth
}