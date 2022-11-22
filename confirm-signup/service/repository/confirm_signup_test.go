package repository

import (
	"confirm-signup/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) ConfirmSignUp(input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
	cogInitiateAuthOutput := cognitoidentityprovider.ConfirmSignUpOutput{}
	return &cogInitiateAuthOutput, nil
}

func TestConfirmSignUp(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	email := "abc@bc.com"
	confirmationCode := "123456"

	req := models.RequestParameter{
		Email:            &email,
		ConfirmationCode: &confirmationCode,
	}

	err := CognitoConfirmSignUp(mockCC, &req)

	assert.NoError(err)
}
