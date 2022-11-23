package handler

import (
	"context"
	"fmt"
	"login/service"
	"login/service/models"
)

func HandleRequest(ctx context.Context, request *map[string]string) (models.HttpResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request)

	response, err := service.LoginUser(request)

	if err != nil {
		statusCode := 500
		errMessage := err.Error()
		errorResponse := models.HttpResponse{
			StatusCode: &statusCode,
			Message:    &errMessage,
		}
		return errorResponse, err
	}

	statusCode := 200
	httpResponse := models.HttpResponse{
		Session:              response.Session,
		AuthenticationResult: response.AuthenticationResult,
		ChallengeName:        response.ChallengeName,
		StatusCode:           &statusCode,
	}
	return httpResponse, nil
}
