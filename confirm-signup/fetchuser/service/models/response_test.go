package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpResponse(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"username\": \"123456789\"}"

	fetchUserResponse := FetchUserResponse{}
	err := json.Unmarshal([]byte(body), &fetchUserResponse)

	assert.NoError(err)

	assert.Equal(*fetchUserResponse.Username, "123456789")
}
