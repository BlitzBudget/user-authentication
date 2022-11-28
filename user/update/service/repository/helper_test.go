package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert := assert.New(t)
	body := `{"access_token": "123456789012345678900"}`

	requestParameter, err := ParseRequest(&body)

	assert.NoError(err)

	// assert for not nil (good when you expect something)
	if assert.NotNil(requestParameter) {
		assert.Equal(*requestParameter.AccessToken, "123456789012345678900")
	}
}

func TestEmptyAccessToken(t *testing.T) {
	assert := assert.New(t)
	body := `{}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal(err.Error(), "the access token is empty")

	// assert for not nil (good when you expect something)
	assert.Nil(requestParameter)
}

func TestInvalidAccessToken(t *testing.T) {
	assert := assert.New(t)
	body := `{ "access_token": "123456789456123"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal(err.Error(), "the access token entered is invalid")

	// assert for not nil (good when you expect something)
	assert.Nil(requestParameter)
}
