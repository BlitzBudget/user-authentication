package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"update-user/service"
	"update-user/service/models"

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

	err := service.UpdateUserAttributes(&request.Body)
	if err != nil {
		errorMessage := err.Error()
		errorRespose := models.ErrorHttpResponse{
			Message: &errorMessage,
		}
		errorAsBytes, _ := json.Marshal(errorRespose)
		return events.APIGatewayProxyResponse{Body: string(errorAsBytes), StatusCode: 500, Headers: header}, nil
	}

	message := "Update of User Attributes is successful"
	httpResponse := models.HttpResponse{
		Message: &message,
	}
	responseAsbytes, _ := json.Marshal(httpResponse)
	return events.APIGatewayProxyResponse{Body: string(responseAsbytes), StatusCode: 200, Headers: header}, nil
}
