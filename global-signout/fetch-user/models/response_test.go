package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpResponse(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"username\": \"123456789\"}"

	errorHttpResponse := FetchUserResponse{}
	err := json.Unmarshal([]byte(body), &errorHttpResponse)

	assert.NoError(err)

	assert.Equal(*errorHttpResponse.Username, "123456789")
}
