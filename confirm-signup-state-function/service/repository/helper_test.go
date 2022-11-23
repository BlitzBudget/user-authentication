package repository

import (
	"confirm-signup/service/config"
	"confirm-signup/service/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	email := "abc@anc.com"
	request := models.RequestParameter{
		Email: &email,
	}

	err := MandatoryFieldsCheck(&request)

	assert.Error(err)
	assert.Equal(err.Error(), "the confirmation code is empty")
}

func TestEmptyEmail(t *testing.T) {
	assert := assert.New(t)
	confirmationCode := "226546"
	request := models.RequestParameter{
		ConfirmationCode: &confirmationCode,
	}

	err := MandatoryFieldsCheck(&request)

	assert.Error(err)
	assert.Equal(err.Error(), "the email id is empty")
}

func TestInvalidConfirmationCode(t *testing.T) {
	assert := assert.New(t)
	confirmationCode := "226"
	email := "abc@abc.com"
	request := models.RequestParameter{
		ConfirmationCode: &confirmationCode,
		Email:            &email,
	}

	err := MandatoryFieldsCheck(&request)

	assert.Error(err)
	assert.Equal(err.Error(), "the confirmation code entered is invalid")
}

func TestInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	email := "abc"
	confirmationCode := "123456"
	request := models.RequestParameter{
		Email:            &email,
		ConfirmationCode: &confirmationCode,
	}

	err := MandatoryFieldsCheck(&request)

	assert.Error(err)
	assert.Equal(err.Error(), "the email entered is invalid")
}
