package service

import (
	"encoding/json"
	"fmt"
	"login/fetchuser/service/models"
	"login/fetchuser/service/repository"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func GetUser(accessToken *string, cognitoClient *cognitoidentityprovider.CognitoIdentityProvider) (*models.FetchUserResponse, error) {

	response, err := repository.CognitoGetUser(cognitoClient, accessToken)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("SignupUser: There was an error logging the user %v", string(respAsBytes))
		return nil, err
	}

	return response, nil
}
