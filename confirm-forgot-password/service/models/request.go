package models

type RequestParameter struct {
	Email            *string `validate:"required" json:"email"`
	Password         *string `validate:"required" json:"new_password"`
	ConfirmationCode *string `validate:"required" json:"confirmation_code"`
}
