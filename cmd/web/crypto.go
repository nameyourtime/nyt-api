package main

import (
	"github.com/dchest/uniuri"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var appKey []byte

func initJwt(key string) {
	appKey = []byte(key)
}

type Claims struct {
	UserID string
	jwt.StandardClaims
}

func hashAndSalt(pwd string) (string, error) {
	bytePwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func generateRefreshToken() (string, time.Time) {
	t := uniuri.NewLen(40)
	exp := time.Now().Add(24 * time.Hour * 30)
	return t, exp
}

func generateAccessToken(userID string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString(appKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
