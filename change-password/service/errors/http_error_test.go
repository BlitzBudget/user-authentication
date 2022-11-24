package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ErrEmptyAccessTokenInRequest.Error(), "the access token is empty")
	assert.Equal(ErrEmptyPreviousPasswordInRequest.Error(), "the previous password provided is empty")
	assert.Equal(ErrMinimumLengthRequiredPreviousPassword.Error(), "the previous password provided must have a minimum length of 8 characters")
	assert.Equal(ErrEmptyProposedPasswordInRequest.Error(), "the proposed password provided is empty")
	assert.Equal(ErrMinimumLengthRequiredProposedPassword.Error(), "the proposed password provided must have a minimum length of 8 characters")
}
