package helper

import (
	"fmt"
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

func GenerateToken() (*string, error) {

	t = jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["user"] = "username"
	jwt, err := t.SignedString(key)
	if err != nil {
		fmt.Println(err, "PRINTING ERROR")
		return nil, err
	}

	return &jwt, nil
}
