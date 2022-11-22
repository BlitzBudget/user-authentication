package repository

import (
	"confirm-signup/service/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com", "confirmation_code": "123456"}`

	requestParameter, err := ParseRequest(&body)

	assert.NoError(err)

	// assert for not nil (good when you expect something)
	if assert.NotNil(requestParameter) {
		assert.Equal(*requestParameter.Email, "abc@bc.com")
		assert.Equal(*requestParameter.ConfirmationCode, "123456")
	}
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

func TestEmptyConfirmationCode(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal(err.Error(), "the confirmation code is empty")
	assert.Nil(requestParameter)
}

func TestEmptyEmail(t *testing.T) {
	assert := assert.New(t)
	body := `{"confirmation_code": "123456"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal(err.Error(), "the email id is empty")
	assert.Nil(requestParameter)
}

func TestInvalidConfirmationCode(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com", "confirmation_code": "225", "new_password": "12345678"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Nil(requestParameter)
	assert.Equal(err.Error(), "the confirmation code entered is invalid")
}

func TestInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc", "confirmation_code": "23456", "new_password": "12345678"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Nil(requestParameter)
	assert.Equal(err.Error(), "the email entered is invalid")
}
