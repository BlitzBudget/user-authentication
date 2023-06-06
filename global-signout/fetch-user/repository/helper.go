package repository

import (
	"global-signout/fetch-user/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ParseResponse(getUserOutput *cognitoidentityprovider.GetUserOutput) *models.FetchUserResponse {

	userAttributes := []*models.UserAttribute{}
	for _, u := range getUserOutput.UserAttributes {
		userAttribute := models.UserAttribute{
			Name:  u.Name,
			Value: u.Value,
		}

		userAttributes = append(userAttributes, &userAttribute)
	}

	response := models.FetchUserResponse{
		Username:       getUserOutput.Username,
		UserAttributes: userAttributes,
	}

	return &response
}
