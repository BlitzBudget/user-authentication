package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ErrEmptyEmailInRequest.Error(), "the email id is empty")
	assert.Equal(ErrInvalidEmailInRequest.Error(), "the email entered is invalid")
	assert.Equal(ErrEmptyPasswordInRequest.Error(), "the password provided is empty")
	assert.Equal(ErrMinimumLengthRequiredPassword.Error(), "the password provided must have a minimum length of 8 characters")
}
