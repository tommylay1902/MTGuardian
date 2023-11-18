package models

import (
	"fmt"
	"strings"
	"time"
)

type Claims struct {
	Email      string `json:"email"`
	Subject    string `json:"sub"`
	Expiration string `json:"exp"`
}

func (c *Claims) Valid() error {
	// Adjusted layout to handle the date format without time-related placeholders and monotonic clock offset
	layout := "2006-01-02"

	fmt.Println("before comparing", c.Expiration, time.Now().Local())

	// Extract the expiration date without monotonic clock offset
	expirationString := strings.Fields(c.Expiration)[0]

	expirationTime, err := time.Parse(layout, expirationString)
	if err != nil {
		fmt.Println("Error parsing from model:", err)
		fmt.Println(*c)
		return err
	}

	present := time.Now().Local()

	fmt.Println(expirationTime, present)
	// Compare dates without considering time
	if expirationTime.Before(present) {
		return fmt.Errorf("token has expired")
	}

	// You can add more validation logic here based on your needs.

	// If everything is valid, return nil.
	return nil
}
