package models

import "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

type HttpResponse struct {
	StatusCode           *int                                              `validate:"required" json:"statusCode"`
	Session              *string                                           `validate:"required" json:"Session"`
	Message              *string                                           `validate:"required" json:"Message"`
	AuthenticationResult *cognitoidentityprovider.AuthenticationResultType `validate:"required" json:"AuthenticationResult"`
	ChallengeName        *string                                           `validate:"required" json:"ChallengeName"`
}

type ResponseModel struct {
	Session              *string                                           `validate:"required" json:"Session"`
	AuthenticationResult *cognitoidentityprovider.AuthenticationResultType `validate:"required" json:"AuthenticationResult"`
	ChallengeName        *string                                           `validate:"required" json:"ChallengeName"`
}
