package repository

import (
	"fmt"
	"global-signout/service/config"
	"global-signout/service/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoGlobalSignout(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) error {

	authInput := cognitoidentityprovider.AdminUserGlobalSignOutInput{
		UserPoolId: aws.String(config.UserPoolId),
		Username:   req.Email,
	}

	resp, err := cognitoClient.AdminUserGlobalSignOut(&authInput)
	fmt.Printf("The response of the AdminUserGlobalSignOut is %v \n", resp)

	return err
}
