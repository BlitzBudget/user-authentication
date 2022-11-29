package repository

import (
	"confirm-signup/add-wallet/service/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func CreateItem(av map[string]*dynamodb.AttributeValue, svc dynamodbiface.DynamoDBAPI) error {
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.TableName),
	}

	ptOutput, err := svc.PutItem(input)
	fmt.Printf("CreateItem:: The consumed capacity is %v \n", ptOutput)

	return err
}
