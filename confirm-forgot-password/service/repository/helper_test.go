package repository

import (
	"confirm-forgot-password/service/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com", "confirmation_code": "123456", "new_password": "12345678"}`

	requestParameter, err := ParseRequest(&body)

	assert.NoError(err)

	// assert for not nil (good when you expect something)
	if assert.NotNil(requestParameter) {
		assert.Equal(*requestParameter.Email, "abc@bc.com")
		assert.Equal(*requestParameter.ConfirmationCode, "123456")
		assert.Equal(*requestParameter.Password, "12345678")
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

func TestEmptyEmailID(t *testing.T) {
	assert := assert.New(t)
	body := `{"confirmation_code": "123456", "new_password": "12345678"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Nil(requestParameter)
	assert.Equal(err.Error(), "the email id is empty")
}

func TestEmptyConfirmationCode(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com", "new_password": "12345678"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Nil(requestParameter)
	assert.Equal(err.Error(), "the confirmation code is empty")
}

func TestEmptyPassword(t *testing.T) {
	assert := assert.New(t)
	body := `{"confirmation_code": "123456", "email": "abc@bc.com"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Nil(requestParameter)
	assert.Equal(err.Error(), "the password provided is empty")
}

func TestInvalidConfirmationCode(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com", "confirmation_code": "234", "new_password": "12345678"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Nil(requestParameter)
	assert.Equal(err.Error(), "the confirmation code entered is invalid")
}

func TestInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc", "confirmation_code": "234567", "new_password": "12345678"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Nil(requestParameter)
	assert.Equal(err.Error(), "the email entered is invalid")
}
