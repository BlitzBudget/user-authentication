package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ErrEmptyAccessTokenInRequest.Error(), "the access token is empty")
	assert.Equal(ErrInvalidAccessTokenInRequest.Error(), "the access token entered is invalid")
}
