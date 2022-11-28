package service

import (
	fetchuserService "confirm-signup/fetchuser/service"
	"confirm-signup/http/helper"
	"confirm-signup/login/service"
	"confirm-signup/service/models"
	"confirm-signup/service/repository"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ConfirmSignUp(body *string) (*models.HttpResponse, error) {
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

	err = repository.CognitoConfirmSignUp(cognitoClient, requestObject)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("ConfirmSignupUser: There was an error confirming the signup for the user %v", string(respAsBytes))
		return nil, err
	}

	loginResponseModel, err := service.LoginUser(requestObject)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("Loginuser: There was an error logging the user %v", string(respAsBytes))
		return nil, err
	}
	httpResponse := helper.ParseResponse(loginResponseModel)

	fetchUserResponse, err := fetchuserService.GetUser(loginResponseModel.AuthenticationResult.AccessToken, cognitoClient)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("FetchUser: There was an error fetching the user attributes %v", string(respAsBytes))
		return nil, err
	}
	httpResponse = helper.ParseFetchUserResponse(fetchUserResponse, httpResponse)

	return httpResponse, nil
}
