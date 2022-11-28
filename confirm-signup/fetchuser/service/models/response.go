package models

import "confirm-signup/service/models"

type FetchUserResponse struct {
	UserAttributes []*models.UserAttribute `validate:"required" json:"user_attributes"`
	Username       *string                 `validate:"required" json:"username"`
}
