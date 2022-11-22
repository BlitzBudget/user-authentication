package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpResponse(t *testing.T) {
	assert := assert.New(t)
	body := "{ \"message\": \"123456789\"}"

	httpResponse := HttpResponse{}
	err := json.Unmarshal([]byte(body), &httpResponse)

	assert.NoError(err)

	assert.Equal(*httpResponse.Message, "123456789")
}
