package repository

import (
	"fmt"
	"login/service/config"
	"login/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoLogin(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) (*string, error) {
	secretHash := ComputeSecretHash(req.Email)

	authInput := cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String(config.AuthFlow),
		AuthParameters: map[string]*string{
			"USERNAME":    aws.String(*req.Email),
			"PASSWORD":    aws.String(*req.Password),
			"SECRET_HASH": aws.String(secretHash),
		},
		ClientId: aws.String(config.ClientId),
	}

	resp, err := cognitoClient.InitiateAuth(&authInput)
	fmt.Printf("The response of the Initiate Auth is %v", resp)

	return resp.Session, err
}
