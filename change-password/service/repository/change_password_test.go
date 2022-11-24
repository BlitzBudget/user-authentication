package repository

import (
	"change-password/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) ChangePassword(input *cognitoidentityprovider.ChangePasswordInput) (*cognitoidentityprovider.ChangePasswordOutput, error) {
	cogInitiateAuthOutput := cognitoidentityprovider.ChangePasswordOutput{}
	return &cogInitiateAuthOutput, nil
}

func TestChangePassword(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	accessToken := "accessToken"
	previousPassword := "previousPassword"
	proposedPassword := "proposedPassword"

	req := models.RequestParameter{
		AccessToken:      &accessToken,
		PreviousPassword: &previousPassword,
		ProposedPassword: &proposedPassword,
	}

	err := CognitoChangePassword(mockCC, &req)

	assert.NoError(err)
}
