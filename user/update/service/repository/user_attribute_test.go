package repository

import (
	"testing"
	"update-user/service/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) UpdateUserAttributes(input *cognitoidentityprovider.UpdateUserAttributesInput) (*cognitoidentityprovider.UpdateUserAttributesOutput, error) {
	cogInitiateAuthOutput := cognitoidentityprovider.UpdateUserAttributesOutput{}
	return &cogInitiateAuthOutput, nil
}

func TestUpdateUserAttributes(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	accessToken := "abc@bc.com"
	firstName := "firstName"
	lastName := "lastName"

	userAttributes := models.UserAttribute{
		FirstName:  &firstName,
		LastName:   &lastName,
	}

	req := models.RequestParameter{
		AccessToken:    &accessToken,
		UserAttributes: &userAttributes,
	}

	err := CognitoUpdateUserAttributes(mockCC, &req)

	assert.NoError(err)
}
