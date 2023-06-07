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

func (m *mockCognitoClient) AdminUserGlobalSignOut(input *cognitoidentityprovider.AdminUserGlobalSignOutInput) (*cognitoidentityprovider.AdminUserGlobalSignOutOutput, error) {
	cogInitiateAuthOutput := cognitoidentityprovider.AdminUserGlobalSignOutOutput{}
	return &cogInitiateAuthOutput, nil
}

func TestCognitoGlobalSignout(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	email := "abc@bc.com"

	err := CognitoGlobalSignout(mockCC, &email)

	assert.NoError(err)
}
