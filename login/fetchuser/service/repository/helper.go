package repository

import (
	loginModel "login/service/models"
	"login/fetchuser/service/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ParseResponse(getUserOutput *cognitoidentityprovider.GetUserOutput) *models.FetchUserResponse {

	userAttributes := []*loginModel.UserAttribute{}
	for _, u := range getUserOutput.UserAttributes {
		userAttribute := loginModel.UserAttribute{
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
