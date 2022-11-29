package repository

import (
	"confirm-forgot-password/service/config"
	"confirm-forgot-password/service/models"
	"fmt"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoConfirmForgotPassword(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) error {
	secretHash := ComputeSecretHash(req.Email)
	confirmForgotPasswordInput := cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         &config.ClientId,
		SecretHash:       &secretHash,
		Username:         req.Email,
		ConfirmationCode: req.ConfirmationCode,
		Password:         req.Password,
	}

	cfgOutput, err := cognitoClient.ConfirmForgotPassword(&confirmForgotPasswordInput)
	fmt.Printf("The response of the Signup is %v \n", cfgOutput.String())

	return err
}
