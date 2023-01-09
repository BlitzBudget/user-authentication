package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"refresh-token/service"
	"refresh-token/service/models"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jsonReq, _ := json.Marshal(request)
	fmt.Printf("Processing request data for request %v.\n", string(jsonReq))

	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}

	response, err := service.InitiateAuth(&request.Body)
	if err != nil {
		errorMessage := err.Error()
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}

	responseAsBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{Body: string(responseAsBytes), StatusCode: 200, Headers: header}, nil
}
