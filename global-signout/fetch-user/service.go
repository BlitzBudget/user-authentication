package service

import (
	"encoding/json"
	"fmt"
	"global-signout/fetch-user/models"
	"global-signout/fetch-user/repository"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func GetUser(accessToken *string, cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI) (*models.FetchUserResponse, error) {

	response, err := repository.CognitoGetUser(cognitoClient, accessToken)
	if err != nil {
		respAsBytes, _ := json.Marshal(err)
		fmt.Printf("SignupUser: There was an error logging the user %v \n", string(respAsBytes))
		return nil, err
	}

	return response, nil
}
