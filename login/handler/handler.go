package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"login/service"
	"login/service/models"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %v.\n", request.Body)
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %v: %v\n", key, value)
	}

	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}
	httpResponse, err := service.LoginUser(&request.Body)

	if err != nil {
		errorMessage := err.Error()
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}

	respAsBytes, _ := json.Marshal(httpResponse)
	return events.APIGatewayProxyResponse{Body: string(respAsBytes), StatusCode: 200, Headers: header}, nil
}
