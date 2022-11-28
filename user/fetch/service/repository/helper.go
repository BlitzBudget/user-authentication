package repository

import (
	"encoding/json"
	"fetch-user/service/errors"
	"fetch-user/service/models"
	"fmt"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ParseRequest(body *string) (*models.RequestParameter, error) {
	requestParameter := models.RequestParameter{}
	err := json.Unmarshal([]byte(*body), &requestParameter)

	if err != nil {
		fmt.Printf("Repository: There was an error marshalling the bytes to struct: %v", err.Error())
		return nil, err
	}

	requestAsBytes, _ := json.Marshal(requestParameter)
	fmt.Printf("Repository: marshalled bytes to struct: %+v. \n", string(requestAsBytes))

	// Mandatory Fiels Check for the Request
	err = mandatoryFieldsCheck(requestParameter)
	if err != nil {
		return nil, err
	}

	return &requestParameter, nil
}

func mandatoryFieldsCheck(requestParameter models.RequestParameter) error {

	if requestParameter.AccessToken == nil {
		return errors.ErrEmptyAccessTokenInRequest
	}

	if len(*requestParameter.AccessToken) < 20 {
		return errors.ErrInvalidAccessTokenInRequest
	}

	return nil
}

func ParseResponse(getUserOutput *cognitoidentityprovider.GetUserOutput) *models.Response {

	userAttributes := []*models.UserAttribute{}
	for _, u := range getUserOutput.UserAttributes {
		userAttribute := models.UserAttribute{
			Name:  u.Name,
			Value: u.Value,
		}

		userAttributes = append(userAttributes, &userAttribute)
	}

	response := models.Response{
		Username:       getUserOutput.Username,
		UserAttributes: userAttributes,
	}

	return &response
}
