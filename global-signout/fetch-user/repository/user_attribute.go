package repository

import (
	"encoding/json"
	"fmt"
	"global-signout/fetch-user/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoGetUser(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, accessToken *string) (*models.FetchUserResponse, error) {

	fpInput := cognitoidentityprovider.GetUserInput{
		AccessToken: accessToken,
	}

	fpOutput, err := cognitoClient.GetUser(&fpInput)
	respAsBytes, _ := json.Marshal(fpOutput)
	fmt.Printf("The response of the Forgot Password is %v \n", string(respAsBytes))

	if err != nil {
		return nil, err
	}

	return ParseResponse(fpOutput), nil
}
