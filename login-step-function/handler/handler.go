package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"login/service"
	"login/service/models"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request *string) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request)

	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}
	session, err := service.LoginUser(request)

	if err != nil {
		errorMessage := err.Error()
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}

	message := "Successfuly sent an SMS to your mobile number"
	response := models.HttpResponse{
		Message: &message,
		Session: session,
	}
	respAsBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{Body: string(respAsBytes), StatusCode: 200, Headers: header}, nil
}
