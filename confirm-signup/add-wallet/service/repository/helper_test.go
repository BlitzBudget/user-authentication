package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	currencyName := "Euro"
	currencySymbol := "€"
	userIdInCognito := "userId"

	got, err := AttributeBuilder(&currencyName, &currencySymbol, &userIdInCognito)
	if err != nil {
		t.Errorf("AttributeBuilder() error = %v \n", err)
		return
	}

	if got == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *(*got["creation_date"]).S == "" {
		t.Errorf("name creationDate to DynamoDB attribute not correct, got = %v \n", *(*got["creation_date"]).S)
		return
	}

	if *(*got["updated_date"]).S == "" {
		t.Errorf("name updated_date to DynamoDB attribute not correct, got = %v \n", *(*got["updated_date"]).S)
		return
	}

	if *(*got["pk"]).S != userIdInCognito {
		t.Errorf("name pk to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["pk"]).S, userIdInCognito)
		return
	}

	if *(*got["sk"]).S == "" {
		t.Errorf("name sk to DynamoDB attribute not correct, got = %v \n", *(*got["sk"]).S)
		return
	}

	if *(*got["wallet_currency"]).S != currencySymbol {
		t.Errorf("name wallet_currency to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["wallet_currency"]).S, currencySymbol)
		return
	}

	if *(*got["wallet_name"]).S != currencyName {
		t.Errorf("name wallet_name to DynamoDB attribute not correct, got = %v, want = %v \n", *(*got["wallet_name"]).S, currencyName)
		return
	}
}
