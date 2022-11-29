package repository

import (
	"confirm-signup/service/config"
	"confirm-signup/service/models"
	"fmt"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoConfirmSignUp(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) error {
	secretHash := ComputeSecretHash(req.Email)
	ConfirmSignUpInput := cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         &config.ClientId,
		SecretHash:       &secretHash,
		Username:         req.Email,
		ConfirmationCode: req.ConfirmationCode,
	}

	cfgOutput, err := cognitoClient.ConfirmSignUp(&ConfirmSignUpInput)
	fmt.Printf("The response of the Signup is %v \n", cfgOutput.String())

	return err
}
