package service

import (
	"confirm-signup/service/models"
	"confirm-signup/service/repository"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ConfirmSignUp(request *models.RequestParameter) error {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	err := repository.MandatoryFieldsCheck(request)
	if err != nil {
		fmt.Printf("Got error marshalling new item: %v", err)
		return err
	}

	err = repository.CognitoConfirmSignUp(cognitoClient, request)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("SignupUser: There was an error logging the user %v", string(respAsBytes))
		return err
	}

	return nil
}
