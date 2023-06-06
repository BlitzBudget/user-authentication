package repository

import (
	"global-signout/service/models"
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
	accessToken := "12345678912345678912"
	req := models.RequestParameter{
		Email:       &email,
		AccessToken: &accessToken,
	}

	err := CognitoGlobalSignout(mockCC, &req)

	assert.NoError(err)
}
