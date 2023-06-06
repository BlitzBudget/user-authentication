package permissions

import (
	"fmt"
	fetchUser "global-signout/fetch-user"
	"global-signout/permissions/helper"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func UserHasPermissions(accessToken *string, cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, userEmail *string) (bool, error) {
	resp, err := fetchUser.GetUser(accessToken, cognitoClient)
	if err != nil {
		fmt.Printf("Unable to Verify AccessToken %v \n", err)
		return false, err
	}

	email := helper.FetchEmailFromUser(resp)
	if *email != *userEmail {
		fmt.Println("You do not have access to do a global signout")
		return false, nil
	}

	fmt.Println("You are authorized to do the global signout")
	return true, nil
}
