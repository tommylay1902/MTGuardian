package dao

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/authmicro/internal/model"
)

type IAuthDAO interface {
	CreateAuth(auth *model.Auth) (*uuid.UUID, error)
	GetHashFromEmail(email *string) (*string, error)
	GetTokenFromEmail(email *string) (*string, error)
	InsertNewRefreshToken(email *string, token *string) error
}
