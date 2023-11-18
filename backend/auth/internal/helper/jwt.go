package helper

import (
	"github.com/golang-jwt/jwt"
)

var (
	key []byte
	t   *jwt.Token
)

func InitJwtHelper(secret string) {
	key = []byte(secret)
}

func GenerateAccessToken(email *string) (*string, error) {

	t = jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	//time.Now().Add(10 * time.Minute)
	claims["exp"] = 10
	claims["authorized"] = true
	claims["email"] = *email
	jwt, err := t.SignedString(key)
	if err != nil {
		return nil, err
	}

	return &jwt, nil
}
