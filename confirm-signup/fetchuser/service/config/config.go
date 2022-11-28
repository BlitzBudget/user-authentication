package config

import "os"

var ClientId = os.Getenv("CLIENT_ID")
var ClientSecret = os.Getenv("CLIENT_SECRET")

var FirstNameUserAttribute = "name"
var LastNameUserAttribute = "family_name"
