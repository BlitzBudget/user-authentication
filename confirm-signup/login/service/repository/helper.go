package repository

import (
	"confirm-signup/login/service/errors"
	"confirm-signup/service/config"
	"confirm-signup/service/models"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/mail"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ParseResponse(adminInitiateAuthOutput *cognitoidentityprovider.AdminInitiateAuthOutput) *models.LoginResponseModel {
	responseModel := models.LoginResponseModel{
		Session:              adminInitiateAuthOutput.Session,
		AuthenticationResult: adminInitiateAuthOutput.AuthenticationResult,
		ChallengeName:        adminInitiateAuthOutput.ChallengeName,
	}

	return &responseModel
}

func ParseRequest(body *string) (*models.RequestParameter, error) {
	requestParamater := models.RequestParameter{}
	err := json.Unmarshal([]byte(*body), &requestParamater)

	if err != nil {
		fmt.Printf("Repository: There was an error marshalling the bytes to struct: %v", err.Error())
		return nil, err
	}

	fmt.Printf("Repository: marshalled bytes to struct: %+v. \n", requestParamater)

	// Mandatory Fiels Check for the Request
	err = mandatoryFieldsCheck(requestParamater)
	if err != nil {
		return nil, err
	}

	return &requestParamater, nil
}

func mandatoryFieldsCheck(requestParamater models.RequestParameter) error {

	if requestParamater.Email == nil {
		return errors.ErrEmptyEmailInRequest
	}

	if requestParamater.Password == nil {
		return errors.ErrEmptyPasswordInRequest
	}

	_, err := mail.ParseAddress(*requestParamater.Email)
	if err != nil {
		return errors.ErrInvalidEmailInRequest
	}

	if len(*requestParamater.Password) <= 7 {
		return errors.ErrMinimumLengthRequiredPassword
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
