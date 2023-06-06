package repository

import (
	"encoding/json"
	"fmt"
	"global-signout/service/errors"
	"global-signout/service/models"
	"net/mail"
)

func ParseRequest(body *string) (*models.RequestParameter, error) {
	requestParamater := models.RequestParameter{}
	err := json.Unmarshal([]byte(*body), &requestParamater)

	if err != nil {
		fmt.Printf("Repository: There was an error marshalling the bytes to struct: %v \n", err.Error())
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

	if requestParamater.AccessToken == nil {
		return errors.ErrEmptyAccessTokenInRequest
	}

	_, err := mail.ParseAddress(*requestParamater.Email)
	if err != nil {
		return errors.ErrInvalidEmailInRequest
	}

	if len(*requestParamater.AccessToken) <= 20 {
		return errors.ErrMinimumLengthRequiredAccessToken
	}

	return nil
}
