package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ErrEmptyEmailInRequest.Error(), "the email id is empty")
	assert.Equal(ErrInvalidEmailInRequest.Error(), "the email entered is invalid")
	assert.Equal(ErrEmptyAccessTokenInRequest.Error(), "the access token provided is empty")
	assert.Equal(ErrMinimumLengthRequiredAccessToken.Error(), "the access token provided is invalid")
}
