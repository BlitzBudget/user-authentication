package models

type RequestParameter struct {
	RefreshToken *string `validate:"required" json:"refresh_token"`
	UserId       *string `validate:"required" json:"user_id"`
}
