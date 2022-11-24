package repository

import (
	"confirm-signup/login/service/config"
	"confirm-signup/service/models"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoLogin(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) (*models.LoginResponseModel, error) {
	secretHash := ComputeSecretHash(req.Email)

	authInput := cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow: aws.String(config.AuthFlow),
		AuthParameters: map[string]*string{
			"USERNAME":    aws.String(*req.Email),
			"PASSWORD":    aws.String(*req.Password),
			"SECRET_HASH": aws.String(secretHash),
		},
		ClientId:   aws.String(config.ClientId),
		UserPoolId: aws.String(config.UserPoolId),
	}

	resp, err := cognitoClient.AdminInitiateAuth(&authInput)
	fmt.Printf("The response of the Initiate Auth is %v", resp)
	if err != nil {
		return nil, err
	}

	return ParseResponse(resp), nil
}
