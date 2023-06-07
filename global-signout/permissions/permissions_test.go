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

	name := "custom:user_role_id"
	value := "userRoleObtained"
	name1 := "email"
	value1 := "abc@com.com"
	name2 := "name2"
	value2 := "value2"
	name3 := "name3"
	value3 := "value3"
	username := "291303e4-cbed-4904-8592-3f28503e0f4c"

	attributeTypes := []*cognitoidentityprovider.AttributeType{
		{
			Name:  &name,
			Value: &value,
		},
		{
			Name:  &name1,
			Value: &value1,
		},
		{
			Name:  &name2,
			Value: &value2,
		},
		{
			Name:  &name3,
			Value: &value3,
		},
	}

	cogInitiateAuthOutput := cognitoidentityprovider.GetUserOutput{
		Username:       &username,
		UserAttributes: attributeTypes,
	}
	return &cogInitiateAuthOutput, nil
}

func TestUserHasPermissions(t *testing.T) {
	assert := assert.New(t)

	accessToken := "accessToken"
	mockCC := &mockCognitoClient{}

	userHasPermissions, err, email := UserHasPermissions(&accessToken, mockCC)

	assert.NoError(err)
	assert.NotNil(userHasPermissions)
	assert.True(userHasPermissions)
	assert.Equal(*email, "abc@com.com")
}
