package helper

import (
	"global-signout/fetch-user/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchUserRoleFromUser(t *testing.T) {
	assert := assert.New(t)

	name := "custom:user_role_id"
	value := "userRoleObtained"
	name1 := "email"
	value1 := "email"
	name2 := "name2"
	value2 := "value2"
	name3 := "name3"
	value3 := "value3"
	email := "email"

	attributeTypes := []*models.UserAttribute{
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

	userResponse := models.FetchUserResponse{
		UserAttributes: attributeTypes,
		Username:       &email,
	}

	respEmail := FetchEmailFromUser(&userResponse)

	assert.Equal(*respEmail, value1)
}
