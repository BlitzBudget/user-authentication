package models

type RequestParameter struct {
	Email            *string `validate:"required" json:"email"`
	ConfirmationCode *string `validate:"required" json:"confirmation_code"`
}
