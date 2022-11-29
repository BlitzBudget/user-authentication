package repository

import (
	"encoding/json"
	"fmt"
	"update-user/service/errors"
	"update-user/service/models"
)

func ParseRequest(body *string) (*models.RequestParameter, error) {
	requestParameter := models.RequestParameter{}
	err := json.Unmarshal([]byte(*body), &requestParameter)

	if err != nil {
		fmt.Printf("Repository: There was an error marshalling the bytes to struct: %v \n", err.Error())
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

	if requestParameter.AccessToken == nil {
		return errors.ErrEmptyAccessTokenInRequest
	}

	if len(*requestParameter.AccessToken) < 20 {
		return errors.ErrInvalidAccessTokenInRequest
	}

	return nil
}
