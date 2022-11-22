package config

import "os"

const AuthFlow = "ADMIN_USER_PASSWORD_AUTH"
const UserIdPrefix = "User#"
var ExportFormat = "XLS"

var (
	ClientId     = os.Getenv("COGNITO_CLIENT_ID")
	UserPoolId   = os.Getenv("COGNITO_USER_POOL_ID")
	ClientSecret = os.Getenv("COGNITO_CLIENT_SECRET")
)
