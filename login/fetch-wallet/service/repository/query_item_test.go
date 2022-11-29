package repository

import (
	"login/fetch-wallet/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
)

type mockDynamodbClient struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamodbClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	result := dynamodb.QueryOutput{}

	wallet := "wallet"
	item := models.WalletResponseItem{
		Pk: &wallet,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return &result, err
	}

	result = dynamodb.QueryOutput{
		Items: []map[string]*dynamodb.AttributeValue{
			av,
		},
	}

	return &result, nil
}

func TestQueryItem(t *testing.T) {
	assert := assert.New(t)

	mockSvc := &mockDynamodbClient{}
	userId := "userId"

	queryOutput, err := QueryItem(&userId, mockSvc)
	if err != nil {
		t.Errorf("QueryItem() error = %v \n", err)
		return
	}

	assert.NotNil(queryOutput)

}
