package config

import "os"

var TableName = os.Getenv("TABLE_NAME")
var ScanIndexForward = false
var ProjectionExpression = "wallet_currency, wallet_name, sk, pk, creation_date"

const (
	SkPrefix = "Wallet#"
)
