package repository

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) GetUser(input *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error) {
	email := "email"
	cogInitiateAuthOutput := cognitoidentityprovider.GetUserOutput{
		Username: &email,
	}
	return &cogInitiateAuthOutput, nil
}

func TestGetUser(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	accessToken := "abc@bc.com"

	response, err := CognitoGetUser(mockCC, &accessToken)

	assert.NoError(err)

	if assert.NotNil(response) {
		assert.Equal(*response.Username, "email")
	}
}
