package repository

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	assert := assert.New(t)

	queryOutput := dynamodb.QueryOutput{
		Items: nil,
	}

	got, err := ParseResponse(&queryOutput)

	if assert.Error(err) {
		assert.Equal(err.Error(), "no Items found")
	}
	assert.Nil(got)
}
