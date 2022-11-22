package repository

import (
	"confirm-mfa/service/config"
	"confirm-mfa/service/models"
	"fmt"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

func CognitoConfirmMFA(cognitoClient cognitoidentityprovideriface.CognitoIdentityProviderAPI, req *models.RequestParameter) (*models.Response, error) {
	secretHash := ComputeSecretHash(req.Email)
	challengeResponses := make(map[string]*string)
	challengeResponses["SECRET_HASH"] = &secretHash
	challengeResponses["SMS_MFA_CODE"] = req.ConfirmationCode
	challengeResponses["USERNAME"] = req.Email
	ConfirmMFAInput := cognitoidentityprovider.RespondToAuthChallengeInput{
		ClientId:           &config.ClientId,
		ChallengeName:      &config.ChallangeName,
		ChallengeResponses: challengeResponses,
		Session:            req.Session,
	}

	cfgOutput, err := cognitoClient.RespondToAuthChallenge(&ConfirmMFAInput)
	fmt.Printf("The response of the Signup is %v", cfgOutput.String())

	return ParseResponse(cfgOutput, err)
}
