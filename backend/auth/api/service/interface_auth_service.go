package service

import (
	dto "github.com/tommylay1902/authmicro/internal/dto"
	"github.com/tommylay1902/authmicro/internal/model"
)

type IAuthService interface {
	CreateAuth(authDTO *dto.AuthDTO) (*string, error)
	Login(authDTO *dto.AuthDTO) (*string, error)
	Refresh(accessToken *model.AccessToken) (*string, error)
}
