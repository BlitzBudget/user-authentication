package config

import "os"

var ClientId = os.Getenv("CLIENT_ID")
var ClientSecret = os.Getenv("CLIENT_SECRET")
var UserPoolId = os.Getenv("USERPOOL_ID")

var UserIdInCognito = "custom:financialPortfolioId"
