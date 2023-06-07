package helper

import (
	fetchUserModels "global-signout/fetch-user/models"
	"global-signout/service/config"
)

func FetchEmailFromUser(user *fetchUserModels.FetchUserResponse) *string {
	for _, ua := range user.UserAttributes {
		if *ua.Name == config.CognitoEmailAttribute {
			return ua.Value
		}
	}

	return nil
}
