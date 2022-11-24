package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert := assert.New(t)
	body := `{ "access_token": "123468789", "previous_password": "12345678", "proposed_password": "321654879"}`

	requestParameter, err := ParseRequest(&body)

	assert.NoError(err)

	// assert for not nil (good when you expect something)
	if assert.NotNil(requestParameter) {
		assert.Equal(*requestParameter.AccessToken, "123468789")
		assert.Equal(*requestParameter.PreviousPassword, "12345678")
		assert.Equal(*requestParameter.ProposedPassword, "321654879")
	}
}
