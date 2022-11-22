package repository

import (
	"confirm-forgot-password/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) ConfirmForgotPassword(input *cognitoidentityprovider.ConfirmForgotPasswordInput) (*cognitoidentityprovider.ConfirmForgotPasswordOutput, error) {
	cogInitiateAuthOutput := cognitoidentityprovider.ConfirmForgotPasswordOutput{}
	return &cogInitiateAuthOutput, nil
}

func TestConfirmForgotPassword(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	email := "abc@bc.com"
	confirmationCode := "123456"
	password := "Password"

	req := models.RequestParameter{
		Email:            &email,
		ConfirmationCode: &confirmationCode,
		Password:         &password,
	}

	err := CognitoConfirmForgotPassword(mockCC, &req)

	assert.NoError(err)
}
