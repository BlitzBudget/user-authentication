package repository

import (
	"confirm-signup/fetchuser/service/models"
	confirmsignupModels "confirm-signup/service/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ParseResponse(getUserOutput *cognitoidentityprovider.GetUserOutput) *models.FetchUserResponse {

	userAttributes := []*confirmsignupModels.UserAttribute{}
	for _, u := range getUserOutput.UserAttributes {
		userAttribute := confirmsignupModels.UserAttribute{
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
