package service

import (
	"errors"
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

	userHasPermission, err, email := permissions.UserHasPermissions(mo.AccessToken, cognitoClient)
	if err != nil {
		fmt.Printf("The User does not have permissions to perform this acction %v \n", err)
		return err
	}
	if !userHasPermission {
		err = errors.New("the User does not have permission to perform the global signout")
		fmt.Printf("%v \n", err.Error())
		return err
	}

	err = repository.CognitoGlobalSignout(cognitoClient, email)
	if err != nil {
		fmt.Printf("globalSignoutUser: There was an error logging the user %v \n", err)
		return err
	}

	return nil
}
