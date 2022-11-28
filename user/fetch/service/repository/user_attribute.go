package repository

import (
	"encoding/json"
	"fetch-user/service/models"
	"fmt"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoGetUser(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) (*models.Response, error) {

	fpInput := cognitoidentityprovider.GetUserInput{
		AccessToken: req.AccessToken,
	}

	fpOutput, err := cognitoClient.GetUser(&fpInput)
	respAsBytes, _ := json.Marshal(fpOutput)
	fmt.Printf("The response of the Forgot Password is %v", string(respAsBytes))

	if err != nil {
		return nil, err
	}

	return ParseResponse(fpOutput), nil
}
