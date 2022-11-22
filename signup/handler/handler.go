package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"signup/service"
	"signup/service/models"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %v.\n", request.RequestContext.RequestID)
	requstAsByte, _ := json.Marshal(request)
	fmt.Printf("Body size = %v.\n", string(requstAsByte))
	localeHeader := request.Headers["CloudFront-Viewer-Country"]
	fmt.Printf("The Locale in the header is %v", localeHeader)

	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}

	err := service.SignupUser(&request.Body)
	if err != nil {
		errorMessage := err.Error()
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200, Headers: header}, nil
}
