package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert := assert.New(t)
	body := "{\"email\":\"abc@bc.com\", \"access_token\": \"12345678912345678912345\"}"

	requestParameter, err := ParseRequest(&body)

	assert.NoError(err)

	// assert for not nil (good when you expect something)
	if assert.NotNil(requestParameter) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("abc@bc.com", *requestParameter.Email)
		assert.Equal("12345678912345678912345", *requestParameter.AccessToken)
	}

}

func TestEmptyEmail(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"access_token\": \"12345678912345678912345\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the email id is empty", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	body := "{\"email\":\"\", \"access_token\": \"12345678912345678912345\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the email entered is invalid", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestEmptyAccessToken(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"email\": \"abc@bc.com\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the access token provided is empty", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestInvalidAccessToken(t *testing.T) {
	assert := assert.New(t)
	body := "{\"email\":\"abc@bc.com\", \"access_token\": \"1234568\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the access token provided is invalid", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}
