package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"forgot-password/service/config"
	"forgot-password/service/errors"
	"forgot-password/service/models"
	"net/mail"
)

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

	if requestParameter.Email == nil {
		return errors.ErrEmptyEmailInRequest
	}

	_, err := mail.ParseAddress(*requestParameter.Email)
	if err != nil {
		return errors.ErrInvalidEmailInRequest
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
