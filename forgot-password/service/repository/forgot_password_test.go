package repository

import (
	"forgot-password/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) ForgotPassword(input *cognitoidentityprovider.ForgotPasswordInput) (*cognitoidentityprovider.ForgotPasswordOutput, error) {
	cogInitiateAuthOutput := cognitoidentityprovider.ForgotPasswordOutput{}
	return &cogInitiateAuthOutput, nil
}

func TestForgotPassword(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	email := "abc@bc.com"

	req := models.RequestParameter{
		Email: &email,
	}

	err := CognitoForgotPassword(mockCC, &req)

	assert.NoError(err)
}
