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

func (m *mockCognitoClient) AdminInitiateAuth(input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
	session := "not empty"
	cogInitiateAuthOutput := cognitoidentityprovider.AdminInitiateAuthOutput{
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

	loginResponseModel, err := CognitoLogin(mockCC, &req)

	assert.NoError(err)
	assert.NotNil(loginResponseModel)
	assert.NotNil(loginResponseModel.Session)
	assert.Equal(*loginResponseModel.Session, "not empty")
}
