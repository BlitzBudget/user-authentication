package repository

import (
	"encoding/json"
	"fmt"
	"global-signout/service/errors"
	"global-signout/service/models"
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

	if requestParamater.AccessToken == nil {
		return errors.ErrEmptyAccessTokenInRequest
	}

	if len(*requestParamater.AccessToken) <= 20 {
		return errors.ErrMinimumLengthRequiredAccessToken
	}

	return nil
}
