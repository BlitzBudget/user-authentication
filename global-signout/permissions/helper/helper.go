package helper

import (
	fetchUserModels "global-signout/fetch-user/models"
)

func FetchEmailFromUser(user *fetchUserModels.FetchUserResponse) *string {
	if user.Username != nil {
		return user.Username
	}

	return nil
}
