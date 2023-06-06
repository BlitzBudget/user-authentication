package permissions

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
	email := "abc@com.com"
	cogInitiateAuthOutput := cognitoidentityprovider.GetUserOutput{
		Username: &email,
	}
	return &cogInitiateAuthOutput, nil
}

func TestUserHasPermissions(t *testing.T) {
	assert := assert.New(t)

	accessToken := "accessToken"
	mockCC := &mockCognitoClient{}
	userEmail := "abc@com.com"

	userHasPermissions, err := UserHasPermissions(&accessToken, mockCC, &userEmail)

	assert.NoError(err)
	assert.NotNil(userHasPermissions)
	assert.True(userHasPermissions)
}
