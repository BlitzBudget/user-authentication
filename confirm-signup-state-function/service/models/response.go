package models

type HTTPResponse struct {
	Email    *string `validate:"required" json:"email"`
	Password *string `validate:"required" json:"new_password"`
}
