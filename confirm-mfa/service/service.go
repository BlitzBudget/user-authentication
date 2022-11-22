package service

import (
	"confirm-mfa/service/models"
	"confirm-mfa/service/repository"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ConfirmMFA(body *string) (*models.Response, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	requestObject, err := repository.ParseRequest(body)
	if err != nil {
		fmt.Printf("Got error marshalling new item: %v", err)
		return nil, err
	}

	response, err := repository.CognitoConfirmMFA(cognitoClient, requestObject)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("SignupUser: There was an error logging the user %v", string(respAsBytes))
		return nil, err
	}

	return response, nil
}
