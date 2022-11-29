package repository

import (
	"confirm-signup/add-wallet/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type mockDynamodbClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamodbClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	result := dynamodb.PutItemOutput{}

	wallet := "wallet"
	startsWithDate := "2022-10-20"
	endsWithDate := "2022-12-01"
	item := models.QueryParameter{
		Pk:           &wallet,
		CreationDate: &startsWithDate,
		UpdatedDate:  &endsWithDate,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return &result, err
	}

	result = dynamodb.PutItemOutput{
		Attributes: av,
	}

	return &result, nil
}

func TestPutItem(t *testing.T) {
	mockSvc := &mockDynamodbClient{}
	wallet := "wallet"
	startsWithDate := "2022-10-20"
	endsWithDate := "2022-12-01"
	item := models.QueryParameter{
		Pk:           &wallet,
		CreationDate: &startsWithDate,
		UpdatedDate:  &endsWithDate,
	}

	av, _ := dynamodbattribute.MarshalMap(item)

	err := CreateItem(av, mockSvc)
	if err != nil {
		t.Errorf("PutItem() error = %v \n", err)
		return
	}

}
