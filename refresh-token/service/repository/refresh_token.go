package repository

import (
	"fmt"
	"refresh-token/service/config"
	"refresh-token/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoInitiateAuth(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) (*models.Response, error) {
	secretHash := ComputeSecretHash(req.UserId)
	initiateAuthInput := cognitoidentityprovider.InitiateAuthInput{
		ClientId: &config.ClientId,
		AuthFlow: aws.String(config.AuthFlow),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": aws.String(*req.RefreshToken),
			"SECRET_HASH":   aws.String(secretHash),
		},
	}

	cfgOutput, err := cognitoClient.InitiateAuth(&initiateAuthInput)
	fmt.Printf("The response of the Refresh Token is %v \n", cfgOutput.String())

	return ParseResponse(cfgOutput, err)
}
