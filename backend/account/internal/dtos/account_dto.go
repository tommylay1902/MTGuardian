package dto

import "github.com/google/uuid"

type AccountDTO struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email    *string   `json:"email"`
	Password *string   `json:"password"`
}
