package helper

import "golang.org/x/crypto/bcrypt"

func HashAndSaltPassword(pass string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashPassword := string(hash)

	return &hashPassword, nil
}
