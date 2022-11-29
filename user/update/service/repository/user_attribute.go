package repository

import (
	"encoding/json"
	"fmt"
	"update-user/service/config"
	"update-user/service/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoUpdateUserAttributes(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) error {
	userAttributes := []*cognitoidentityprovider.AttributeType{}

	userAttributes = updateFirstName(userAttributes, req)
	userAttributes = updateLastName(userAttributes, req)

	fpInput := cognitoidentityprovider.UpdateUserAttributesInput{
		AccessToken:    req.AccessToken,
		UserAttributes: userAttributes,
	}

	fpOutput, err := cognitoClient.UpdateUserAttributes(&fpInput)
	respAsBytes, _ := json.Marshal(fpOutput)
	fmt.Printf("The response of the Forgot Password is %v \n", string(respAsBytes))

	return err
}

func updateFirstName(userAttributes []*cognitoidentityprovider.AttributeType, request *models.RequestParameter) []*cognitoidentityprovider.AttributeType {

	if request.UserAttributes == nil || request.UserAttributes.FirstName == nil {
		return nil
	}

	userAttribute := cognitoidentityprovider.AttributeType{
		Name:  &config.FirstNameUserAttribute,
		Value: request.UserAttributes.FirstName,
	}

	userAttributes = append(userAttributes, &userAttribute)
	return userAttributes
}

func updateLastName(userAttributes []*cognitoidentityprovider.AttributeType, request *models.RequestParameter) []*cognitoidentityprovider.AttributeType {

	if request.UserAttributes == nil || request.UserAttributes.LastName == nil {
		return nil
	}

	userAttribute := cognitoidentityprovider.AttributeType{
		Name:  &config.LastNameUserAttribute,
		Value: request.UserAttributes.LastName,
	}

	userAttributes = append(userAttributes, &userAttribute)
	return userAttributes
}
