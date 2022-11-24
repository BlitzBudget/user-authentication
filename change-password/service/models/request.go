package models

type RequestParameter struct {
	AccessToken      *string `validate:"required" json:"access_token"`
	PreviousPassword *string `validate:"required" json:"previous_password"`
	ProposedPassword *string `validate:"required" json:"proposed_password"`
}
