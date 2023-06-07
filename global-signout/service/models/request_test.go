package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"access_token\": \"12345678\"}"

	request := RequestParameter{}
	err := json.Unmarshal([]byte(body), &request)

	assert.NoError(err)
	assert.Equal(*request.AccessToken, "12345678")
}
