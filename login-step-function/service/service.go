package service

import (
	"fmt"
	"login/service/models"
	"login/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func LoginUser(requestMap *map[string]string) (*models.ResponseModel, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cognitoClient := cognitoidentityprovider.New(sess)

	email := (*requestMap)["email"]
	password := (*requestMap)["password"]
	requestObject := models.RequestParameter{
		Email:    &email,
		Password: &password,
	}

	resp, err := repository.CognitoLogin(cognitoClient, &requestObject)
	if err != nil {
		fmt.Printf("LoginUser: There was an error logging the user %v", err)
		return nil, err
	}

	return resp, nil
}
