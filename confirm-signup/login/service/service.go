package service

import (
	"confirm-signup/login/service/repository"
	"confirm-signup/service/models"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func LoginUser(requestObject *models.RequestParameter) (*models.LoginResponseModel, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	resp, err := repository.CognitoLogin(cognitoClient, requestObject)
	if err != nil {
		fmt.Printf("LoginUser: There was an error logging the user %v \n", err)
		return nil, err
	}

	return resp, nil
}
