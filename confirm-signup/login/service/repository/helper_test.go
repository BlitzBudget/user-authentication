package repository

import (
	"confirm-signup/login/service/config"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	assert := assert.New(t)

	session := "session"
	accessToken := "accessToken"
	refreshToken := "refreshToken"
	idToken := "idToken"
	authResultType := cognitoidentityprovider.AuthenticationResultType{
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
		IdToken:      &idToken,
	}

	adminInitAuthOut := cognitoidentityprovider.AdminInitiateAuthOutput{
		Session:              &session,
		AuthenticationResult: &authResultType,
	}

	loginResponse := ParseResponse(&adminInitAuthOut)

	assert.NotNil(loginResponse)
	assert.Equal(*loginResponse.Session, session)
	assert.Equal(*loginResponse.AuthenticationResult.AccessToken, accessToken)
	assert.Equal(*loginResponse.AuthenticationResult.IdToken, idToken)
	assert.Equal(*loginResponse.AuthenticationResult.RefreshToken, refreshToken)
}

func TestParseRequest(t *testing.T) {
	assert := assert.New(t)
	body := "{\"email\":\"abc@bc.com\", \"new_password\": \"12345678\"}"

	requestParameter, err := ParseRequest(&body)

	assert.NoError(err)

	// assert for not nil (good when you expect something)
	if assert.NotNil(requestParameter) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("abc@bc.com", *requestParameter.Email)
		assert.Equal("12345678", *requestParameter.Password)
	}

}

func TestEmptyEmail(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"new_password\": \"12345678\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the email id is empty", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	body := "{\"email\":\"\", \"new_password\": \"12345678\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the email entered is invalid", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestEmptyPassword(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"email\": \"abc@bc.com\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the password provided is empty", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestInvalidPassword(t *testing.T) {
	assert := assert.New(t)
	body := "{\"email\":\"abc@bc.com\", \"new_password\": \"1234568\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the password provided must have a minimum length of 8 characters", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestComputeHashMap(t *testing.T) {
	assert := assert.New(t)
	email := "abc@bc.com"
	config.ClientId = "clientID"
	config.ClientSecret = "clientSecret"

	hashedSecret := ComputeSecretHash(&email)

	assert.NotNil(hashedSecret)
	assert.Equal("aMIWMCPgXsRusULKA/OJ1AfztPK2JgBykLKnCZ1k/yQ=", hashedSecret)
}
