package repository

import (
	"confirm-signup/add-wallet/service/config"
	"confirm-signup/add-wallet/service/models"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func AttributeBuilder(currencyName *string, currencySymbol *string, userIdInCognito *string) (map[string]*dynamodb.AttributeValue, error) {
	date := time.Now().Format(time.RFC3339Nano)
	queryParameter := models.QueryParameter{
		Currency:     currencySymbol,
		Name:         currencyName,
		CreationDate: &date,
		UpdatedDate:  &date,
		Sk:           config.SkPrefix + date,
		Pk:           userIdInCognito,
	}

	av, err := dynamodbattribute.MarshalMap(queryParameter)
	fmt.Printf("marshalled struct: %+v", av)
	return av, err
}
