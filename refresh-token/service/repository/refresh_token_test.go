package repository

import (
	"refresh-token/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) InitiateAuth(input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	challengeName := "Challenge Name"
	accessToken := "accessToken"
	expiresIn := int64(1234)
	idToken := "idToken"
	deviceGroupKey := "deviceGroupKey"
	deviceKey := "deviceKey"
	refreshToken := "refreshToken"
	tokenType := "tokenType"
	cogInitiateAuthOutput := cognitoidentityprovider.InitiateAuthOutput{
		ChallengeName: &challengeName,
		AuthenticationResult: &cognitoidentityprovider.AuthenticationResultType{
			AccessToken:  &accessToken,
			ExpiresIn:    &expiresIn,
			IdToken:      &idToken,
			RefreshToken: &refreshToken,
			TokenType:    &tokenType,
			NewDeviceMetadata: &cognitoidentityprovider.NewDeviceMetadataType{
				DeviceGroupKey: &deviceGroupKey,
				DeviceKey:      &deviceKey,
			},
		},
	}
	return &cogInitiateAuthOutput, nil
}

func TestInitiateAuth(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	user_id := "abc@bc.com"
	refreshToken := "123456"

	req := models.RequestParameter{
		UserId:       &user_id,
		RefreshToken: &refreshToken,
	}

	response, err := CognitoInitiateAuth(mockCC, &req)

	assert.NoError(err)
	assert.Equal(*response.ChallangeName, "Challenge Name")
	assert.Equal(*response.AccessToken, "accessToken")
	assert.Equal(*response.ExpiresIn, int64(1234))
	assert.Equal(*response.IdToken, "idToken")
}
