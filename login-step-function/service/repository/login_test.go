package repository

import (
	"login/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) InitiateAuth(input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	session := "not empty"
	cogInitiateAuthOutput := cognitoidentityprovider.InitiateAuthOutput{
		Session: &session,
	}
	return &cogInitiateAuthOutput, nil
}

func TestCognitoLogin(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	email := "abc@bc.com"
	password := "12345678"
	req := models.RequestParameter{
		Email:    &email,
		Password: &password,
	}

	session, err := CognitoLogin(mockCC, &req)

	assert.NoError(err)
	assert.Equal(*session, "not empty")
}
