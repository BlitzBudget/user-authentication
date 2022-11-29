package repository

import (
	"encoding/json"
	"fmt"
	"signup/service/config"
	"signup/service/models"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoSignup(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) error {
	date := time.Now().Format(time.RFC3339Nano)
	userId := config.UserIdPrefix + date
	lowercaseEmail := strings.ToLower(*req.Email)
	secretHash := ComputeSecretHash(&lowercaseEmail)
	locale := "EN"

	fmt.Printf("Creating user for %v \n", lowercaseEmail)
	attributeTypeArray := []*cognitoidentityprovider.AttributeType{
		{
			Name:  aws.String("email"),
			Value: &lowercaseEmail,
		},
		{
			Name:  aws.String("locale"),
			Value: &locale,
		},
		{
			Name:  aws.String("custom:exportFileFormat"),
			Value: &config.ExportFormat,
		},
		{
			Name:  aws.String("name"),
			Value: req.Name,
		},
		{
			Name:  aws.String("family_name"),
			Value: req.Name,
		},
		{
			Name:  aws.String("custom:financialPortfolioId"),
			Value: aws.String(userId),
		},
	}

	authInput := cognitoidentityprovider.SignUpInput{
		Username:       aws.String(lowercaseEmail),
		Password:       req.Password,
		ClientId:       aws.String(config.ClientId),
		SecretHash:     &secretHash,
		UserAttributes: attributeTypeArray,
	}

	resp, err := cognitoClient.SignUp(&authInput)
	respAsBytes, _ := json.Marshal(resp)
	fmt.Printf("The response of the Signup is %v \n", string(respAsBytes))

	return err
}
