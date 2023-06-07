package repository

import (
	"fmt"
	"global-signout/service/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoGlobalSignout(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, email *string) error {

	authInput := cognitoidentityprovider.AdminUserGlobalSignOutInput{
		UserPoolId: aws.String(config.UserPoolId),
		Username:   email,
	}

	resp, err := cognitoClient.AdminUserGlobalSignOut(&authInput)
	fmt.Printf("The response of the AdminUserGlobalSignOut is %v \n", resp)

	return err
}
