package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	assert := assert.New(t)
	body := `{ "access_token": "123468789", "previous_password": "12345678", "proposed_password": "321654879"}`

	request := RequestParameter{}
	err := json.Unmarshal([]byte(body), &request)

	assert.NoError(err)
	assert.Equal(*request.AccessToken, "123468789")
	assert.Equal(*request.PreviousPassword, "12345678")
	assert.Equal(*request.ProposedPassword, "321654879")
}
