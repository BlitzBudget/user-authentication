package service

import (
	"fmt"
	"login/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func LoginUser(body *string) (*string, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	mo, err := repository.ParseRequest(body)
	if err != nil {
		fmt.Printf("Got error marshalling new item: %v", err)
		return nil, err
	}

	session, err := repository.CognitoLogin(cognitoClient, mo)
	if err != nil {
		fmt.Printf("LoginUser: There was an error logging the user %v", err)
		return nil, err
	}

	return session, nil
}
