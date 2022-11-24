package repository

import (
	"change-password/service/models"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoChangePassword(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) error {
	fpInput := cognitoidentityprovider.ChangePasswordInput{
		AccessToken:      req.AccessToken,
		PreviousPassword: req.PreviousPassword,
		ProposedPassword: req.ProposedPassword,
	}

	fpOutput, err := cognitoClient.ChangePassword(&fpInput)
	respAsBytes, _ := json.Marshal(fpOutput)
	fmt.Printf("The response of the Forgot Password is %v", string(respAsBytes))

	return err
}
