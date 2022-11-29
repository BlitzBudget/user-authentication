package repository

import (
	"testing"
)

func TestAttributeBuilder(t *testing.T) {
	pk := "Wallet#2022-05-12T20:25:19Z"
	creationDate := "2022-05-12T20:25:19Z"
	updatedDate := "2022-05-12T20:25:19Z"
	walletCurrency := "Dollar"
	walletName := "walletName"
	body := "{\"pk\":\"" + pk + "\",\"creation_date\":\"" + creationDate + "\",\"wallet_currency\":\"" + walletCurrency + "\",\"wallet_name\":\"" + walletName + "\",\"updated_date\":\"" + updatedDate + "\"}"

	got, err := AttributeBuilder(&body)
	if err != nil {
		t.Errorf("AttributeBuilder() error = %v", err)
		return
	}

	if got == nil {
		t.Errorf("AttributeBuilder() is null")
		return
	}

	if *(*got["creation_date"]).S == "" {
		t.Errorf("name creationDate to DynamoDB attribute not correct, got = %v", *(*got["creation_date"]).S)
		return
	}

	if *(*got["updated_date"]).S == "" {
		t.Errorf("name updated_date to DynamoDB attribute not correct, got = %v", *(*got["updated_date"]).S)
		return
	}

	if *(*got["pk"]).S != pk {
		t.Errorf("name pk to DynamoDB attribute not correct, got = %v, want = %v", *(*got["pk"]).S, pk)
		return
	}

	if *(*got["sk"]).S == "" {
		t.Errorf("name sk to DynamoDB attribute not correct, got = %v", *(*got["sk"]).S)
		return
	}

	if *(*got["wallet_currency"]).S != walletCurrency {
		t.Errorf("name wallet_currency to DynamoDB attribute not correct, got = %v, want = %v", *(*got["wallet_currency"]).S, walletCurrency)
		return
	}

	if *(*got["wallet_name"]).S != walletName {
		t.Errorf("name wallet_name to DynamoDB attribute not correct, got = %v, want = %v", *(*got["wallet_name"]).S, walletName)
		return
	}
}
