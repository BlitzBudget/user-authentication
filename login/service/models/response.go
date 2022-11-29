package models

import (
	fetchWalletModels "login/fetch-wallet/service/models"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type HttpResponse struct {
	StatusCode           *int                                              `validate:"required" json:"statusCode"`
	Session              *string                                           `validate:"required" json:"Session"`
	Message              *string                                           `validate:"required" json:"Message"`
	AuthenticationResult *cognitoidentityprovider.AuthenticationResultType `validate:"required" json:"AuthenticationResult"`
	ChallengeName        *string                                           `validate:"required" json:"ChallengeName"`
	UserAttributes       []*UserAttribute                                  `validate:"required" json:"UserAttributes"`
	Wallet               []*fetchWalletModels.WalletResponseItem           `validate:"required" json:"Wallet"`
}

type LoginResponseModel struct {
	Session              *string                                           `validate:"required" json:"Session"`
	AuthenticationResult *cognitoidentityprovider.AuthenticationResultType `validate:"required" json:"AuthenticationResult"`
	ChallengeName        *string                                           `validate:"required" json:"ChallengeName"`
}

// Specifies whether the attribute is standard or custom.
type UserAttribute struct {

	// The name of the attribute.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The value of the attribute.
	//
	// Value is a sensitive parameter and its value will be
	// replaced with "sensitive" in string returned by AttributeType's
	// String and GoString methods.
	Value *string `type:"string" sensitive:"true"`
}
