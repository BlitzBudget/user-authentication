package service

import (
	"add-wallet/service/repository"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func SaveRequest(body *string) {
	// snippet-start:[dynamodb.go.create_item.session]
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	// snippet-end:[dynamodb.go.create_item.session]

	av, err := repository.AttributeBuilder(body)
	if err != nil {
		panic(fmt.Sprintf("SaveRequest: Got error marshalling new item: %v", err))
	}

	err = repository.CreateItem(av, svc)
	if err != nil {
		panic(fmt.Sprintf("SaveRequest: Got error calling PutItem: %v", err))
	}

	fmt.Println("Successfully added the item!")
}
