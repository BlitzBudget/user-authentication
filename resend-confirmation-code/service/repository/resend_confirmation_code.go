package repository

import (
	"encoding/json"
	"fmt"
	"resend-confirmation-code/service/config"
	"resend-confirmation-code/service/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoResendConfirmationCode(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) error {
	secretHash := ComputeSecretHash(req.Email)
	rccInput := cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId:   &config.ClientId,
		SecretHash: &secretHash,
		Username:   req.Email,
	}

	rccOutput, err := cognitoClient.ResendConfirmationCode(&rccInput)
	respAsBytes, _ := json.Marshal(rccOutput)
	fmt.Printf("The response of the Resend Confirmation Code is %v \n", string(respAsBytes))

	return err
}
