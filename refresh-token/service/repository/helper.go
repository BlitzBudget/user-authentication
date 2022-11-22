package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"refresh-token/service/config"
	customErrors "refresh-token/service/errors"
	"refresh-token/service/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ParseResponse(acOutput *cognitoidentityprovider.InitiateAuthOutput, err error) (*models.Response, error) {
	var responseItems models.Response

	if err != nil {
		return nil, err
	}

	if acOutput == nil {
		msg := "parse Auth Response: Null Object Found"
		return nil, errors.New(msg)
	}

	responseItems.AccessToken = acOutput.AuthenticationResult.AccessToken
	responseItems.ExpiresIn = acOutput.AuthenticationResult.ExpiresIn
	responseItems.IdToken = acOutput.AuthenticationResult.IdToken
	responseItems.RefreshToken = acOutput.AuthenticationResult.RefreshToken
	responseItems.TokenType = acOutput.AuthenticationResult.TokenType
	responseItems.ChallangeName = acOutput.ChallengeName
	responseItems.ChallangeParameter = acOutput.ChallengeParameters

	fmt.Printf("Successfully parsed response to model response %v. \n", *responseItems.AccessToken)
	return &responseItems, nil
}

func ParseRequest(body *string) (*models.RequestParameter, error) {
	requestParameter := models.RequestParameter{}
	err := json.Unmarshal([]byte(*body), &requestParameter)

	if err != nil {
		fmt.Printf("Repository: There was an error marshalling the bytes to struct: %v", err.Error())
		return nil, err
	}

	requestAsBytes, _ := json.Marshal(requestParameter)
	fmt.Printf("Repository: marshalled bytes to struct: %+v. \n", string(requestAsBytes))

	// Mandatory Fiels Check for the Request
	err = mandatoryFieldsCheck(requestParameter)
	if err != nil {
		return nil, err
	}

	return &requestParameter, nil
}

func mandatoryFieldsCheck(requestParameter models.RequestParameter) error {

	if requestParameter.RefreshToken == nil {
		return customErrors.ErrEmptyRefreshTokenInRequest
	}

	if len(*requestParameter.RefreshToken) < 20 {
		return customErrors.ErrInvalidRefreshTokenInRequest
	}

	if requestParameter.UserId == nil {
		return customErrors.ErrEmptyUserIdInRequest
	}

	return nil
}

// Secret hash is not a client secret itself, but a base64 encoded hmac-sha256
// hash.
// The actual AWS documentation on how to compute this hash is here:
// https://docs.aws.amazon.com/cognito/latest/developerguide/signing-up-users-in-your-app.html#cognito-user-pools-computing-secret-hash
func ComputeSecretHash(username *string) string {
	mac := hmac.New(sha256.New, []byte(config.ClientSecret))
	mac.Write([]byte(*username + config.ClientId))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
