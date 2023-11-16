package models

import "github.com/google/uuid"

type Account struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email        *string   `json:"email"`
	Password     *string   `json:"password"`
	RefreshToken *string   `json:"refresh"`
}
