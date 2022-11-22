package repository

import (
	"resend-confirmation-code/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) ResendConfirmationCode(input *cognitoidentityprovider.ResendConfirmationCodeInput) (*cognitoidentityprovider.ResendConfirmationCodeOutput, error) {
	cogInitiateAuthOutput := cognitoidentityprovider.ResendConfirmationCodeOutput{}
	return &cogInitiateAuthOutput, nil
}

func TestResendConfirmationCode(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	email := "abc@bc.com"

	req := models.RequestParameter{
		Email: &email,
	}

	err := CognitoResendConfirmationCode(mockCC, &req)

	assert.NoError(err)
}
