package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	key []byte
	t   *jwt.Token
)

func InitJwtHelper(secret string) {
	key = []byte(secret)
}

func GetKey() []byte {
	return key
}

func GenerateRefreshToken(email *string) (*string, error) {

	claims := jwt.RegisteredClaims{
		Subject:   *email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := t.SignedString(key)
	if err != nil {
		return nil, err
	}
	return &jwt, nil
}

func GenerateAccessToken(email *string) (*string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   *email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(720 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := t.SignedString(key)
	if err != nil {

		return nil, err
	}
	return &jwt, nil
}

func IsValidToken(tokenString string) bool {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Check for parsing errors
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return false
	}

	// Check if the token is valid
	if token.Valid {
		return true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		// Handle validation errors
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("Token is malformed")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			fmt.Println("Token has expired or not valid yet")
		} else {
			fmt.Println("Token validation error:", err)
		}
		return false
	} else {
		fmt.Println("Token validation error:", err)
		return false
	}
}
