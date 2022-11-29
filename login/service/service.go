package service

import (
	"encoding/json"
	"fmt"
	fetchuserService "login/fetchuser/service"
	"login/http/helper"
	"login/service/models"
	"login/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func LoginUser(body *string) (*models.HttpResponse, error) {
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

	loginResponseModel, err := repository.CognitoLogin(cognitoClient, mo)
	if err != nil {
		fmt.Printf("LoginUser: There was an error logging the user %v", err)
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

	// TODO Fetch all wallet

	return httpResponse, nil
}
