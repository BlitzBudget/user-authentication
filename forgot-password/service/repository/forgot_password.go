package repository

import (
	"encoding/json"
	"fmt"
	"forgot-password/service/config"
	"forgot-password/service/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoForgotPassword(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) error {
	secretHash := ComputeSecretHash(req.Email)
	fpInput := cognitoidentityprovider.ForgotPasswordInput{
		ClientId:   &config.ClientId,
		SecretHash: &secretHash,
		Username:   req.Email,
	}

	fpOutput, err := cognitoClient.ForgotPassword(&fpInput)
	respAsBytes, _ := json.Marshal(fpOutput)
	fmt.Printf("The response of the Forgot Password is %v \n", string(respAsBytes))

	return err
}
