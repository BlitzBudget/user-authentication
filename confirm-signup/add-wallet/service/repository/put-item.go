package repository

import (
	"confirm-signup/add-wallet/service/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func CreateItem(av map[string]*dynamodb.AttributeValue, svc dynamodbiface.DynamoDBAPI) error {
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.TableName),
	}

	_, err := svc.PutItem(input)
	return err
}
