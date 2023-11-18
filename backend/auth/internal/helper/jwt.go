package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/tommylay1902/authmicro/internal/models"
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

	t = jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)
	claims["sub"] = *email
	claims["exp"] = time.Now().Add(720 * time.Hour).Local().String()
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
	claims["sub"] = *email
	claims["exp"] = time.Now().Add(15 * time.Minute).Local().String()
	claims["email"] = *email
	jwt, err := t.SignedString(key)
	if err != nil {
		return nil, err
	}
	return &jwt, nil
}

func IsValidToken(tokenString string) bool {
	claims := &models.Claims{}
	fmt.Println("printing from isvalidtoken", tokenString)
	//parse the expired token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	fmt.Println(token.Valid, "and the claims: ", claims)

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
