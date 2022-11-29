package service

import (
	"encoding/json"
	"fmt"
	"signup/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func SignupUser(body *string) error {
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

	err = repository.CognitoSignup(cognitoClient, mo)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("SignupUser: There was an error logging the user %v \n", string(respAsBytes))
		return err
	}

	return nil
}
