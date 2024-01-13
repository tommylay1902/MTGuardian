package model

import "github.com/google/uuid"

type Auth struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email        *string   `json:"email" gorm:"unique;"`
	Password     *string   `json:"password"`
	RefreshToken *string   `json:"token"`
}
