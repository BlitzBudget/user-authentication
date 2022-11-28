package models

type RequestParameter struct {
	AccessToken    *string        `validate:"required" json:"access_token"`
	UserAttributes *UserAttribute `validate:"required" json:"user_attributes"`
}

type UserAttribute struct {
	FirstName     *string `validate:"required" json:"first_name"`
	LastName *string `validate:"required" json:"last_name"`
}
