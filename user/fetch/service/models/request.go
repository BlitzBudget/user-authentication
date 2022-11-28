package models

type RequestParameter struct {
	AccessToken    *string        `validate:"required" json:"access_token"`
}