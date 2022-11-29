package service

import (
	"confirm-signup/add-wallet/locale"
	"confirm-signup/add-wallet/service/repository"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func AddNewWallet(localeHeader *string, sess *session.Session, userIdInCognito *string) error {

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	currencyName := locale.ConvertToCurrencyName(localeHeader)
	currencySymbol := locale.ConvertToSymbol(currencyName)

	av, err := repository.AttributeBuilder(currencyName, currencySymbol, userIdInCognito)
	if err != nil {
		panic(fmt.Sprintf("SaveRequest: Got error marshalling new item: %v \n", err))
	}

	err = repository.CreateItem(av, svc)
	if err != nil {
		panic(fmt.Sprintf("SaveRequest: Got error calling PutItem: %v \n", err))
	}

	fmt.Println("Successfully added the item!")
	return err
}
