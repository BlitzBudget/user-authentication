package models

type RequestParameter struct {
	Email            *string `validate:"required" json:"email"`
	ConfirmationCode *string `validate:"required" json:"confirmation_code"`
	Password         *string `validate:"required" json:"new_password"`
}
