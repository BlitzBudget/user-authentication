package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(ErrEmptyUserIdInRequest.Error(), "the user id is empty")

	assert.Equal(ErrEmptyRefreshTokenInRequest.Error(), "the refresh token is empty")
	assert.Equal(ErrInvalidRefreshTokenInRequest.Error(), "the refresh token entered is invalid")
}
