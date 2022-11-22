package repository

import (
	"confirm-mfa/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/stretchr/testify/assert"
)

type mockCognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

func (m *mockCognitoClient) RespondToAuthChallenge(input *cognitoidentityprovider.RespondToAuthChallengeInput) (*cognitoidentityprovider.RespondToAuthChallengeOutput, error) {
	challengeName := "Challenge Name"
	accessToken := "accessToken"
	expiresIn := int64(1234)
	idToken := "idToken"
	deviceGroupKey := "deviceGroupKey"
	deviceKey := "deviceKey"
	refreshToken := "refreshToken"
	tokenType := "tokenType"
	cogInitiateAuthOutput := cognitoidentityprovider.RespondToAuthChallengeOutput{
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

func TestConfirmMFA(t *testing.T) {
	assert := assert.New(t)

	mockCC := &mockCognitoClient{}
	email := "abc@bc.com"
	confirmationCode := "123456"

	req := models.RequestParameter{
		Email:            &email,
		ConfirmationCode: &confirmationCode,
	}

	response, err := CognitoConfirmMFA(mockCC, &req)

	assert.NoError(err)
	assert.Equal(*response.ChallangeName, "Challenge Name")
	assert.Equal(*response.AccessToken, "accessToken")
	assert.Equal(*response.ExpiresIn, int64(1234))
	assert.Equal(*response.IdToken, "idToken")
}
