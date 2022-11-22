package config

import "os"

var ClientId = os.Getenv("CLIENT_ID")
var ClientSecret = os.Getenv("CLIENT_SECRET")
var AuthFlow = "REFRESH_TOKEN_AUTH"
