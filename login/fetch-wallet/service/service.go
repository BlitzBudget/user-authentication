package service

import (
	"fmt"
	"log"
	"login/fetch-wallet/service/models"
	"login/fetch-wallet/service/repository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func FetchWallet(userIdInCognito *string, sess *session.Session) ([]*models.WalletResponseItem, error) {

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	var queryOutput *dynamodb.QueryOutput
	queryOutput, err := repository.QueryItem(userIdInCognito, svc)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %v \n", err)
		return nil, err
	}

	var walletResponseItems []*models.WalletResponseItem
	walletResponseItems, err = repository.ParseResponse(queryOutput)
	if err != nil {
		log.Fatalf("Got error parsing Response Item: %v \n", err)
		return nil, err
	}

	fmt.Printf("Successfully retrieved %v items with the consumed capacity of %v'", *queryOutput.Count, queryOutput.ConsumedCapacity)
	return walletResponseItems, nil
}
