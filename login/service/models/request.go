package models

type RequestParameter struct {
	Email    *string `validate:"required" json:"email"`
	Password *string `validate:"required" json:"password"`
}
