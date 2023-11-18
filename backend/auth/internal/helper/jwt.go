package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	key []byte
	t   *jwt.Token
)

func InitJwtHelper(secret string) {
	key = []byte(secret)
}

func GenerateRefreshToken(email *string) (*string, error) {

	t = jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(720 * time.Hour)
	claims["authorized"] = true
	claims["email"] = *email
	jwt, err := t.SignedString(key)
	if err != nil {
		return nil, err
	}
	return &jwt, nil
}

func GenerateAccessToken(email *string) (*string, error) {

	t = jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(15 * time.Minute)
	claims["authorized"] = true
	claims["email"] = *email
	jwt, err := t.SignedString(key)
	if err != nil {
		return nil, err
	}
	return &jwt, nil
}
