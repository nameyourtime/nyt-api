package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrNonUniqueEmail = errors.New("models: user with this email already registered")
var ErrNonUniqueCode = errors.New("models: verification code for user already exists")

type Healthcheck struct {
	Status     string    `json:"status"`
	ServerTime time.Time `json:"server_time"`
}

type Token struct {
	AccessToken     string    `json:"access_token,omitempty"`
	RefreshToken    string    `json:"refresh_token,omitempty"`
	RefreshTokenExp time.Time `json:"refresh_token_exp"`
}

type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	Status   string    `json:"status"`
	Created  time.Time `json:"created"`
	Token    Token     `json:"token,omitempty"`
}

type VerificationCode struct {
	UserID  string
	Code    string
	CodeExp time.Time
}
