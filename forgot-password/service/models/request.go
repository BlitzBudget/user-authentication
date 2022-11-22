package models

type RequestParameter struct {
	Email *string `validate:"required" json:"email"`
}
