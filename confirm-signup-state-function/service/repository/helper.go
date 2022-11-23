package repository

import (
	"confirm-signup/service/config"
	"confirm-signup/service/errors"
	"confirm-signup/service/models"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/mail"
)

func MandatoryFieldsCheck(requestParameter *models.RequestParameter) error {

	if requestParameter == nil {
		return errors.ErrEmptyRequest
	}

	if requestParameter.Email == nil {
		return errors.ErrEmptyEmailInRequest
	}

	_, err := mail.ParseAddress(*requestParameter.Email)
	if err != nil {
		return errors.ErrInvalidEmailInRequest
	}

	if requestParameter.ConfirmationCode == nil {
		return errors.ErrEmptyConfirmationCodeInRequest
	}

	if len(*requestParameter.ConfirmationCode) < 4 || len(*requestParameter.ConfirmationCode) > 6 {
		return errors.ErrInvalidConfirmationCodeInRequest
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
