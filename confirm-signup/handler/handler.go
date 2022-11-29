package handler

import (
	"confirm-signup/service"
	"confirm-signup/service/models"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Body size = %v.\n", request.Body)

	locale := ""
	for key, value := range request.Headers {
		if key == "CloudFront-Viewer-Country" {
			fmt.Printf("The Cloud front local selected is %v \n", value)
			locale = value
		}
	}

	header := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS,POST",
		"Access-Control-Allow-Credentials": "true",
	}

	httpResponse, err := service.ConfirmSignUp(&request.Body, &locale)
	if err != nil {
		errorMessage := err.Error()
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}

	responseAsBytes, _ := json.Marshal(httpResponse)
	return events.APIGatewayProxyResponse{Body: string(responseAsBytes), StatusCode: 200, Headers: header}, nil
}
