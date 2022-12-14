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

func (m *mockCognitoClient) AdminInitiateAuth(input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
	session := "not empty"
	cogAdminInitiateAuthOutput := cognitoidentityprovider.AdminInitiateAuthOutput{
		Session: &session,
	}
	return &cogAdminInitiateAuthOutput, nil
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

	resp, err := CognitoLogin(mockCC, &req)

	assert.NoError(err)
	assert.NotNil(resp)
}
