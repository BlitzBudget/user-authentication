package service

import (
	"fmt"
	"global-signout/permissions"
	"global-signout/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func GlobalSignoutUser(body *string) error {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	mo, err := repository.ParseRequest(body)
	if err != nil {
		fmt.Printf("Got error marshalling new item: %v \n", err)
		return err
	}

	userHasPermission, err := permissions.UserHasPermissions(mo.AccessToken, cognitoClient, mo.Email)
	if err != nil || !userHasPermission {
		fmt.Printf("The User does not have permissions to perform this acction %v \n", err)
		return err
	}

	err = repository.CognitoGlobalSignout(cognitoClient, mo)
	if err != nil {
		fmt.Printf("globalSignoutUser: There was an error logging the user %v \n", err)
		return err
	}

	return nil
}
