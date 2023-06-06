package repository

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	assert := assert.New(t)

	name := "name"
	value := "value"
	name1 := "name1"
	value1 := "value1"
	name2 := "name2"
	value2 := "value2"
	name3 := "name3"
	value3 := "value3"

	attributeTypes := []*cognitoidentityprovider.AttributeType{
		{
			Name:  &name,
			Value: &value,
		},
		{
			Name:  &name1,
			Value: &value1,
		},
		{
			Name:  &name2,
			Value: &value2,
		},
		{
			Name:  &name3,
			Value: &value3,
		},
	}

	userOutput := cognitoidentityprovider.GetUserOutput{
		UserAttributes: attributeTypes,
	}

	response := ParseResponse(&userOutput)
	if assert.NotNil(response.UserAttributes) {
		assert.Equal(*response.UserAttributes[0].Name, name)
		assert.Equal(*response.UserAttributes[0].Value, value)

		assert.Equal(*response.UserAttributes[1].Name, name1)
		assert.Equal(*response.UserAttributes[1].Value, value1)

		assert.Equal(*response.UserAttributes[2].Name, name2)
		assert.Equal(*response.UserAttributes[2].Value, value2)

		assert.Equal(*response.UserAttributes[3].Name, name3)
		assert.Equal(*response.UserAttributes[3].Value, value3)
	}
}
