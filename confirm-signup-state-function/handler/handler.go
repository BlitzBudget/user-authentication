package handler

import (
	"confirm-signup/service"
	"confirm-signup/service/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request models.RequestParameter) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request)

	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}

	err := service.ConfirmSignUp(&request)
	if err != nil {
		errorMessage := err.Error()
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}

	response := models.HTTPResponse{
		Email:    request.Email,
		Password: request.NewPassword,
	}
	responseAsBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{Body: string(responseAsBytes), StatusCode: 500, Headers: header}, nil
}
