package repository

import (
	"signup/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) SignUp(input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error) {
	cogInitiateAuthOutput := cognitoidentityprovider.SignUpOutput{}
	return &cogInitiateAuthOutput, nil
}

func TestCognitoSignUp(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	email := "abc@bc.com"
	password := "12345678"
	name := "Name"
	lastName := "lastName"

	req := models.RequestParameter{
		Email:    &email,
		Password: &password,
		Name:     &name,
		LastName: &lastName,
	}

	err := CognitoSignup(mockCC, &req)

	assert.NoError(err)
}
