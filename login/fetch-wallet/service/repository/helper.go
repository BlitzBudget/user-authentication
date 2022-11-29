package repository

import (
	"errors"
	"fmt"
	"login/fetch-wallet/service/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func ParseResponse(result *dynamodb.QueryOutput) (models.WalletResponseItems, error) {

	if result.Items == nil {
		msg := "no Items found"
		return nil, errors.New(msg)
	}

	walletResponseItems := models.WalletResponseItems{}
	var err error

	for k, v := range result.Items {
		walletResponseItem := models.WalletResponseItem{}

		err = dynamodbattribute.UnmarshalMap(v, &walletResponseItem)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record %v, %v", k, err))
		}
		walletResponseItems = append(walletResponseItems, &walletResponseItem)
	}

	fmt.Printf("Parsed %v Items. \n", len(walletResponseItems))
	return walletResponseItems, nil
}
