package main

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type userConfig struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(username string, jwtKey []byte, expirationTime int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &userConfig{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Subject:   username,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "API - Xongl Cloud Pvt. Ltd.",
			ExpiresAt: expirationTime,
		},
	}
	return token.SignedString(jwtKey)
}

func verifyToken(tokenString string, jwtKey []byte) (*userConfig, bool, error) {
	claims := &userConfig{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err == jwt.ErrSignatureInvalid || !token.Valid {
		return claims, false, err
	}
	return claims, true, nil
}
