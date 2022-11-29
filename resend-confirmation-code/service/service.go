package service

import (
	"encoding/json"
	"fmt"
	"resend-confirmation-code/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ResendConfirmationCode(body *string) error {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	requestObject, err := repository.ParseRequest(body)
	if err != nil {
		fmt.Printf("Got error marshalling new item: %v \n", err)
		return err
	}

	err = repository.CognitoResendConfirmationCode(cognitoClient, requestObject)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("SignupUser: There was an error logging the user %v \n", string(respAsBytes))
		return err
	}

	return nil
}
