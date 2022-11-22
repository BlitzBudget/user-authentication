package models

type RequestParameter struct {
	Email         *string `validate:"required" json:"email"`
	Password      *string `validate:"required" json:"password"`
	Name          *string `validate:"required" json:"name"`
	LastName      *string `validate:"required" json:"last_name"`
	CognitoUserID string  `json:"cognito_user_id"`
}
