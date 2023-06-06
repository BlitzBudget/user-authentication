package models

type RequestParameter struct {
	Email       *string `validate:"required" json:"email"`
	AccessToken *string `validate:"required" json:"access_token"`
}
