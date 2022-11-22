package repository

import (
	"signup/service/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com", "password": "123456789", "user_role_id": "UserRole#12345", "organization_id": "Organization#12345", "phone_number": "666666666", "name": "MockName", "last_name": "MockLastName"}`

	requestParameter, err := ParseRequest(&body)

	assert.NoError(err)

	// assert for not nil (good when you expect something)
	if assert.NotNil(requestParameter) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("abc@bc.com", *requestParameter.Email)
		assert.Equal("123456789", *requestParameter.Password)
		assert.Equal("MockName", *requestParameter.Name)
	}

}

func TestEmptyEmail(t *testing.T) {
	assert := assert.New(t)
	body := `{ "password": "123456789"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the email id is empty", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestInvalidEmail(t *testing.T) {
	assert := assert.New(t)
	body := "{\"email\":\"\", \"password\": \"123456789\"}"

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
	body := "{\"email\":\"abc@bc.com\", \"password\": \"1234568\"}"

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the password provided must have a minimum length of 8 characters", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestEmptyName(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com", "password": "123456789", "user_role_id": "UserRole#12345", "organization_id": "Organization#12345", "phone_number": "666666666", "last_name": "MockLastName"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the name should have only characters in the request", err.Error())
	// assert for nil
	assert.Nil(requestParameter)
}

func TestNumberedName(t *testing.T) {
	assert := assert.New(t)
	body := `{"email": "abc@bc.com", "password": "123456789", "user_role_id": "UserRole#12345", "organization_id": "Organization#12345", "phone_number": "666666666", "name": "MockName2", "last_name": "MockLastName"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal("the name should have only characters in the request", err.Error())
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