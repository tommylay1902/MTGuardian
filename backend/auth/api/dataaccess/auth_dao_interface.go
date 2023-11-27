package dataaccess

import (
	"github.com/google/uuid"
	"github.com/tommylay1902/authmicro/internal/models"
)

type IAuthDAO interface {
	CreateAuth(auth *models.Auth) (*uuid.UUID, error)
	GetHashFromEmail(email *string) (*string, error)
	GetTokenFromEmail(email *string) (*string, error)
	InsertNewRefreshToken(email *string, token *string) error
}
