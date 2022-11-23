package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"email\": \"abc@bc.com\"}"

	request := RequestParameter{}
	err := json.Unmarshal([]byte(body), &request)

	assert.NoError(err)
	assert.Equal(*request.Email, "abc@bc.com")
}
