package dto

type AuthDTO struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
